<template>
  <div class="container">
    <h1>Todo List</h1>

    <form @submit.prevent="submitTodo">
      <input v-model="title" placeholder="Tambah todo" />
      <button>Tambah</button>
    </form>

    <p v-if="store.loading">Loading...</p>

    <table class="todo-table" v-if="store.todos.length">
      <thead>
        <tr>
          <th>No</th>
          <th>Status</th>
          <th>Judul</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(todo, index) in store.todos" :key="todo.id">
          <td>{{ index + 1 }}</td>

          <td class="center">
            <input
              type="checkbox"
              :checked="todo.done"
              @change="store.toggleTodo(todo.id)"
            />
          </td>

          <td :class="{ done: todo.done }">
            {{ todo.title }}
          </td>

          <td class="center">
            <button class="btn-delete" @click="store.deleteTodo(todo.id)">
              Hapus
            </button>
          </td>
        </tr>
      </tbody>
    </table>  

    <p v-else>Tidak ada todo</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useTodoStore } from '../stores/todo'

const store = useTodoStore()
const title = ref('')

onMounted(() => {
  store.fetchTodos()
})

const submitTodo = () => {
  if (!title.value) return
  store.addTodo(title.value)
  title.value = ''
}
</script>

<style>
.container {
  width: 400px;
  margin: 40px auto;
  font-family: Arial;
}
button {
  margin-left: 8px;
}
.todo-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

.todo-table th,
.todo-table td {
  border: 1px solid #ddd;
  padding: 8px;
}
tbody tr td {
  color: black;
  background-color: #fff;
}

.todo-table th {
  background-color: #2090a8;
  text-align: center;
}

.center {
  text-align: center;
}

.done {
  text-decoration: line-through;
  color: gray;
}

.btn-delete {
  background: #e74c3c;
  color: white;
  border: none;
  padding: 4px 8px;
  cursor: pointer;
}

.btn-delete:hover {
  background: #c0392b;
}
</style>