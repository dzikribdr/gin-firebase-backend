package main

import (
	"log"
	"os"

	"github.com/dzikribdr/gin-firebase-backend/config"
	"github.com/dzikribdr/gin-firebase-backend/pkg/logger"
	"github.com/dzikribdr/gin-firebase-backend/routes"

	"github.com/joho/godotenv"
)

func main() {
    // Load .env
    if err := godotenv.Load(); err != nil {
        log.Println("File .env tidak ditemukan, menggunakan environment variable sistem")
    }

    // Inisialisasi logger (WAJIB)
    logger.Init()

    // Firebase
    config.InitFirebase()

    // Database
    config.InitDatabase()

    // Router
    router := routes.SetupRouter()

    // Jalankan server
    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8081"
    }

    log.Printf("Server berjalan di http://localhost:%s", port)
    log.Printf("Health check: http://localhost:%s/v1/health", port)

    if err := router.Run(":" + port); err != nil {
        log.Fatalf("Gagal menjalankan server: %v", err)
    }
}