package usecase

import (
	"backend/model" // Definisi model data.
	"backend/repository" // Mengelola akses data.
)

type TodoUsecase interface { // Antarmuka untuk logika bisnis todo.
	GetTodos() []model.Todo // Mendapatkan semua todo.
	AddTodo(title string) model.Todo // Menambahkan todo baru.
	ToggleStatus(id int) (model.Todo, bool) // Mengubah status todo (selesai/belum selesai).
	RemoveTodo(id int) bool // Menghapus todo.
}

type todoUsecase struct { // Implementasi TodoUsecase.
	repo repository.TodoRepository // Menggunakan TodoRepository untuk akses data.
}

func NewTodoUsecase(r repository.TodoRepository) TodoUsecase { // Fungsi untuk membuat instance TodoUsecase.
	return &todoUsecase{repo: r} // Mengembalikan instance todoUsecase dengan repository yang diberikan.
}

func (u *todoUsecase) GetTodos() []model.Todo { // Metode untuk mendapatkan semua todo.
	return u.repo.GetAll() // Mengambil semua todo dari repository.
}

func (u *todoUsecase) AddTodo(title string) model.Todo { // Metode untuk menambahkan todo baru.
	return u.repo.Create(model.Todo{ //	 Membuat todo baru di repository.
		Title: title, // Judul todo.
		Done:  false, // Status awal todo adalah belum selesai.	
	})
}

func (u *todoUsecase) ToggleStatus(id int) (model.Todo, bool) { // Metode untuk mengubah status todo.
	return u.repo.UpdateStatus(id) // Memperbarui status todo di repository.
}

func (u *todoUsecase) RemoveTodo(id int) bool { // Metode untuk menghapus todo.
	return u.repo.Delete(id) // Menghapus todo dari repository.
}
