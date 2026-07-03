package services

import (
	"errors"

	"github.com/dzikribdr/gin-firebase-backend/models"
	"github.com/dzikribdr/gin-firebase-backend/repositories"
)

type OrderService struct {
	orderRepo   *repositories.OrderRepository
	cartRepo    *repositories.CartRepository
	productRepo *repositories.ProductRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderRepo:   repositories.NewOrderRepository(),
		cartRepo:    repositories.NewCartRepository(),
		productRepo: repositories.NewProductRepository(),
	}
}

func (s *OrderService) Checkout(userID uint, req *models.CheckoutRequest) (*models.Order, error) {
	// Ambil cart user
	cartItems, err := s.cartRepo.GetByUserID(userID)
	if err != nil || len(cartItems) == 0 {
		return nil, errors.New("keranjang belanja kosong")
	}

	// Bangun order items + hitung total
	var orderItems []models.OrderItem
	var totalAmount float64

	for _, item := range cartItems {
		product, err := s.productRepo.FindByID(item.ProductID)
		if err != nil {
			return nil, errors.New("produk " + item.Product.Name + " tidak ditemukan")
		}
		if product.Stock < item.Quantity {
			return nil, errors.New("stok produk " + product.Name + " tidak mencukupi")
		}

		subtotal := product.Price * float64(item.Quantity)
		totalAmount += subtotal

		orderItems = append(orderItems, models.OrderItem{
			ProductID:   product.ID,
			ProductName: product.Name,
			Price:       product.Price,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})

		// Kurangi stok
		newStock := product.Stock - item.Quantity
		if err := s.productRepo.UpdateStock(product.ID, newStock); err != nil {
			return nil, err
		}
	}

	// Buat order
	order := &models.Order{
		UserID:          userID,
		Status:          models.OrderStatusPending,
		TotalAmount:     totalAmount,
		ShippingAddress: req.ShippingAddress,
		Notes:           req.Notes,
		PaymentMethod:   req.PaymentMethod,
		Items:           orderItems,
	}

	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	// Kosongkan cart
	_ = s.cartRepo.ClearCart(userID)

	return order, nil
}

func (s *OrderService) GetMyOrders(userID uint, page, limit int) ([]models.Order, int64, error) {
	return s.orderRepo.GetByUserID(userID, page, limit)
}

func (s *OrderService) GetOrderByID(orderID, userID uint) (*models.Order, error) {
	return s.orderRepo.GetByID(orderID, userID)
}

func (s *OrderService) GetAllOrders(page, limit int) ([]models.Order, int64, error) {
	return s.orderRepo.GetAll(page, limit)
}

// ConfirmPayment dipanggil setelah aplikasi menerima callback sukses dari
// CashLess (deep link pasarmalam://payment-callback?status=success).
// Ini satu-satunya tempat status order pindah dari "pending" ke "processing"
// untuk metode global_institute_pay — sebelumnya TIDAK ADA endpoint ini,
// sehingga status order tidak pernah berubah walau pembayaran di CashLess sukses.
func (s *OrderService) ConfirmPayment(orderID, userID uint) (*models.Order, error) {
	order, err := s.orderRepo.GetByID(orderID, userID)
	if err != nil {
		return nil, errors.New("order tidak ditemukan")
	}

	if order.PaymentMethod != "global_institute_pay" {
		return nil, errors.New("order ini tidak menggunakan CashLess (global_institute_pay)")
	}

	if order.Status != models.OrderStatusPending {
		// Sudah dikonfirmasi sebelumnya — anggap sukses (idempotent)
		return order, nil
	}

	if err := s.orderRepo.UpdateStatus(orderID, models.OrderStatusProcessing); err != nil {
		return nil, err
	}

	order.Status = models.OrderStatusProcessing
	return order, nil
}

func (s *OrderService) UpdateOrderStatus(orderID uint, status models.OrderStatus) error {
	validStatuses := map[models.OrderStatus]bool{
		models.OrderStatusPending:    true,
		models.OrderStatusProcessing: true,
		models.OrderStatusShipped:    true,
		models.OrderStatusDelivered:  true,
		models.OrderStatusCancelled:  true,
	}
	if !validStatuses[status] {
		return errors.New("status tidak valid")
	}
	return s.orderRepo.UpdateStatus(orderID, status)
}
