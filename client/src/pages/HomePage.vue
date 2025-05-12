<template>
  <div class="q-pa-md">
    <q-table flat bordered title="Books" :rows="rows" :columns="columns" row-key="name">
      <template v-slot:top>
        <q-btn icon="add" @click="onCreate()" />
      </template>

      <template v-slot:body-cell-image="props">
        <q-td :props="props">
          <q-img :src="`http://127.0.0.1:3000/files/${props.row.image}`" />
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
  { name: 'status', label: 'Status', field: 'status' },
  { name: 'description', label: 'Description', field: 'description' },
  { name: 'publishedAt', label: 'Published Date', field: 'published_at' },
  { name: 'CreatedAt', label: 'Created At', field: 'created_at', sortable: true },
  { name: 'UpdatedAt', label: 'Updated At', field: 'updated_at', sortable: true },
  { name: 'actions', label: '', field: 'actions', sortable: true },
])

const rows = ref([])

const router = useRouter()

const fetchData = async () => {
  const endPoint = 'http://127.0.0.1:3000/book'

  try {
    const response = await fetch(endPoint)

    console.log(response)

    if (!response.ok) {
      throw new Error('Cannot fetch book data')
    }

    const responseJson = await response.json()

    console.log(responseJson)

    rows.value = responseJson.data
  } catch (error) {
    console.error(error)
  }
}

const onEdit = (id) => {
  router.push(`edit/${id}`)
}

const onDelete = (id) => {
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
</script>
