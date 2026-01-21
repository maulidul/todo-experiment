package handler

import (
	"net/http" // Untuk status kode HTTP.
	"strconv" // Untuk konversi string ke integer.
	"github.com/gin-gonic/gin" // Framework web Gin.
	"backend/usecase" // Logika bisnis aplikasi.
)

type TodoHandler struct { // Struct handler untuk todo.
	uc usecase.TodoUsecase // Menggunakan TodoUsecase untuk logika bisnis.
}

func NewTodoHandler(r * gin.Engine, uc usecase.TodoUsecase) { // Fungsi untuk membuat route handler.
	h := &TodoHandler{uc: uc} // Inisialisasi TodoHandler dengan usecase.

	// Mendefinisikan route dan mengaitkannya dengan metode handler.	

	r.GET("/todos", h.getTodos) // Mendapatkan daftar todo.
	r.POST("/todos", h.createTodo) // Membuat todo baru.
	r.PUT("/todos/:id", h.toggleTodo) // Mengubah status todo (selesai/belum selesai).
	r.DELETE("/todos/:id", h.deleteTodo) // Menghapus todo.
}

func (h * TodoHandler) getTodos(c * gin.Context) { // Metode untuk mendapatkan daftar todo.
	c.JSON(http.StatusOK, h.uc.GetTodos()) // Mengembalikan daftar todo dalam format JSON.
}

func (h * TodoHandler) createTodo(c * gin.Context) { // Metode untuk membuat todo baru.
	var req struct {
		Title string `json:"title"`//binding:"required"` // Judul todo yang diterima dari request JSON.
	}

	if err := c.ShouldBindJSON(&req); err != nil { // Memeriksa kesalahan dalam binding JSON.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Mengembalikan kesalahan jika ada.
		return
	}

	todo := h.uc.AddTodo(req.Title)// Menambahkan todo baru menggunakan usecase.
	c.JSON(http.StatusCreated, todo)// Mengembalikan todo yang baru dibuat dalam format JSON.
}

func (h * TodoHandler) toggleTodo(c * gin.Context) { // Metode untuk mengubah status todo.
	id, _ := strconv.Atoi(c.Param("id")) // Mengonversi ID todo dari string ke integer.
	todo, ok := h.uc.ToggleStatus(id) // Mengubah status todo menggunakan usecase.

	if !ok { //	 Memeriksa apakah todo ditemukan.
		c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"}) // Mengembalikan kesalahan jika todo tidak ditemukan.
		return
	}

	c.JSON(http.StatusOK, todo)// Mengembalikan todo yang diperbarui dalam format JSON.
}

func (h * TodoHandler) deleteTodo(c * gin.Context) { // Metode untuk menghapus todo.
	id, _ := strconv.Atoi(c.Param("id")) // Mengonversi ID todo dari string ke integer.
	if !h.uc.RemoveTodo(id) { // Menghapus todo menggunakan usecase.
		c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"}) // Mengembalikan kesalahan jika todo tidak ditemukan.
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"}) // Mengembalikan pesan sukses jika todo berhasil dihapus.
}
