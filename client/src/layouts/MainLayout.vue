<template>
  <q-layout view="lHh Lpr lFf" class="shadow-2 rounded-borders">
    <q-header class="bg-green-14">
      <q-toolbar>
        <router-link to="/">
          <q-avatar>
            <img src="https://i.pinimg.com/564x/a3/d1/9c/a3d19cd17a81c12be4bb7bcb2a6bd486.jpg" />
          </q-avatar>
        </router-link>
        <q-toolbar-title> Books Management </q-toolbar-title>
      </q-toolbar>
    </q-header>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>
