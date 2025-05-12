<template>
  <div class="q-pa-md">
    <q-table flat bordered title="Books" :rows="rows" :columns="columns" row-key="name">
      <template v-slot:top>
        <q-btn icon="add" @click="onCreate()" />
        <!-- <q-space /> -->
        <!-- <q-input borderless dense debounce="300" v-model="filter" placeholder="Search">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input> -->
      </template>

      <template v-slot:body-cell-img>
        <q-td props="props">
          <q-img :src="props.row.image" />
        </q-td>
      </template>

      <template v-slot:body-cell-actions="props">
        <q-td :props="props">
          <q-btn icon="mode_edit" @click="onEdit(props.row.id)" />
          <q-btn icon="delete" @click="onDelete(props.row.id)" />
        </q-td>
      </template>
    </q-table>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const columns = ref([
  { name: 'id', label: 'ID', align: 'left', field: 'id', sortable: true },
  { name: 'image', label: 'Image', field: 'image', sortable: true },
  { name: 'title', label: 'Title', field: 'title', sortable: true },
  { name: 'author', label: 'Author', field: 'author', sortable: true },
  { name: 'category', label: 'Category', field: 'category' },
  { name: 'description', label: 'Description', field: 'description' },
  { name: 'published_date', label: 'Published Date', field: 'published_date' },
  { name: 'CreatedAt', label: 'Created At', field: 'CreatedAt', sortable: true },
  { name: 'UpdatedAt', label: 'Updated At', field: 'UpdatedAt', sortable: true },
  { name: 'actions', label: '', field: 'actions', sortable: true },
])

const rows = ref([])

const router = useRouter()

const fetchData = () => {
  fetch('https://www.melivecode.com/api/users')
    .then((response) => response.json())
    .then((result) => {
      rows.value = result
    })
}

const onEdit = (id) => {
  router.push(`edit/${id}`)
}

const onDelete = (id) => {
  console.log(id)
  deleteBook(id)
}

const onCreate = () => {
  router.push('/create')
}

const deleteBook = async (id) => {
  const endPoint = `http://127.0.0.1:3000/book/${id}`

  const requestOptions = {
    method: 'DELETE',
  }

  try {
    const response = await fetch(endPoint, requestOptions)

    if (!response.ok) {
      throw new Error(`CANNOT DELETE BOOK DATA ID: ${id}`)
    }

    const responseJson = await response.json()

    let message = `RESPONSE: ${JSON.stringify(responseJson)}\n`
    message += 'DELETED SUCCESS\n'
    message += `BOOK ID: ${id}`

    console.log(message)

    alert(responseJson.message)

    fetchData()
  } catch (error) {
    console.error(error)
  }
}

fetchData()
// [
//   {
//     id: 'Frozen Yogurt',
//     title: 159,
//     author: 6.0,
//     published_date: 24,
//     category: 4.0,
//     description: 87,
//     CreatedAt: '14%',
//     UpdatedAt: '1%'
//   },
// ]
</script>
