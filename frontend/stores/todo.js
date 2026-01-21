import { defineStore } from 'pinia' // Mengimpor fungsi defineStore dari Pinia untuk mendefinisikan store.
import axios from 'axios' // Mengimpor axios untuk melakukan permintaan HTTP.

export const useTodoStore = defineStore('todo', { // Mendefinisikan store Pinia dengan nama 'todo'.
  state: () => ({ // Mendefinisikan state awal untuk store.
    todos: [], // Array kosong untuk menyimpan daftar todo.
    loading: false // Boolean untuk menandai status pemuatan data.
  }),
  actions: { // Mendefinisikan aksi-aksi untuk mengelola todo.
    async fetchTodos() { // Aksi untuk mengambil daftar todo dari backend.
      this.loading = true
      const res = await axios.get('http://localhost:8085/todos')
      this.todos = res.data
      this.loading = false
    },
    async addTodo(title) { // Aksi untuk menambahkan todo baru.
      await axios.post('http://localhost:8085/todos', { title })
      this.fetchTodos() // Memperbarui daftar todo setelah penambahan.
    },
    async toggleTodo(id) { // Aksi untuk mengubah status todo (selesai/belum selesai).
      await axios.put(`http://localhost:8085/todos/${id}`)
      this.fetchTodos() // Memperbarui daftar todo setelah perubahan status.
    },
    async deleteTodo(id) { // Aksi untuk menghapus todo.
      await axios.delete(`http://localhost:8085/todos/${id}`)
      this.fetchTodos() // Memperbarui daftar todo setelah penghapusan.
    }
  }
})