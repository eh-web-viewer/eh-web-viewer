<template>
  <button @click="console.log(src)"></button>
  <input v-model="src" />
  <EhwvImage :src="src" @error="onError"></EhwvImage>
</template>

<script setup lang="ts">
  import { ref } from "vue"
  import EhwvImage from "../components/EhwvImage.vue"
  const src = ref("");
  function onError(){
    console.log("error")
  }
</script>