<template>
  <q-page class="flex justify-center">
    <div class="q-pa-md full-width" style="max-width: 400px">
      <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
        <q-input
          filled
          v-model="title"
          label="Your Title *"
          hint="Title"
          lazy-rules
          :rules="[(val) => (val && val.length > 0 && val.length < 10) || 'Please type something']"
        />

        <q-file
          v-model="file"
          counter
          filled
          use-chips
          label="Image (< 1 MB size and PNG only)"
          :rules="[(val) => val || 'Please type something']"
          :filter="checkFile"
        >
          <template v-slot:prepend>
            <q-icon name="cloud_upload" />
          </template>
        </q-file>

        <q-input
          filled
          v-model="author"
          label="Your Author *"
          hint="Author"
          lazy-rules
          :rules="[(val) => (val && val.length > 0) || 'Please type something']"
        />

        <q-input
          filled
          v-model="description"
          label="Your Description *"
          hint="Description"
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
          <q-btn label="Submit" type="submit" color="green-14" />
          <q-btn label="Reset" type="reset" color="green-14" flat class="q-ml-sm" />
        </div>
      </q-form>
    </div>
  </q-page>
</template>

<script setup>
import dayjs from 'dayjs'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
const title = ref(null)
const file = ref(null)
const author = ref(null)
const description = ref(null)
const publishedAt = ref(null)
const status = ref(null)
const router = useRouter()

const options = ['available', 'borrowed', 'reserved', 'lost']

const checkFile = (file) => {
  return file.filter((file) => file.size < 1024 * 1000 && file.type === 'image/png')
}

const onSubmit = () => {
  const payload = preparePayload()

  createBook(payload)

  // fetch('http://127.0.0.1:3000/book', requestOptions)
  //   .then((response) => response.json())
  //   .then((result) => console.log(result))
  //   .catch((error) => console.error(error))
}

const preparePayload = () => {
  const formData = new FormData()
  formData.append('title', title.value)
  formData.append('image', file.value)
  formData.append('author', author.value)
  formData.append('description', description.value)
  formData.append('published_at', convertDate(publishedAt.value))
  formData.append('status', status.value)

  return formData
}

const createBook = async (formdata) => {
  const endPoint = 'http://127.0.0.1:3000/book'

  const requestOptions = {
    method: 'POST',
    body: formdata,
  }

  try {
    const response = await fetch(endPoint, requestOptions)

    if (!response.ok) {
      throw new Error('cannot create book data')
    }

    const responseJson = await response.json()
    console.log('Response: ', responseJson)
    router.push('/')
  } catch (error) {
    console.error('Error: ', error)
  }
}

const onReset = () => {
  title.value = null
  file.value = null
  author.value = null
  description.value = null
  publishedAt.value = null
  status.value = null
}

/**
 * convert date format
 *
 * @param date
 * @return newDate
 */
const convertDate = (date) => {
  const newDate = dayjs(date).toISOString()

  return newDate
}
</script>

<style></style>
