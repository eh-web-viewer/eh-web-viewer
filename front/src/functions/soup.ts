// general store.

// not used. 
// class NotLRU<K> {
//   maxSize = 5
//   queue: K[] = []
//   counter = new Map<K, number>()
  
//   constructor(maxSize: number) {
//     this.maxSize = maxSize
//   }

//   put(key: K): K|undefined {
//     this.queue.push(key)
//     // add cnt
//     const cnt = this.counter.get(key) ?? 0
//     this.counter.set(key, cnt+1)

//     if (this.queue.length > this.maxSize) {
//       const k = this.queue.shift()!
//       const c = this.counter.get(k) ?? 0
//       if (c == 1) {
//         this.counter.delete(k)
//         return k // remove from 
//       } else {
//         this.counter.set(k, c-1)
//       }
//     }
//   } 
// }

// type MyFunctionType

// not used
// type Entry<V> = {
//   status: string; 
//   value: V;
// };
// not used
// function covToEntry<V>(status:string, value:V): Entry<V> {
//   const entry = { status, value };
//   return entry;
// }

// Soup
// for cache something
// fetch([url]).then(r => {...}).catch(r => {...})
// change to 
// someSoup = new Soup<string, Response>(fetch)
// someSoup.get([url]).then
type ResolveFunc<V> = (value: V | PromiseLike<V>) => void
type RejectFunc = (reason?: any) => void
type Executor<V> = [ResolveFunc<V>, RejectFunc]

const MISS = 'miss'
const DONE = 'done'
const FETCHING = 'fetching'
const ERROR = 'error'

class Soup<K,V> {
  // run this method when data not in cache
  getter: (key:K) => Promise<V>;
  private _cache: Map<K,V> = new Map<K, V>() 
  private _status: Map<K, string> = new Map<K, string>()
  // resolveCallbacks: Map<K, ResolveFunc<V>[]> = new Map<K, ResolveFunc<V>[]>()
  // rejectCallbacks: Map<K, RejectFunc[]> = new Map<K, RejectFunc[]>()
  private _executors = new Map<K, Executor<V>[]>()

  constructor(
    // the raw getter function
    getter : (key:K) => Promise<V>,
    initalMap?: Map<K,V>,
  ){
    this.getter = getter
    initalMap?.forEach((v,k) => {
      this._status.set(k, DONE)
      this._cache.set(k, v)
    })
  }

  // status
  // miss => fetching => done or
  // miss => fetching => error
  private _getStatus(key: K): string {
    const status = this._status.get(key)
    if (typeof status === 'undefined') 
      return MISS
    else 
      return status
  }
  private  _setStatus(key: K, status: string) {
    this._status.set(key, status)
  }

  // cache
  private  _hasCache(key: K): boolean {
    return this._cache.has(key)
  }
  private  _getCache(key: K): V | undefined {
    return this._cache.get(key)
  }
  private  _setCache(key: K, value: V) {
    this._cache.set(key, value)
  }

  // executors
  private  _addExecutor(key: K, resolve:ResolveFunc<V>, reject:RejectFunc) {
    if (!this._executors.has(key)) {
      this._executors.set(key, [])
    }
    this._executors.get(key)!.push([resolve, reject])
  }
  private  _delExecutor(key: K) {
    this._executors.delete(key)
  }
  private _getExecutor(key: K): Executor<V>[] {
    const arr = this._executors.get(key)
    if (typeof arr === 'undefined'){
      return []
    }
    this._delExecutor(key)
    return arr
  }

  //
  delAll(key: K) {
    this._status.delete(key)
    this._cache.delete(key)
    this._getExecutor(key).forEach((executor) => {
      executor[1]("deleted")
    })
    this._delExecutor(key)
  }

  // core function
  forceRenew(key: K) {
    this._setStatus(key, FETCHING)
    // this is the original getter
    this.getter(key) // async
    .then((value) => {
      this._setCache(key, value)
      this._setStatus(key, DONE)
      this._getExecutor(key).forEach((executor) => {
        const resolve = executor[0]
        resolve(value)
      })
      this._delExecutor(key)
    })
    .catch((reason) => {
      this._setStatus(key, ERROR)
      console.log(reason)
      this._getExecutor(key).forEach((executor) => {
        const reject = executor[1]
        reject(reason)
      })
      this._delExecutor(key)
    })
  }
  // core function
  renew(key: K) {
    const status = this._getStatus(key)
    if (status === MISS || status === ERROR) {
      this.forceRenew(key)
    } else { // fetching or done
      return
    }
  }

  patch(key: K, value: V) {
    this._setCache(key, value)
    this._setStatus(key, DONE)
  }
  
  // async getWithCache(key: K, tryInterval?: number): V {
  //   let value = this.getCache(key)
  //   if (typeof value === 'undefined') {
  //     this.addCache(key)      
  //   }
  //   return value
  // }

  // add call back function
  private _addCallbacks(key: K, resolve: ResolveFunc<V>, reject: RejectFunc) {
    if (this._getStatus(key) === DONE && this._hasCache(key)) {
      // if done, just return what in cache
      resolve(this._getCache(key)!)
    } else {
      // addCache will do nothing when fetching or done
      this.renew(key)
      // if fetching, then add Executor that runs later
      this._addExecutor(key, resolve, reject)
    }
  }

  // interface 
  // usage: someSoup.get(key).then((v)=>{...}).catch(r=>{...})
  get(key: K): Promise<V> {
    
    return new Promise((resolve,reject) => {
      this._addCallbacks(key, resolve, reject)
      // original usage
      // const v = this.getCacheValue(key)
      // if (typeof v === 'undefined') {
      //   reject("undefined")
      // }else{
      //   resolve(v)
      // }
    })
  }

}

export { Soup }

// usage:
// function fetchWithLog(url: string): Promise<string> {
//   console.log("fetching "+ url)
//   return fetch(url).then(r => r.text())
// }
// someSoup.get("https://moonchan.xyz/api/cookie")
// .then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("https://moonchan.xyz/api/cookie")
// .then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("https://moonchan.xyz/api/cookie")
// .then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("http://127.0.0.1/run.sh")
// .then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("http://127.0.0.1/run.sh")
// .then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("http://127.0.0.1/run.sh")
// .then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("http://127.0.0.1/run.sh")
// .then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("https://moonchan.xyz/api/cookie")
// .then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("https://moonchan.xyz/api/cookie")
// .then(t => console.log(t))
// .catch(r => console.log(r))



// in this example body can't be read more than once so it may error
// function fetchWithLog(url: string): Promise<Response> {
//   console.log("fetching "+ url)
//   return fetch(url)
// }

// const someSoup = new Soup<string, Response>(fetchWithLog)
// someSoup.get("https://moonchan.xyz/api/cookie")
// .then(r => r.text()).then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("https://moonchan.xyz/api/cookie")
// .then(r => r.text()).then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("https://moonchan.xyz/api/cookie")
// .then(r => r.text()).then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("http://127.0.0.1/run.sh")
// .then(r => r.text()).then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("http://127.0.0.1/run.sh")
// .then(r => r.text()).then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("http://127.0.0.1/run.sh")
// .then(r => r.text()).then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("http://127.0.0.1/run.sh")
// .then(r => r.text()).then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("https://moonchan.xyz/api/cookie")
// .then(r => r.text()).then(t => console.log(t))
// .catch(r => console.log(r))
// someSoup.get("https://moonchan.xyz/api/cookie")
// .then(r => r.text()).then(t => console.log(t))
// .catch(r => console.log(r))



// there should be some 

// function delay(milliseconds: number): Promise<void> {
//   return new Promise<void>((resolve) => {
//     setTimeout(resolve, milliseconds);
//   });
// }
// // 函数实现，参数 delay 单位 毫秒 ；
// function sleep(delay: number) {
//   const start = (new Date()).getTime();
//   while ((new Date()).getTime() - start < delay) {
//       // 使用  continue 实现；
//       continue; 
//   }
// }
// // 调用方法，同步执行，阻塞后续程序的执行；
// let r : any
// function myFunction() {
//   return new Promise((resolve, reject) => {
//     // Perform asynchronous operation or any logic here
//     r = reject
//     // if (true) {
//     //   resolve([1,2,3]);
//     // } else {
//     //   reject("Error occurred");
//     // }
//   });
// }
// myFunction()
//   .then((result) => {
//     console.log(result);
//     // Handle the resolved value
//   })
//   .catch(error => {
//     console.error(error);
//     // Handle the rejected value or error
//   });
//   console.log("213")
// for(let i=0; i<20; i++){
//   sleep(1000*60*1);
//   console.log("213")
// }

// r("som")
// r("som")


//-------------------------
// const myMap = new Map();
// myMap.set("key1", "value1");
// myMap.set("key2", "value2");
// myMap.set("key3", "value3");

// myMap.forEach((v,k) => {
//   console.log(k,v)
// })

// > key1 value1
// > key2 value2
// > key3 value3
