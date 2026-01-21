package repository

import "backend/model" // Mengimpor paket model untuk struktur data Todo.

type TodoRepository interface { // Antarmuka untuk operasi data todo.
	GetAll() []model.Todo // Mendapatkan semua todo.
	Create(todo model.Todo) model.Todo // Membuat todo baru.
	UpdateStatus(id int) (model.Todo, bool) // Memperbarui status todo.
	Delete(id int) bool // Menghapus todo.
}

type todoRepo struct { // Implementasi TodoRepository.
	data []model.Todo // Menyimpan data todo dalam slice.
}

func NewTodoRepository() TodoRepository { // Fungsi untuk membuat instance TodoRepository.
	return &todoRepo{ // Mengembalikan instance todoRepo dengan data awal kosong.
		data: []model.Todo{}, // Inisialisasi slice kosong.
	}
}

func (r *todoRepo) GetAll() []model.Todo { // Metode untuk mendapatkan semua todo.
	return r.data // Mengembalikan semua data todo.
}

func (r *todoRepo) Create(todo model.Todo) model.Todo { // Metode untuk membuat todo baru.
	todo.ID = len(r.data) + 1 // Menetapkan ID unik berdasarkan panjang data saat ini.
	r.data = append(r.data, todo) // Menambahkan todo baru ke slice data.
	return todo // Mengembalikan todo yang baru dibuat.
}

func (r *todoRepo) UpdateStatus(id int) (model.Todo, bool) { // Metode untuk memperbarui status todo.
	for i, t := range r.data { // Mencari todo berdasarkan ID.
		if t.ID == id { // Jika todo ditemukan.
			r.data[i].Done = !t.Done // Membalikkan status Done.
			return r.data[i], true //
		}
	}
	return model.Todo{}, false // Mengembalikan false jika todo tidak ditemukan.
}

func (r *todoRepo) Delete(id int) bool { // Metode untuk menghapus todo.
	for i, t := range r.data { // Mencari todo berdasarkan ID.
		if t.ID == id { // Jika todo ditemukan.
			r.data = append(r.data[:i], r.data[i+1:]...) // Menghapus todo dari slice.
			return true 
		}
	}
	return false
}
