<template>
  <v-app>
    <v-app-bar color="primary" dark prominent>
      <v-app-bar-nav-icon></v-app-bar-nav-icon>
      <v-toolbar-title>TodoList</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn text @click="$router.push('/')">
        <v-icon left>mdi-home</v-icon>
        首页
      </v-btn>
      <v-btn text @click="$router.push('/about')">
        <v-icon left>mdi-information</v-icon>
        关于
      </v-btn>
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>
