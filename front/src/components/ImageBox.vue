<template>
  <img :src="imgSrc" :alt="imgSrc" @error="imgError"
    style="width: 100%; height: auto;"
  >
</template>

<script setup lang="ts">
import { 
  ref,
  onMounted,
} from 'vue';
import { useRoute, useRouter } from 'vue-router'
const route = useRoute()
const router = useRouter()
import { fetchImage, IImage } from '@/functions/api'

const props = defineProps<{
  image : IImage,
}>()

const imgSrc = ref(props.image.image)

let nextLink = props.image.altQuery

function findPath(fullPath: string): string {
  const arr = fullPath.split("?")
  return arr[0]
}
async function imgError() { // TODO: shorten the timeout limit
  // to next Image
  console.log("load image error and try next link")
  // try next link
  await router.isReady(); // use this or will get '/' only
  const path = findPath(route.fullPath)
  const nextImage = await fetchImage(path + "?nl=" + nextLink)
  imgSrc.value = nextImage.image
  nextLink = nextImage.altQuery
}
onMounted(() => {
  // console.log(imgSrc) // seems well
  // console.log(props.image) 
})


</script>