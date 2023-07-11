<template>
  <div class="search-box" :class="{ 'show': showSearchBox || showOptions }">
    <div class="search-box-input-continer">
      <input 
        class="search-box-input" 
        :placeholder="placeholder" 
        @focus="expandOptions" 
        @blur="undefined" 
        v-model="inputValue"
        @keydown.enter="search"
      >
      <button 
        class="search-box-clear-button" 
        @click="clickButton"
      >
        🔍
      </button>
    </div>
    <div v-show="showOptions" class="search-box-options">
      <!-- 更多选项内容 -->
      <div v-for="category in categories"
        :key="category" 
        class="search-box-catagory" 
        @click="trigger(category)"
      >
        <category-label 
          :category="category" 
          :selected="switcher[category] ?? true"
        ></category-label>
      </div>
      <div class="search-box-catagory">
        <label style="font-size: 24px; color: white;">
          <input type="checkbox" v-model="searchExpunged" style="height: 24px; width: 24px;">
          search expunged
        </label>
      </div>
      <div class="search-box-catagory" style="font-size: 24px; color: white;">
        min score:
        <select v-model="minScore" style="font-size: 24px; color: white;">
          <option value="">Not selected</option>
          <option value="2">2</option>
          <option value="3">3</option>
          <option value="4">4</option>
          <option value="5">5</option>
        </select>
      </div>


    </div>
  </div>
  <div 
    :class="{ 'dummy-show':showOptions }" 
    @click="collapseOptions"
  ></div>
  <div 
    :class="{ 'dummy':true, 'show':false }" 
    @click="collapseOptions"
  ></div>
</template>

<script setup lang="ts">
// import { fchmod } from 'fs';
import CategoryLabel from './CategoryLabel.vue';

import { 
  ref,
  onMounted,
  reactive,
onBeforeUnmount, 
} from 'vue';
import { 
  useRoute, useRouter,
} from 'vue-router'
const route = useRoute()
const router = useRouter()

const props = defineProps<{
  showSearchBox: boolean
}>()
console.log(props)
const emit = defineEmits(['submit'])


const categories = [
  'doujinshi'   ,
  'manga'       ,
  'artist CG'   ,
  'game CG'     ,
  'western'      ,
  'non-h'       ,
  'image set'   ,
  'cosplay'     ,
  'asian porn'  ,
  'misc'        ,
]

const placeholder = ref("搜索")
const showOptions = ref(false)
const inputValue = ref("")
const switcher = reactive({
  'doujinshi'  : (true),
  'manga'      : (true),
  'artist CG'  : (true),
  'game CG'    : (true),
  'western'     : (true),
  'non-h'      : (true),
  'image set'  : (true),
  'cosplay'    : (true),
  'asian porn' : (true),
  'misc'       : (true),
} as Record<string, boolean>) 
const searchExpunged = ref(false)
const minScore = ref("")

function trigger(category: string) {
  // console.log("onclick", category, switcher) // seems well
  switcher[category] = !switcher[category] 
}

function expandOptions() {
  showOptions.value = true
  document.body.classList.add('scroll-lock');
  router.push(route.fullPath+"#search")
}
function collapseOptions() {
  showOptions.value = false
  document.body.classList.remove('scroll-lock');
}
function clickButton() {
  search()
  // not used
  // console.log(inputValue.value)
  // if (inputValue.value !== '') {
  //   // do search
  // } 
  // collapseOptions()
}
const cats = ['misc','doujinshi','manga','artist CG','game CG','image set','cosplay','asian porn','non-h','western']
function search() {
  let searchParams = []
  let fCats = 0
  for (let i=0; i<10; i++) {
    if (!switcher[cats[i]])
      fCats += (1<<i)
  }
  if (fCats !== 0) searchParams.push(['f_cats',fCats.toString()])

  if (minScore.value !== '' || searchExpunged.value)
    searchParams.push(['advsearch', '1'])
  if (searchExpunged.value) searchParams.push(['f_sh', 'on'])
  if (minScore.value !== '') searchParams.push(['f_srdd', minScore.value])

  if (inputValue.value !== '') searchParams.push(['f_search', inputValue.value])

  const searchString = "/?" + (new URLSearchParams(searchParams)).toString()
  emit('submit', searchString)
}

onMounted(async () => {
  console.log("onMounted", "SearchBar")
  await router.isReady() // use this or will get '/' only
  window.addEventListener('popstate', collapseOptions);
})

onBeforeUnmount(() => {
  window.removeEventListener('popstate', collapseOptions);
})

</script>