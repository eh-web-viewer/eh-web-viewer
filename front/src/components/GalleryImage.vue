not used

<template>

  <img 
    :id="id"
    :src="src" 
    :alt="src"
    :style="{width: '100%', height : minHeight}"
    :loading="lazy"
    @load="onImageLoad"
  >

</template>

<script setup lang="ts">
// vue-core
import { ref, onBeforeMount, onMounted } from 'vue'
const src = ref("https://moonchan.xyz/favicon.ico")
const id = ref("")
const lazy = ref("eager")
const minHeight = ref("1080px")
// functions
import { Image } from '@/functions/store';
import { fetchImage } from '@/functions/api';
const props = defineProps<{
  image : Image,
}>()

// console.log(props.image)
function getIdFromUrl(url: string): string{
  const arr = url.split('/')
  return "_" + arr[arr.length-1]
}

onBeforeMount(() => {
  console.log(props.image)
  const image = props.image
  id.value = getIdFromUrl(image.url)
})

// const options = {
//   root: null,
//   rootMargin: '0px',
// }
const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      // console.log(getIdFromUrl(props.image.url), ":" , entry.intersectionRatio)
      const remainingDistance = entry.boundingClientRect.top - entry.rootBounds!.top;
      console.log(getIdFromUrl(props.image.url), ":", `Remaining distance to initial entry: ${remainingDistance}px`);

      if ( props.image.preview.replace(/s.exhentai.org/g, "s-ex.moonchan.xyz") !== src.value )
        return 
      // if (entry.intersectionRatio > 0) {
      if (remainingDistance < 6*1080) { // will affect initial images loaded
        // minHeight.value = ""
        fetchImage(props.image.url.substring('https://exhentai.org'.length))
          .then((img) => {
            // console.log(img, src)
            src.value = img.image || src.value
            lazy.value = ""
          })
      }
    });
  }, {
    root: null,
    rootMargin: '0px 0px 8000px', // will affect how much images load before, and this is two times when use 'Xpx'
    threshold: [0, 0.2, 0.4, 0.6, 0.8, 1]
  }
)

function onImageLoad() {
  minHeight.value = ""
}

onMounted(async () => {
  // await new Promise(resolve => setTimeout(resolve, 100));
  const imageOverride = props.image.image.replace(/s.exhentai.org/g, "s-ex.moonchan.xyz")
  const previewOverride = props.image.preview.replace(/s.exhentai.org/g, "s-ex.moonchan.xyz")
  src.value  = imageOverride || previewOverride // preview only

  console.log("GalleryImage: "+"img#" + getIdFromUrl(props.image.url)+" onMounted")
  if ( props.image.preview.replace(/s.exhentai.org/g, "s-ex.moonchan.xyz") !== src.value )
    return
  // console.log("entry")
  const entry = document.querySelector("img#" + getIdFromUrl(props.image.url))
  if (entry !== null)
    setTimeout( () => observer.observe(entry), 500);
})



</script>