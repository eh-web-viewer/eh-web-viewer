ImageBox:
show the image on exhentai.org/s/:key/:id 
params:
  image: object that get on exhentai.org/s/:key/:id, will show the image and request the altImage when failed
  src: if image is not set, show this.

<template>
  <img :src="imgSrc" :alt="imgSrc" @error="imgError" @load="imgLoaded" @loadstart="imgLoadStart" >
</template>

<script setup lang="ts">
import { 
  ref,
  onMounted,
} from 'vue';
// not used
// import { useRoute, useRouter } from 'vue-router'
// const route = useRoute()
// const router = useRouter()
import { fetchImage, IImage } from '@/functions/api'
import { findPath } from '@/functions/utils';

const props = defineProps<{
  src?: string,
  image?: IImage,
  // url?: string,
}>()

const imgSrc = ref(props.image?.image ?? props.src ?? "/favicon.ico")
const image = ref(props.image)

const timeout = 5000
let timeoutId : NodeJS.Timeout

// dunno why but works
async function imgError() { // TODO: shorten the timeout limit
  if (typeof image.value === 'undefined') return
  if (image.value.image === imgSrc.value) {    
    const nextImage = await fetchImage(findPath(image.value.query) + "?nl=" + image.value.altQuery)
    image.value = nextImage
    imgSrc.value = nextImage.image
  } else {
    imgSrc.value = image.value.image
  }
  // origin codes
  // nextLink = nextLink ?? props.image?.altQuery
  // // to next Image
  // console.log("load image error and try next link")
  // // try next link
  // await router.isReady(); // use this or will get '/' only
  // const path = props.image?.query ?? findPath(route.fullPath)
  // const nextImage = await fetchImage(path + "?nl=" + nextLink)
  // imgSrc.value = nextImage.image
  // nextLink = nextImage.altQuery
  // // reset the timer (not tested)
  timeoutId = setTimeout(() => {
    imgError()
  }, timeout);
}
// not tested
function imgLoaded() {
  // cancel the timer
  clearTimeout(timeoutId);
}
function imgLoadStart() {
  timeoutId = setTimeout(() => {
    imgError()
  }, timeout);
}

onMounted(() => {
  // console.log(imgSrc) // seems well
  // console.log(props.image) 
  // not tested
})


</script>