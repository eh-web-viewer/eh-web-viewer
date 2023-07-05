async function myFetch(query: string): Promise<any> {
  if (!query.startsWith("/")) {
    query = "/"+query
  }
  const respJson = await fetch("/api" + query)
                        .then(response => response.json())
  return respJson
}


interface IGallerySummary {
  query         : string   ;   // the original query string
  preview       : string   ;   // the cover url
  title         : string   ;   // title in english
  originTitle   : string   ;   // title in japanese
  tags          : string[] ;   // language:chinese ...
  category      : string   ;   // doujin, manga ...
  rate          : string   ;   // stars from 1 to 10
  date          : string   ;   // publish date
  pages         : string   ;   // how many pages are there in gallery
  seeds         : string   ;   // url to the seeds link
  url           : string   ;   // url to the gallery page
  error         : string   ;   // error
} 

function parseGallerySummary(obj: any): IGallerySummary {
  // query, preview, title, originTitle
  const query : string = obj.query!
  const preview : string = obj.preview!
  const title : string = obj.title!
  const originTitle : string = obj.origin_title!
  // TODO: tags
  const tags : string[] = ["todo"]
  // caeagory
  const category : string = obj.category!
  // TODO: rage
  const rate : string = obj.rate!
  // date
  const date : string = obj.date!
  // results like "1234 pages."
  const str = obj.pages!
  const regex = /[\d]+/g 
  const pages = str.match(regex)?.join("") ?? ""
  // galleries
  const seeds : string = obj.seeds!
  const url : string = obj.url!
  // error
  const error : string = obj.error!

  const gallerySummeryObj : IGallerySummary = {
    query       : query        ,
    preview     : preview      ,
    title       : title        ,
    originTitle : originTitle  ,
    tags        : tags         ,
    category    : category     ,
    rate        : rate         ,
    date        : date         ,
    pages       : pages        ,
    seeds       : seeds        ,
    url         : url          ,
    error       : error        ,
  } 
  return gallerySummeryObj
}

async function fetchGallerySummmary(query: string): Promise<IGallerySummary>  {
  return await myFetch(query).then((obj) => {
    return parseGallerySummary(obj)
  })
}

interface IIndex {
	query       : string            ;  // which query it executes
	results     : string            ;  // how many results
	nextPage   : string            ;  // url to next page
	prevPage   : string            ;  // url to previous page
	galleries   : IGallerySummary[] ;  // search results
	error       : string            ;  // error
}

function parseIndex(obj: any): IIndex {
  // query
  const query : string = obj.query!
  // results like "123,456 results."
  const str = obj.results!
  const regex = /[\d]+/g 
  const results = str.match(regex)?.join("") ?? ""
  // nextpage "var nexturl=\"https://exhentai.org/?next=2559818\";"
  const nextPage : string = obj.next_page!
  const prevPage : string = obj.prev_page!
  // galleries
  const galleries : IGallerySummary[] = obj.galleries!.map((item:any) => parseGallerySummary(item))
  // error
  const error : string = obj.error!

  const indexObj : IIndex = {
    query: query,
    results: results,
    nextPage: nextPage,
    prevPage: prevPage,
    galleries: galleries,
    error : error,
  }
  return indexObj
}

async function fetchIndex(query: string): Promise<IIndex> {
  return await myFetch(query).then((obj) => {
    return parseIndex(obj)
  })
}

interface IGallery {
  query   : string   ; // the original query string
  preview : string ;
  title   : string ;
  originalTitle: string; 
  tags    : string[][]; 
  category: string;
  pages : string;
  images  : string[] ; // url of previewing pics
  url     : string[] ; // url of pics
  error   : string   ; // error
}

function parseGallery(obj : any): IGallery { 
  const query = obj.query!
  const images = obj.images!.map((s:string) => s)
  const url = obj.url!.map((s:string) => s)
  const error = obj.error!
  const galleryObj: IGallery = {
    query : query,
    preview : obj.preview! , 
    title: obj.title!, 
    originalTitle : obj.origin_title!,
    tags : obj.tags!,
    category: obj.category!,
    pages: obj.pages!,
    images : images,
    url : url,
    error : error,
  }
  return galleryObj
}

async function fetchGallery(query: string): Promise<IGallery> {
  return await myFetch(query).then((obj) => {
    return parseGallery(obj)
  })
}

// IImage
function chopString(s:string, prefix:string, surfix?:string): string {
  const s1 = s.substring(prefix.length)
  if (typeof surfix === 'undefined') return s1
  const s2 = s1.substring(0, s1.length-surfix.length)
  return s2
}
interface IImage {
  query : string
  galleryQuery : string
  nextPageQuery : string
  prevPageQuery : string
  image : string
  altQuery : string
  error : string 
}
function parseImage(obj : any): IImage { 
  const query = obj.query!
  const galleryQuery = chopString(obj.gallery_page, "https://exhentai.org")
  const image = obj.image!
  const error = obj.error!
  const nextPageQuery = chopString(obj.next_page, "https://exhentai.org")
  const prevPageQuery = chopString(obj.prev_page, "https://exhentai.org")
  const altQuery = chopString(obj.next_page, "return nl('", "')")
  const imageObj: IImage = {
    query : query,
    galleryQuery: galleryQuery,
    nextPageQuery: nextPageQuery,
    prevPageQuery: prevPageQuery,
    image : image,
    altQuery: altQuery,
    error : error,
  }
  return imageObj
}
async function fetchImage(query: string): Promise<IImage>{
  return await myFetch(query).then((obj) => {
    return parseImage(obj)
  })
}

// exports
export type { IIndex, IGallery, IGallerySummary, IImage }
export { fetchIndex, fetchGallerySummmary, fetchGallery, fetchImage }