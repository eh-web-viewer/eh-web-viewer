<template>
  <button @click="()=>{console.log('image',image);console.log('imageList,route.fullPath',imageList, route.fullPath)}"></button>
  <router-link :to="{path: nextPageQuery}">    
    <div v-for="image in imageList" :key="image.query" v-show="image.query === route.fullPath">
      <!-- {{ image.query }} -->
      <ImageBox :image="image" v-if="(typeof image !== 'undefined')" ></ImageBox>
    </div>
  </router-link>
  
  <!-- <router-link :to="{path: firstPageQuery}">first page</router-link> -->
  <router-link :to="{path: prevPageQuery}"><buttom> &lt; </buttom></router-link>
  <router-link :to="{path: nextPageQuery}"><buttom> &gt; </buttom></router-link>
  <!-- <router-link :to="{path: lastPageQuery}">last page</router-link> -->
</template>

<script setup lang="ts">
// components
import ImageBox from "@/components/ImageBox.vue";

// vue-core
import { 
  ref,
  onUpdated, onMounted, 
} from "vue";

// vue-route
import { useRoute, useRouter } from "vue-router";
const route = useRoute();
const router = useRouter();

// get datas
import { store } from "@/functions/store";
import { fetchImage, IImage } from "@/functions/api";
import { loadFullPath } from '@/functions/utils'

const prevPageQuery = ref("")
const nextPageQuery = ref("")
// const firstPageQuery = ref("")
// const lastPageQuery = ref("")
const imageSrc = ref("/favicon.ico")
const imageList = ref<IImage[]>([])

let image : IImage|undefined

const data = store.imageData // this line works well as it's in hook function
// functions
// for use set
function findIndex(query:string): number{
  const arr = query.split('-')
  return parseInt(arr[arr.length - 1])
}
function hasImage(query:string): boolean {
  return data.imageRecords.has(findIndex(query))
}
function setImage(query:string, image:IImage) {
  return data.imageRecords.set(findIndex(query), image)
}
function getImage(query:string): IImage|undefined {
  return data.imageRecords.get(findIndex(query))
}
// load image item to template
async function reloadTemplate(image:IImage) {
  // console.log("reloadTemplate", image) // seems well
  // didn't outputed .what the fuck..
  // firstPageQuery.value = image.firstPageQuery
  // lastPageQuery.value = image.lastPageQuery
  prevPageQuery.value = image.prevPageQuery
  nextPageQuery.value = image.nextPageQuery
  imageSrc.value = image.image
  imageList.value = []
  let idx = findIndex(image.query)
  idx -= 3
  if (idx < 1) idx = 1
  for (let i = 0; i<data.preloadLength; i++) {
    const image = data.imageRecords.get(i+idx)
    if (typeof image !== 'undefined')
      imageList.value.push(image)
  }
}
// preload image
async function preloadImage(image: IImage) {
  const tag = new Image()
  tag.onerror = () => {
    console.log("load", image.image, "error")
  }
  tag.src = image.image
}
// preload next N images
async function nextNImages(query:string, n:number) {
  if (n < 0) return
  // try this query in records.
  // console.log(query) // has '/'
  let image :IImage|undefined
  if (hasImage(query)) {
    image = getImage(query)
  } else {
    console.log("miss", query)
    if (data.fetchedRecords.has(query)) return;
    data.fetchedRecords.add(query)
    image = await fetchImage(query)
    if (typeof image === 'undefined') { // tempory error handler
      image = await fetchImage(query)
    }
    if (typeof image !== 'undefined') {
      setImage(query, image)    
      preloadImage(image)
    }
  }
  // if not found, (1)download and put it into records and (2)preload the image.
  // try next query
  const nextQuery = image?.nextPageQuery
  // console.log(nextQuery, image?.nextPageQuery, image) // 
  if (typeof nextQuery !== 'undefined') {
    nextNImages(nextQuery, n-1)
  }
}

// preload prev N images
async function prevNImages(query:string, n:number) {
  if (n < 0) return
  // try this query in records.
  // console.log(query) // has '/'
  let image :IImage|undefined
  if (hasImage(query)) {
    image = getImage(query)
  } else {
    console.log("miss", query)
    if (data.fetchedRecords.has(query)) return;
    data.fetchedRecords.add(query)
    image = await fetchImage(query)
    if (typeof image === 'undefined') { // tempory error handler
      image = await fetchImage(query)
    }
    if (typeof image !== 'undefined') {
      setImage(query, image)    
      preloadImage(image)
    }
  }
  // if not found, (1)download and put it into records and (2)preload the image.
  // try prev query
  const prevQuery = image?.prevPageQuery
  // console.log(nextQuery, image?.nextPageQuery, image) // 
  if (typeof prevQuery !== 'undefined') {
    prevNImages(prevQuery, n-1)
  }
}

// first get into `/s/*` will only trigger onMounted function
onMounted(async () => { 
  console.log("Gallery: onMount");
  await router.isReady(); // use this or will get '/' only
  console.log(route.fullPath) // /s/:key/:id ( /s/b028d14f3d/2599914-1 )
  
  // if a new gallery
  if (route.fullPath != data.query) {
    data.imageRecords = new Map<number, IImage>() // initial a new view
    data.fetchedRecords = new Set<string>()
  }
  // debug
  console.log(data)
  data.query = route.fullPath
  console.log(data)
  // get image json data
  const query = loadFullPath(route.fullPath)
  image = getImage(query)
  if (typeof image === 'undefined') {
    console.log(query, "not in records")
    image = await fetchImage(query)
    if (typeof image !== 'undefined') {
      setImage(query, image)
    } else {
      console.log(query, "fetchefailed")
    }
  }
  reloadTemplate(image)

  // preload
  // nextNImages(query, data.preloadLength)
});

onUpdated(async () => {
  console.log("Gallery: onUpdated")
  await router.isReady()
  console.log(route.fullPath)

  // get image json data
  const query = loadFullPath(route.fullPath)
  image = getImage(query)
  if (typeof image === 'undefined') { // if not hit preloads
    console.log(query, "not in records")
    image = await fetchImage(query)
    if (typeof image !== 'undefined') {
      setImage(query, image)
    } else {
      console.log(query, "fetch failed")
    }
  }    
  reloadTemplate(image)

  // preload
  // console.log(query, data.preloadLength)
  nextNImages(query, data.preloadLength)
  prevNImages(query, data.preloadLength/2)

});

</script>
