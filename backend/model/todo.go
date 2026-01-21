package model

type Todo struct { // Struktur data untuk item todo.
	ID     int    `json:"id"` // Unique identifier for the todo item.
	Title  string `json:"title"` // Judul atau deskripsi todo.
	Done   bool   `json:"done"` // Status apakah todo sudah selesai atau belum.
}