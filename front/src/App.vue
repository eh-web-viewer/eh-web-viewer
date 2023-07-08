<template>
  
  <pre id="scope" style="display: none; position: fixed; left: 0px; max-width: 500px; text-align: left; overflow: scroll; max-height: 100%;">
    <button @click="console.log(reactives); console.log(route.fullPath)"></button>
    <br>
    {{ reactives.scope }}
    <br>
    reactives.indexMetaData: {{ reactives.indexMetaData }}
    <br>
    reactives.galleryMetaData: {{ reactives.galleryMetaData }}
  </pre>  
  
  <!-- <header> -->
  <header style="display:none">
    <span><router-link :to="{path: '/'}">home</router-link></span>
    <span><router-link :to="{path: '/g/2597892/20f9db69dd/'}">gallery</router-link></span>
    <span><router-link :to="{path: '/s/b028d14f3d/2599914-1'}">image</router-link></span>
  </header>

  <!-- keep-alive not tested -->
  <!-- <keep-alive> -->
    <router-view />    
  <!-- </keep-alive> -->
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { RouterView, RouterLink, useRoute, useRouter } from 'vue-router'
const route = useRoute()
const router = useRouter()

import { Image, reactives } from '@/functions/store'


onMounted(async () => {

  reactives.galleryImageLists = ref<Image[][]>([])
  await router.isReady()
  // fetch home page
  // do not remove this for it do the first time 
  if (route.fullPath.startsWith('/g')){
    console.log("App.vue: renewGallery")
    // renewGallery()
  }else{
    // fetchIndex(route.fullPath)
    //   .then((index) => {
    //     const indexMetaData = reactives.indexMetaData
    //     const indexPreviewLists = reactives.indexPreviewLists
    //     indexMetaData.query = index.query
    //     indexMetaData.nextPage = index.nextPage
    //     indexMetaData.results = index.results

    //     indexPreviewLists.length = 1
    //     console.log(index.galleries)
    //     indexPreviewLists[0] = index.galleries.map(x => {
    //       return {
    //         url: x.url,
    //         preview: x.preview,
    //         title: x.title,
    //         category: x.category,
    //         pages: x.pages,
    //         seed: x.seeds,
    //       } as GallerySummary
    //     })
    //   })
  }
})

</script>