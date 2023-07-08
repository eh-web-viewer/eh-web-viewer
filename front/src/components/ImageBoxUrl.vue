<template>
  <image-box :image="image" v-if="(typeof image !== 'undefined')"></image-box>  
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import ImageBox from './ImageBox.vue';
import { IImage, fetchImage } from '@/functions/api';

const props = defineProps<{
  url : string,
}>()

const image = ref<IImage>()

onMounted(async () => {
  const newImage = await fetchImage(props.url)
  image.value = newImage
})
</script>