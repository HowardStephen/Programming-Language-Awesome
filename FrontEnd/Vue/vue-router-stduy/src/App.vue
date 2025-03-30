<template>
  <div class="container">
    <h2 class="title">{{ title }}</h2>
    <div class="navigation">
      <RouterLink
        v-for="(navObj, index) in navList"
        :key="index"
        :to="navObj.href"
        :class="beClicked"
        >{{ navObj.title }}</RouterLink
      >
    </div>
    <div class="show">
      <RouterView></RouterView>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'

import { RouterView, RouterLink } from 'vue-router'

const navList = reactive([
  { href: '/home', title: 'home' },
  { href: '/topic', title: 'topic' },
  { href: '/about', title: 'about' },
])

const title = ref("henry's blog")
</script>

<style scoped lang="scss">
.container {
  .title {
    width: 80%;

    padding: 20px;
    margin: 20px auto;

    background-color: black;
    border-radius: 10px;
    box-shadow: 0 0 5px black;

    color: white;
    font-size: 30px;
    text-align: center;
  }

  .navigation {
    display: flex;
    align-items: baseline;

    a {
      font-size: 20px;
      color: white;
      background-color: black;
      border-radius: 10px;
      padding: 10px 30px;
      margin: 0 auto;

      &:hover {
        background-color: cornflowerblue;
      }
    }
  }

  .show {
    width: 80%;
    height: 600px;

    margin: 20px auto;
    border-radius: 20px;
    border: 2px solid black;
  }
}
.beClicked {
  background-color: darkgreen;
}
</style>
