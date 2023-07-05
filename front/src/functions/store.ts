import {
   ref, reactive,
   Ref,
} from 'vue'
// import { RouteLocationNormalizedLoaded, Router } from 'vue-router'
// import { useRoute, useRouter } from 'vue-router'
import { IImage } from "@/functions/api"

class Store<T> {
  private _cnt: number = 0
  private _value: T
  private _refMap: Record<string, Ref<T>> // nullptr
  
  constructor(value: T) {
    this._value = value
    this._refMap = {}
  }

  get value(): T {
    return this._value  
  }

  set value(value: T) {
    this._value = value
    for (const k in this._refMap) {
      this._refMap[k].value = value
    }
  }

  addSubscribe(sub: Ref<T>): () => void {
    // get count
    const k = this._cnt
    this._cnt++;
    // add sub to map
    this._refMap[k] = sub
    // return canceller
    return () => {
      delete this._refMap[k]
    }
  }

}

// interface IStore<T> {
//   value: T
// }

const recoder: Record<string, Store<any>> = {
  test : new Store<number[]>([1]),
}


type Image = {
  preview: string
  url: string
  image: string
}

type GallerySummary = {
  query: string
  url: string
  preview: string
  title: string

  originalTitle: string // 2
  category: string 
  pages: string
  uploader: string
  rate: string
  createAt: string
  seed: string

  tags: Record<string, string[]> // 2
}

const referances: Record<string, Ref<any>> = {
  // 'indexSummaryLists' : Ref<I>([])
  // 'indexSummaryLists' : [] as Ref<Image[][]>,
}



const reactives: Record<string, any> = {
  'scope' : reactive({} as Record<string, any>),
  'indexMetaData' : reactive({
    query : "",
    nextPage : "",
  }), 
  'indexPreviewLists' : reactive<GallerySummary[][]>([]),
  'galleryMetaData' : reactive({
    query : "",
    lastQuery : "",
    
    title : "",
    originalTitle: "",
    tags : {} as Record<string, string[]>,

    preview: "",
    category: "",
    pages: "",
  }) , 
  'galleryImageLists' : reactive<Image[][]>([]),
}

reactives.test = reactive({a:1,b:""})

// const route = useRoute()
// const router = useRouter()
// const route = {
//   route: null as any as RouteLocationNormalizedLoaded,
//   router: null as any as Router,
// }

const store = {
  indexData : {
    query : "invalid",
    nextPage : "",
    showingLists : reactive<GallerySummary[]>([]),
    showableLists : reactive<GallerySummary[]>([]),
  },
  galleryData : {
    query : "invalid",
    showingLists : reactive<Image[]>([]),
    showableLists : reactive<Image[]>([]),
  },
  imageData : {
    query : "invalid",
    preloadLength : 5,
    // imageLists : reactive<Image[]>([]), // Proxy(Array)
    imageRecords : new Map<number, IImage>(), // key: query, value: the Image of that query
    fetchedRecords : new Set<string>(),
    // image : ref<IImage>({
    //   query : "undefined",
    //   galleryQuery : "",
    //   nextPageQuery : "",
    //   prevPageQuery : "",
    //   image : "",
    //   altQuery : "",
    //   error : "",
    // } as IImage), // not used
  },
}

export type { Image, GallerySummary }
export { recoder, store, referances, reactives } 