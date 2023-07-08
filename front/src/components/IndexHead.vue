not used

<template>
  <input v-model="query" @keyup.enter="search">
  {{ query }}
  <br>
</template>

<script setup lang="ts">

  import { useRoute, useRouter } from 'vue-router'
  const route = useRoute()
  const router = useRouter()

  import { ref } from 'vue';
  const query = ref(route.fullPath.substring(2));

  function search() {
    router.push("/" + query.value)  
  }
// watch(query, (oldValue, newValue) =>{
//   console.log(oldValue, newValue)
// })
</script>