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
		ImageURL:    "https://picsum.photos/400",
	},
	{
		Name:        "Sosis Ayam",
		Price:       28000,
		Category:    "Olahan Ayam",
		Stock:       60,
		Description: "Sosis ayam premium siap masak",
		ImageURL:    "https://picsum.photos/401",
	},
	{
		Name:        "Dimsum Ayam",
		Price:       32000,
		Category:    "Olahan Ayam",
		Stock:       45,
		Description: "Dimsum ayam isi 10 pcs",
		ImageURL:    "https://picsum.photos/402",
	},
	{
		Name:        "Karaage Ayam",
		Price:       38000,
		Category:    "Olahan Ayam",
		Stock:       35,
		Description: "Potongan ayam karaage berbumbu",
		ImageURL:    "https://picsum.photos/403",
	},
	{
		Name:        "Chicken Katsu",
		Price:       40000,
		Category:    "Olahan Ayam",
		Stock:       30,
		Description: "Chicken katsu crispy siap goreng",
		ImageURL:    "https://picsum.photos/404",
	},

	{
		Name:        "Fish Roll",
		Price:       30000,
		Category:    "Seafood",
		Stock:       40,
		Description: "Olahan ikan roll siap masak",
		ImageURL:    "https://picsum.photos/405",
	},
	{
		Name:        "Tempura Udang",
		Price:       42000,
		Category:    "Seafood",
		Stock:       25,
		Description: "Tempura udang renyah siap goreng",
		ImageURL:    "https://picsum.photos/406",
	},
	{
		Name:        "Cumi Ring",
		Price:       45000,
		Category:    "Seafood",
		Stock:       20,
		Description: "Cumi ring beku siap olah",
		ImageURL:    "https://picsum.photos/407",
	},
	{
		Name:        "Bakso Ikan",
		Price:       29000,
		Category:    "Seafood",
		Stock:       50,
		Description: "Bakso ikan kenyal dan gurih",
		ImageURL:    "https://picsum.photos/408",
	},
	{
		Name:        "Crab Stick",
		Price:       27000,
		Category:    "Seafood",
		Stock:       55,
		Description: "Crab stick praktis untuk berbagai masakan",
		ImageURL:    "https://picsum.photos/409",
	},
}
	for _, p := range products {
		config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
