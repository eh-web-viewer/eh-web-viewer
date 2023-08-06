GalleryInfo:
show the infomation that get from exhentai.org/g/:id/:key
support the one get from index page.
params:
  gallery: object that indicated a gallery
<template>

  <div class="gallery-info-continer">
    <div class="gallery-info-cover-container" style="flex: 1; height: 100%; display: flex; align-items: center; justify-content: center;">      
      <!-- <img class="gallery-info-cover" :src="src" :alt="src" loading="lazy" style="max-height: 100%; max-width: 100%; width: auto; height: auto"> -->
      <image-box v-if="(typeof image === undefined)" :src="src" ></image-box>
      <image-box v-else :image="image"></image-box>
    </div>
    <div class="gallery-info-data" style="height: 100%; flex: 2; display: flex; position: relative; flex-direction: column; max-width: 70%; overflow:hidden">      
      <div class="gallery-info-data-title">
        {{ gallery.title }}
      </div>

      <div class="gallery-info-data-catagory">
        <category-label :category="gallery.category" :selected="true"></category-label>  
      </div>      

      <div class="gallery-info-data-score" style="display: none;">

      </div>        

      <div class="gallery-info-data-tags">
        <div class="gallery-info-data-tag-group" v-for="(valueArr,namespace) in gallery.tags">
          <span class="gallery-info-tag-element"><b>{{ namespace }}</b></span>
          <span class="gallery-info-tag-element" v-for="value in valueArr">
            {{ value }}
          </span>         
        </div>        
      </div>
      <div style="height:24px"></div>
      <div class="gallery-info-data-bottom" style="position: absolute; bottom:0; right:0; white-space: nowrap;">
        {{ gallery.date }} 
        {{ gallery.pages }}
      </div>
      <div style="display: none;">
        {{ gallery }}
      </div>
    </div>
  </div>

</template>


<script lang="ts" setup>
// components
import ImageBox from '@/components/ImageBox.vue';
import CategoryLabel from "@/components/CategoryLabel.vue";

import { onMounted, ref } from 'vue'
import { IGallery, IImage, fetchGallery, fetchImage } from '@/functions/api';

// what happened if props are changed.
const props = defineProps<{
  gallery : IGallery,
}>()

const gallery = ref(props.gallery)
const image = ref<IImage>()
const src = ref(props.gallery.preview.replace(/s.exhentai.org/g, 's-ex.moonchan.xyz'))

onMounted(() => {
  if (props.gallery.query !== "") return
  fetchGallery(props.gallery.url)
  .then((newGallery) => {
    // console.log(newGallery)
    gallery.value = newGallery
    if (newGallery.urls.length > 0) {
      fetchImage(newGallery.urls[0])
      .then((newImage) => {
        // console.log(newImage)
        image.value = newImage
      })
    }
  })
})
</script>