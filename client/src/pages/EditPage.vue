<template>
  <q-page class="flex justify-center">
    <div class="q-pa-md full-width" style="max-width: 400px">
      <q-form @submit="onSubmit" class="q-gutter-md">
        <q-input filled v-model="id" label="ID" readonly />

        <q-input
          filled
          v-model="title"
          label="Your Title *"
          lazy-rules
          :rules="[(val) => (val && val.length > 0 && val.length < 10) || 'Please type something']"
        />

        <q-input
          filled
          v-model="author"
          label="Your Author *"
          lazy-rules
          :rules="[(val) => (val && val.length > 0) || 'Please type something']"
        />

        <q-input
          filled
          v-model="description"
          label="Your Description *"
          lazy-rules
          :rules="[(val) => (val && val.length > 0) || 'Please type something']"
        />

        <q-input
          filled
          v-model="publishedAt"
          mask="date"
          :rules="[(val) => (val && val.length > 0) || 'Please type something']"
          label="Your Published At *"
        >
          <template v-slot:append>
            <q-icon name="event" class="cursor-pointer">
              <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                <q-date v-model="publishedAt">
                  <div class="row items-center justify-end">
                    <q-btn v-close-popup label="Close" color="primary" flat />
                  </div>
                </q-date>
              </q-popup-proxy>
            </q-icon>
          </template>
        </q-input>

        <q-select
          label="Your Status *"
          lazy-rules
          :rules="[(val) => (val && val.length > 0) || 'Please type something']"
          transition-show="flip-up"
          transition-hide="flip-down"
          filled
          v-model="status"
          :options="options"
        />

        <div>
          <q-btn label="Edit" type="submit" color="green-14" />
        </div>
      </q-form>
    </div>
  </q-page>
</template>

<script setup>
import dayjs from 'dayjs'
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const title = ref(null)
const author = ref(null)
const description = ref(null)
const publishedAt = ref(null)
const status = ref(null)

const route = useRoute()
const router = useRouter()
const id = ref(route.params.id)

const options = ['available', 'borrowed', 'reserved', 'lost']

const fetchBook = async () => {
  const endPoint = `http://127.0.0.1:3000/book/${id.value}`

  try {
    const response = await fetch(endPoint)

    if (!response.ok) {
      throw new Error(`Cannot fetch book data, ID: ${id.value}`)
    }

    const responseJson = await response.json()
    console.log('Response: ', responseJson)

    setBook(responseJson.data)
  } catch (error) {
    console.error('Error: ', error)
  }
}

const setBook = (data) => {
  title.value = data.title
  author.value = data.author
  description.value = data.description
  publishedAt.value = data.published_at
  status.value = data.status
}

const updateBook = async (data) => {
  const myHeaders = new Headers()
  myHeaders.append('Content-Type', 'application/json')
  const endPoint = `http://127.0.0.1:3000/book/${id.value}`

  const requestOptions = {
    method: 'PUT',
    headers: myHeaders,
    body: data,
  }

  try {
    const response = await fetch(endPoint, requestOptions)

    if (!response.ok) {
      throw new Error(`Cannot update book, ID: ${id.value}`)
    }

    const responseJson = await response.json()

    console.log('Response: ', responseJson)

    alert(responseJson.status)
  } catch (error) {
    console.error(error)
  }
}

const preparePayload = () => {
  const payload = JSON.stringify({
    title: title.value,
    author: author.value,
    description: description.value,
    published_at: convertDate(publishedAt.value),
    status: status.value,
  })

  return payload
}

const onSubmit = () => {
  const payload = preparePayload()

  updateBook(payload)

  router.push('/')
}

const convertDate = (date) => {
  const newDate = dayjs(date).toISOString()

  return newDate
}

fetchBook()
</script>

<style></style>
