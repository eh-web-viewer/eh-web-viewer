<template>
  <!-- <index-head></index-head> -->
  {{ metaData }}
  <template v-for="previewList in previewLists" :key="previewList">
    <template v-for="summary in previewList" :key="summary">
      <!-- <div> -->
        <!-- {{ summary.title }} -->
      <!-- </div> -->
      <div>
        <gallery-preview :summary="summary" />
      </div>
    </template>
  </template>
  <div id="index-hook"></div>
</template>

<script setup lang="ts">
// components
import GalleryPreview from "@/components/GalleryPreview.vue";
// import IndexHead from "@/components/IndexHead.vue";
// vue-core
import { onBeforeUpdate, onBeforeMount, onMounted } from "vue"
// vue-route
import { useRoute, useRouter } from 'vue-router'
const route = useRoute()
const router = useRouter()
// api
import { fetchIndex } from "@/functions/api";
// store
import { GallerySummary } from '@/functions/store'
import { reactives } from '@/functions/store'
const metaData = reactives.indexMetaData
const previewLists = reactives.indexPreviewLists

let lastPath : string

function updateIndex(path: string) {
  if (lastPath !== path) {
    // on change
    reactives.scope.path = path // it will give '/' only, but supposed to have query.  // for debug
    // renew IndexView
    fetchIndex(path) // remove the prefix '/'
    .then((index) => {
      const indexMetaData = reactives.indexMetaData
      const indexPreviewLists = reactives.indexPreviewLists
      indexMetaData.query = index.query
      indexMetaData.nextPage = index.nextPage
      indexMetaData.results = index.results

      indexPreviewLists.length = 1
      console.log(index.galleries)
      indexPreviewLists[0] = index.galleries.map(x => {
        return {
          url: x.url,
          preview: x.preview,
          title: x.title,
          category: x.category,
          pages: x.pages,
          seed: x.seeds,
        } as GallerySummary
      })
    })
  }
}

function updateIndexNextPage() {
  const regex = /https:\/\/exhentai.org(\/.*)";/
  const nextPage = reactives.indexMetaData.nextPage.match(regex)[1] || "/"
  fetchIndex(nextPage) // remove the prefix '/'
  .then((index) => {
    const indexMetaData = reactives.indexMetaData
    const indexPreviewLists = reactives.indexPreviewLists
    indexMetaData.query = index.query
    indexMetaData.nextPage = index.nextPage
    indexMetaData.results = index.results

    console.log(index.galleries)
    indexPreviewLists.push(index.galleries.map(x => {
      return {
        url: x.url,
        preview: x.preview,
        title: x.title,
        category: x.category,
        pages: x.pages,
        seed: x.seeds,
      } as GallerySummary
    }))
  })
}

const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      // console.log(getIdFromUrl(props.image.url), ":" , entry.intersectionRatio)
      const remainingDistance = entry.boundingClientRect.top - entry.rootBounds!.top;
      console.log("index:", `Remaining distance to initial entry: ${remainingDistance}px`);

      if (remainingDistance < 1080) { // will affect initial images loaded
        updateIndexNextPage()
      }
    });
  }, {
    root: null,
    rootMargin: '0px 0px 0px', // will affect how much images load before, and this is two times when use 'Xpx'
    threshold: [0, 0.2, 0.4, 0.6, 0.8, 1]
  }
)


onBeforeMount(async () => {
  document.title = "EhWebViewer"
  console.log("indexView: onBeforeMount")
  await router.isReady() // use this or will get '/' only
  updateIndex(route.fullPath)
  lastPath = route.fullPath
})

onMounted(() => {
  const entry = document.querySelector("div#index-hook")
  console.log(entry)
  if (entry !== null)
    setTimeout( () => observer.observe(entry), 500);
})

onBeforeUpdate(async () => {
  console.log("indexView: onBeforeUpdate")
  await router.isReady() // use this or will get '/' only
  updateIndex(route.fullPath)
  lastPath = route.fullPath
})

</script>
  