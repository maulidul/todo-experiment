package main

import (
	"github.com/gin-gonic/gin" // Framework web Gin.
	"github.com/gin-contrib/cors" // Middleware CORS untuk Gin.
	"time"// Untuk pengaturan waktu.
	"backend/handler" // Menangani request dan response HTTP.
	"backend/repository" // Mengelola akses data.
	"backend/usecase" // Logika bisnis aplikasi.
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{ // Konfigurasi CORS untuk mengizinkan permintaan dari frontend.
		AllowOrigins:     []string{"http://localhost:5173"}, // Mengizinkan asal dari frontend.
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Metode HTTP yang diizinkan.
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, //
		AllowCredentials: true, // Mengizinkan kredensial (cookies, header otorisasi, dll).
		MaxAge: 12 * time.Hour, // Durasi cache preflight request.
	}))

	repo := repository.NewTodoRepository()
	uc := usecase.NewTodoUsecase(repo)
	handler.NewTodoHandler(r, uc)
	r.Run(":8085") // Menjalankan server pada port 8085.
}
