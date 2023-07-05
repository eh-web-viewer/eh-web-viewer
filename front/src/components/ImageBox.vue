<template>
  <img :src="image.image" :alt="image.image" @error="imgError"
    style="width: 100%; height: auto;"
  >
</template>

<script setup lang="ts">

import { IImage } from '@/functions/api'

function imgError() {
  console.log("load image error")
}

// const props = defineProps<{
defineProps<{
  image : IImage,
}>()




</script>