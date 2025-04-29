<template>
  <div class="q-pa-md">
    <q-table
      grid
      flat bordered
      card-class="bg-light-green-7 text-white"
      title="Books"
      :rows="rows"
      :columns="columns"
      row-key="name"
      :filter="filter"
      hide-header
    >
      <template v-slot:body-cell-actions="props">
        <q-td :props="props">
          <q-btn icon="mode_edit" @click="onEdit(props.row.id)" />
          <q-btn icon="delete" @click="onDelete(props.row.id)" />
        </q-td>
      </template>

      <template v-slot:top>
        <q-btn icon="add"/>
        <q-space />
        <q-input borderless dense debounce="300" v-model="filter" placeholder="Search">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>

      </template>
    </q-table>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const columns = ref([
  { name: 'id', label: 'ID', align: 'left', field: 'id', sortable: true },
  { name: 'title', label: 'Title', field: 'title', sortable: true },
  { name: 'author', label: 'Author', field: 'author', sortable: true },
  { name: 'category', label: 'Category', field: 'category' },
  { name: 'description', label: 'Description', field: 'description' },
  { name: 'published_date', label: 'Published Date', field: 'published_date' },
  { name: 'CreatedAt', label: 'Created At', field: 'CreatedAt', sortable: true },
  { name: 'UpdatedAt', label: 'Updated At', field: 'UpdatedAt', sortable: true },
  { name: 'actions', label: 'id', field: 'id', sortable: true },
])

const rows = ref([])

const fetchData = () => {
  fetch("https://www.melivecode.com/api/users")
    .then(response => response.json())
    .then((result) => {
      rows.value = result
    })
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
