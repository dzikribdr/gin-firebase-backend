package main

import (
	"github.com/dzikribdr/gin-firebase-backend/config"
	"github.com/dzikribdr/gin-firebase-backend/models"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()
	config.InitDatabase()
	products := []models.Product{
	{
		Name:        "Nugget Ayam",
		Price:       35000,
		Category:    "Olahan Ayam",
		Stock:       50,
		Description: "Nugget ayam crispy siap goreng",
		ImageURL:    "https://i.ibb.co.com/MyNMzxgG/chicken-original.png",
	},
	{
		Name:        "Sosis Ayam",
		Price:       28000,
		Category:    "Olahan Ayam",
		Stock:       60,
		Description: "Sosis ayam premium siap masak",
		ImageURL:    "https://i.ibb.co.com/R423wGLh/sosis-ayam-premium-spicy-chicken-sausage-so-good-sogood5-1.jpg",
	},
	{
		Name:        "Dimsum Ayam",
		Price:       32000,
		Category:    "Olahan Ayam",
		Stock:       45,
		Description: "Dimsum ayam isi 10 pcs",
		ImageURL:    "https://i.ibb.co.com/Z6dYd0d3/lg-678603ec68281.jpg",
	},
	{
		Name:        "Karaage Ayam",
		Price:       38000,
		Category:    "Olahan Ayam",
		Stock:       35,
		Description: "Potongan ayam karaage berbumbu",
		ImageURL:    "https://i.ibb.co.com/N6WVgsP9/2p6yblhoj.jpg",
	},
	{
		Name:        "Chicken Katsu",
		Price:       40000,
		Category:    "Olahan Ayam",
		Stock:       30,
		Description: "Chicken katsu crispy siap goreng",
		ImageURL:    "https://i.ibb.co.com/TxcqqbPK/lg-60b4a32ca0fee.jpg",
	},

	{
		Name:        "Fish Roll",
		Price:       30000,
		Category:    "Seafood",
		Stock:       40,
		Description: "Olahan ikan roll siap masak",
		ImageURL:    "https://i.ibb.co.com/NdGZQmdC/brd-44261-ellafroze-fish-roll-500-gr-olahan-ikan-frozen-full01-c26f6f4e.jpg",
	},
	{
		Name:        "Tempura Udang",
		Price:       42000,
		Category:    "Seafood",
		Stock:       25,
		Description: "Tempura udang renyah siap goreng",
		ImageURL:    "https://i.ibb.co.com/Cp1L9ctV/photo.jpg",
	},
	{
		Name:        "Cumi Ring",
		Price:       45000,
		Category:    "Seafood",
		Stock:       20,
		Description: "Cumi ring beku siap olah",
		ImageURL:    "https://image1ws.indotrading.com/s3/productimages/webp/co47833/p370046/w600-h600/82df194e-19e1-45cf-82f8-e3c6e777d4cc.jpg",
	},
	{
		Name:        "Bakso Ikan",
		Price:       29000,
		Category:    "Seafood",
		Stock:       50,
		Description: "Bakso ikan kenyal dan gurih",
		ImageURL:    "https://i.ibb.co.com/5x9d32gW/md-641428eb2cb82.jpg",
	},
	{
		Name:        "Crab Stick",
		Price:       27000,
		Category:    "Seafood",
		Stock:       55,
		Description: "Crab stick praktis untuk berbagai masakan",
		ImageURL:    "https://image.made-in-china.com/202f0j00pfFvAcBtAGbe/Good-Quality-Surimi-Crab-Stick-Frozen-Crab-Meat-Sticks.jpg",
	},
}
	for _, p := range products {
		config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
