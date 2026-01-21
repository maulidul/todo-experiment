import { defineStore } from 'pinia'
import axios from 'axios'

export const useTodoStore = defineStore('todo', {
  state: () => ({
    todos: [],
    loading: false
  }),
  actions: {
    async fetchTodos() {
      this.loading = true
      const res = await axios.get('http://localhost:8085/todos')
      this.todos = res.data
      this.loading = false
    },
    async addTodo(title) {
      await axios.post('http://localhost:8085/todos', { title })
      this.fetchTodos()
    },
    async toggleTodo(id) {
      await axios.put(`http://localhost:8085/todos/${id}`)
      this.fetchTodos()
    },
    async deleteTodo(id) {
      await axios.delete(`http://localhost:8085/todos/${id}`)
      this.fetchTodos()
    }
  }
})