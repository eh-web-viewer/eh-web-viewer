# eh-web-viewer/back

the backend of eh-web-viewer.

convert eh webpage to api format

mainly contain these catalogs
 
- index page
  - it contains the gallery facebooks, and the next page and previous page.
- seed page
  - view only the seed info in the gallery.
- gallery page
  - view the infomation of the gallery
- view page
  - where to read the comic, in both water fall mode and book mode (maybe should add from left or from right)

## usage

router:
- `/g/[gid]/[key]?p=[page]` will return gallery data
- `/s/[gid]/[key]?nl=[some_string]` will return image data
- `/{others}` will return index data

samples: see [samples](#samples)

## sub modules

- my_if: set ipv6 ips
- my_curl: set fetch function and client object (V6POOL mode based on *my_if*)
- exproxy: vanilla exhentai proxy, use client from *my_curl*
- api: api backend

### my_if

used for get new ips from /64 ipv6 address.

params:
- PREFIX : the ip address you have (default = [v6prefix])
- BATCH_SIZE : how many v6 ips provided at once (default = 8)

usage:
```go
my_if.SetPrefix(my_if.PREFIX)
defer my_if.Cleanup()

...

my_if.GetIPBatchAndShift() // return [BATCH_SIZE] v6 ips in `net.IP`
```

### mycurl

curl pages with some clients

usage:
```go
mycurl.SetClient(mycurl.VANILLA|mycurl.V6POOL|mycurl.INSECURE|mycurl.PROXY)  // set the mod, if V6POOL should be set, my_if should set first

mycurl.Fetch(method, url string, headers map[string]string, body io.Reader) (*http.Response, error) // fetch some web resources
mycurl.Client() // return next client should use.
```

### exproxy

exhentai proxy

usage:
see [back\exproxy\exhentai_test.go](exproxy\exhentai_test.go)

### api

core api

usage:
see [usage](#usage)

#### samples

##### index
```json
{
  "query": "",
  "results": "1,286,483 results.",
  "next_page": "var nexturl=\"https://exhentai.org/?next=2571751\";",
  "prev_page": "",
  "galleries": [
    {
      "preview": "https://s.exhentai.org/t/aa/99/aa99c7f636168a42326532eab6006c8c46a79dc5-526666-707-1000-jpg_250.jpg",
      "title": "[Artist] Hong",
      "category": "Non-H",
      "rate": "background-position:0px -21px;opacity:1",
      "date": "2023-06-04 10:53",
      "pages": "503 pages",
      "seeds": "https://exhentai.org/img/t.png",
      "url": "https://exhentai.org/g/2571776/ec49bddf1c/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/7f/87/7f87bb5c16c72ac6a660b054c50b7389a97a9d43-1865222-1600-900-png_250.jpg",
      "title": "●PIXIV● shang [32326413]",
      "category": "Misc",
      "rate": "background-position:-48px -1px;opacity:1",
      "date": "2023-06-04 10:51",
      "pages": "1504 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571775/9e8e65f34f/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/e3/60/e3608b02109ac842bbb3ef2ca449aa4d01fdcef9-5697093-2854-2160-png_250.jpg",
      "title": "(Deadbolt) Slutty Sarah on the Casting Couch [The Last of Us]",
      "category": "Misc",
      "rate": "background-position:-48px -1px;opacity:0.93333333333333",
      "date": "2023-06-04 10:48",
      "pages": "10 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571774/b3f6a7af1c/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/82/e5/82e520361853665f618cb2d9849491ead8735956-170433-1200-630-jpg_250.jpg",
      "title": "Artist | ナツ (26.12.2022 - )",
      "category": "Image Set",
      "rate": "background-position:-48px -1px;opacity:1",
      "date": "2023-06-04 10:47",
      "pages": "518 pages",
      "seeds": "https://exhentai.org/img/t.png",
      "url": "https://exhentai.org/g/2571773/914c9bc591/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/a1/e0/a1e02d1b15edd8219c1a605402321447b3ed3212-1860966-3000-2000-jpg_250.jpg",
      "title": "[lacanishu] Ana in captivity (Overwatch) [Korean]",
      "category": "Western",
      "rate": "background-position:-32px -1px;opacity:1",
      "date": "2023-06-04 10:42",
      "pages": "14 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571771/197101da7f/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/c0/31/c03150ac6aaa33cf4d9b089c9728f3f2d1903d20-239260-1280-960-jpg_250.jpg",
      "title": "Lighting Studio 1/4 WIDOWMAKER",
      "category": "Misc",
      "rate": "background-position:0px -21px;opacity:0.86666666666667",
      "date": "2023-06-04 10:37",
      "pages": "22 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571770/fe4c841e3f/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/f0/ff/f0ff11b857e69f010497a972ce2438d7859c1e3a-167475-1280-853-jpg_250.jpg",
      "title": "DT\u0026UME Studio 1/4 Yuuki Asuna",
      "category": "Misc",
      "rate": "background-position:0px -21px;opacity:1",
      "date": "2023-06-04 10:37",
      "pages": "38 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571769/22dd553976/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/0f/46/0f460d3f8030861b0184d1d9605fe8e3ec15402e-226194-1280-853-jpg_250.jpg",
      "title": "DT\u0026UME Studio 1/4 Selvaria·Bles",
      "category": "Misc",
      "rate": "background-position:0px -21px;opacity:0.8",
      "date": "2023-06-04 10:37",
      "pages": "51 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571768/ce21ea7533/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/94/45/9445b2449ae69ee6b0bb1358360f7be7bc5297ba-210128-1846-1278-jpg_250.jpg",
      "title": "Fantastic Capsule Studio - Beauty Drunken Girl",
      "category": "Misc",
      "rate": "background-position:-16px -1px;opacity:0.66666666666667",
      "date": "2023-06-04 10:37",
      "pages": "45 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571767/7532291475/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/e4/86/e486bdf70d3c10e7b69c47dd5fb0e0d3b108990b-227796-750-1280-jpg_250.jpg",
      "title": "GREEN LEAF STUDIO 1/4 Ace Sniper 'Quiet' Jing Jing",
      "category": "Misc",
      "rate": "background-position:0px -21px;opacity:0.73333333333333",
      "date": "2023-06-04 10:37",
      "pages": "38 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571766/6712fb8142/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/4f/18/4f18efc41152043a119c387514724013f786fa02-278174-959-1280-jpg_250.jpg",
      "title": "Creation Epic Studio Mirror Girl",
      "category": "Misc",
      "rate": "background-position:-16px -1px;opacity:0.73333333333333",
      "date": "2023-06-04 10:37",
      "pages": "28 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571765/dd5196b3e8/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/d3/31/d3317ae8a4cb613043abd777375950af67be5651-162719-1620-1080-jpg_250.jpg",
      "title": "Dragon Studio : Ada Wong",
      "category": "Misc",
      "rate": "background-position:0px -21px;opacity:0.66666666666667",
      "date": "2023-06-04 10:37",
      "pages": "36 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571764/209fd249b6/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/1d/58/1d58dbe39933399ed8fc346e28eecbeda3998f0b-104068-1010-1536-jpg_250.jpg",
      "title": "Asuka – Neon Genesis Evangelion (FA Studio)",
      "category": "Misc",
      "rate": "background-position:-16px -1px;opacity:0.73333333333333",
      "date": "2023-06-04 10:37",
      "pages": "39 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571763/5ba8b579b3/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/44/12/441249465b64edc34e415fde0aeed22e5f70967c-265810-1280-720-jpg_250.jpg",
      "title": "Long Studio 1/4 Yae Miko",
      "category": "Misc",
      "rate": "background-position:0px -21px;opacity:0.86666666666667",
      "date": "2023-06-04 10:37",
      "pages": "38 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571762/0421ba6041/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/06/c7/06c7619e311ada7b369a5a7728cbf0c7969e1390-348889-1280-1180-jpg_250.jpg",
      "title": "MF Studio 1/4 NieR 2B",
      "category": "Misc",
      "rate": "background-position:-16px -1px;opacity:0.66666666666667",
      "date": "2023-06-04 10:36",
      "pages": "49 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571761/39201290c4/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/2b/28/2b28f6f81826673b5300cd47b6e6dc94a6f9807c-218660-854-1280-jpg_250.jpg",
      "title": "HEBE STUDIO 1/4 'PanJinLian'",
      "category": "Misc",
      "rate": "background-position:0px-1px;opacity:0.86666666666667",
      "date": "2023-06-04 10:36",
      "pages": "42 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571760/543cdebf9d/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/58/8b/588bfae6447e10f0611d5df600118b2cfcc96530-278261-560-420-jpg_250.jpg",
      "title": "[Ellshed (Kuraido)] Futakko -Tamami- Jii hen [Chinese] [不咕鸟汉化组]",
      "category": "Artist CG",
      "rate": "background-position:-16px -1px;opacity:1",
      "date": "2023-06-04 10:36",
      "pages": "137 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571759/63616bd4b6/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/c7/5d/c75d2dd647c71da91da0cb7a8f7742fa2afde139-133734-750-1127-jpg_250.jpg",
      "title": "Shenhe - Genshin Impact (Acy Studio)",
      "category": "Misc",
      "rate": "background-position:-16px -1px;opacity:0.8",
      "date": "2023-06-04 10:36",
      "pages": "27 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571758/d72b183b14/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/ef/08/ef085a05a537aaa84731d3afd8c68efc7c1db364-90095-1280-850-jpg_250.jpg",
      "title": "Scratch Studio Albedo",
      "category": "Misc",
      "rate": "background-position:-16px -1px;opacity:0.66666666666667",
      "date": "2023-06-04 10:36",
      "pages": "37 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571757/ce51747b46/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/e8/24/e8248dd310dd3e02c9f798fb61d9840537f68ba5-138091-1280-905-jpg_250.jpg",
      "title": "TR STUDIO 1/4 Mai Shiranui",
      "category": "Misc",
      "rate": "background-position:-16px -1px;opacity:0.66666666666667",
      "date": "2023-06-04 10:36",
      "pages": "50 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571756/528489c0d9/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/e4/30/e43049a808a0250076cd1c21b669fe4be486fa38-235934-1280-798-jpg_250.jpg",
      "title": "Long Studio 1/4 Genshin Impact Ganyu",
      "category": "Misc",
      "rate": "background-position:0px -21px;opacity:0.86666666666667",
      "date": "2023-06-04 10:36",
      "pages": "45 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571755/1576e9be03/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/7a/f3/7af3f4f2150f0c2d73077d400ebd4192a1de111c-188588-1280-720-jpg_250.jpg",
      "title": "LengShi Studio 1/4 Ayanami Rei",
      "category": "Misc",
      "rate": "background-position:-16px -1px;opacity:0.66666666666667",
      "date": "2023-06-04 10:36",
      "pages": "24 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571754/c6cb263241/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/2f/37/2f37f650b2e49db057d3ede542c335d2f6ebd8f6-91084-726-1080-jpg_250.jpg",
      "title": "Reze - Chainsaw Man (Good Luck Studio)",
      "category": "Misc",
      "rate": "background-position:-16px -1px;opacity:0.73333333333333",
      "date": "2023-06-04 10:36",
      "pages": "36 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571753/37b082a1d0/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/ea/7f/ea7f68134215aff923ac1f7bff7544fe33194244-3194501-2500-1875-jpg_250.jpg",
      "title": "[Chinjao Girl. (Special G)] Tonari no Onee-san no Shitagi o Nusundara Kiseki ga Okita Hanashi o Shiyou [Russian]",
      "category": "Manga",
      "rate": "background-position:0px -21px;opacity:0.8",
      "date": "2023-06-04 10:27",
      "pages": "114 pages",
      "seeds": "https://exhentai.org/img/td.png",
      "url": "https://exhentai.org/g/2571752/1607a6b8a8/",
      "error": ""
    },
    {
      "preview": "https://s.exhentai.org/t/96/c4/96c4994b0e7b09f5311e00c1baeefb2b0b746d42-3466231-2731-4600-jpg_250.jpg",
      "title": "[Artist] Kie Genshin Impact",
      "category": "Artist CG",
      "rate": "background-position:0px -21px;opacity:1",
      "date": "2023-06-04 10:26",
      "pages": "71 pages",
      "seeds": "https://exhentai.org/img/t.png",
      "url": "https://exhentai.org/g/2571751/3f050eb547/",
      "error": ""
    }
  ],
  "error": ""
}
```
##### gallery
```json
{
  "query": "g/2557923/cd8d42df44/?p=1",
  "preview": "",
  "title": "[Artist] Tsukiyo",
  "origin_title": "[アーティスト] 月夜",
  "tags": {
    "artist:": ["tsukiyo"],
    "female:": ["kemonomimi", "magical girl", "very long hair"],
    "other:": ["non-h imageset"],
    "parody:": ["pokemon | pocket monsters"]
  },
  "category": "Non-H",
  "pages": "398 pages",
  "images": [
    "https://s.exhentai.org/t/9e/23/9e236581d5e2ec33f0a45bcd7ceb9cde06cc07e2-1019546-1326-2220-jpg_l.jpg",
    "https://s.exhentai.org/t/a2/e9/a2e9abae27a98dd54064e089b54c64d275af383c-346946-1099-1016-jpg_l.jpg",
    "https://s.exhentai.org/t/95/d0/95d069d326c88c47685746feb33b53e8981ef1cd-683100-2660-1282-jpg_l.jpg",
    "https://s.exhentai.org/t/2a/02/2a028ef67f585952c15d6e9dd85667c3e258a38d-1426198-1732-2851-jpg_l.jpg",
    "https://s.exhentai.org/t/eb/8b/eb8b8fc5e70c1faa6be284e888e4b7d80cb26de3-594659-1795-2075-jpg_l.jpg",
    "https://s.exhentai.org/t/4f/2e/4f2ea929348f085903b3195e6bd438f9fc80ad3d-1176756-2084-2350-jpg_l.jpg",
    "https://s.exhentai.org/t/e5/b9/e5b95b4de54ada31e1c3e7ffb5b3be401a6fb545-1915840-2054-2812-jpg_l.jpg",
    "https://s.exhentai.org/t/e4/f2/e4f2425641422463f9e2c90dd9a809e856393ca8-689579-2660-1282-jpg_l.jpg",
    "https://s.exhentai.org/t/ca/5d/ca5d95ecc9da69bdaa5fbc09b3aa3886747570a8-792846-1644-2628-jpg_l.jpg",
    "https://s.exhentai.org/t/42/83/42836b0d10348aa1b9a56b09eb1503b7c96c309b-1062329-1689-2986-jpg_l.jpg",
    "https://s.exhentai.org/t/19/3a/193a5cc3212574bad62bd16d6673c399cabd4698-1899177-2045-2733-jpg_l.jpg",
    "https://s.exhentai.org/t/30/ce/30ceaff94ec608cbab8f869270233c90f22399c4-675231-1158-1631-jpg_l.jpg",
    "https://s.exhentai.org/t/d7/70/d7707b71812a1944f5fceb54cf4d68e8825a8b34-1738210-1784-2900-jpg_l.jpg",
    "https://s.exhentai.org/t/c8/4b/c84ba35e38cfdcb2ae522428a2b24cc66897a47f-2110686-2313-1835-jpg_l.jpg",
    "https://s.exhentai.org/t/86/ac/86acf511693ee86b613b0175c5a34c87df990794-1519326-1793-2847-jpg_l.jpg",
    "https://s.exhentai.org/t/61/ad/61ad939979438a725f04a6ec5dbb9a840fd6797d-818482-1725-1882-jpg_l.jpg",
    "https://s.exhentai.org/t/8e/e3/8ee3d4abf2495783959b6748959ae0c76b2a0be6-1546365-2200-2528-jpg_l.jpg",
    "https://s.exhentai.org/t/43/14/4314ee1ccdc4098bdc4396d569747de24adf1960-850129-1889-2534-jpg_l.jpg",
    "https://s.exhentai.org/t/33/24/33241d343d741155cea423e5acdfec42300994b5-1015640-1907-2100-jpg_l.jpg",
    "https://s.exhentai.org/t/27/55/2755bc6c77f2a0b366b9f0addfd2bb9b7a3ffe69-925316-1820-2751-jpg_l.jpg"
  ],
  "url": [
    "https://exhentai.org/s/9e236581d5/2557923-21",
    "https://exhentai.org/s/a2e9abae27/2557923-22",
    "https://exhentai.org/s/95d069d326/2557923-23",
    "https://exhentai.org/s/2a028ef67f/2557923-24",
    "https://exhentai.org/s/eb8b8fc5e7/2557923-25",
    "https://exhentai.org/s/4f2ea92934/2557923-26",
    "https://exhentai.org/s/e5b95b4de5/2557923-27",
    "https://exhentai.org/s/e4f2425641/2557923-28",
    "https://exhentai.org/s/ca5d95ecc9/2557923-29",
    "https://exhentai.org/s/42836b0d10/2557923-30",
    "https://exhentai.org/s/193a5cc321/2557923-31",
    "https://exhentai.org/s/30ceaff94e/2557923-32",
    "https://exhentai.org/s/d7707b7181/2557923-33",
    "https://exhentai.org/s/c84ba35e38/2557923-34",
    "https://exhentai.org/s/86acf51169/2557923-35",
    "https://exhentai.org/s/61ad939979/2557923-36",
    "https://exhentai.org/s/8ee3d4abf2/2557923-37",
    "https://exhentai.org/s/4314ee1ccd/2557923-38",
    "https://exhentai.org/s/33241d343d/2557923-39",
    "https://exhentai.org/s/2755bc6c77/2557923-40"
  ],
  "comments": null,
  "error": ""
}
```
##### image
```json
{
  "query": "s/8681587d21/2005521-5",
  "gallery_page": "https://exhentai.org/g/2005521/b485b3e9e0/",
  "next_page": "https://exhentai.org/s/f9cf5ee50e/2005521-6",
  "prev_page": "https://exhentai.org/s/b1b3ce098f/2005521-4",
  "image": "https://guuafnt.edpccgzarxix.hath.network/h/8681587d2159b49f5eeda1c0da6f523ef9b8de1e-54771-1039-1465-png/keystamp=1685963400-0731e3cc12;fileindex=97431048;xres=org/xCredit.png",
  "alt_image": "return nl('44243-468323')",
  "origin_image": "",
  "error": ""
}
```
##### torrents
```json
{
  "query": "gallerytorrents.php?gid=2558455\u0026t=e86edf3a07",
  "torrents": [
    {
      "date": "Posted: 2023-05-20 16:00",
      "size": "Size: 132.0 MB",
      "seed": "Seeds: 48",
      "peer": "Peers: 0",
      "download": "Peers: 0",
      "uploader": "Uploader: Bowden",
      "url": "https://exhentai.org/torrent/2558455/e9032fe6f5493a87724ced59a3356725a72bb6b9.torrent",
      "error": ""
    },
    {
      "date": "Posted: 2023-05-20 16:01",
      "size": "Size: 7.49 MB",
      "seed": "Seeds: 16",
      "peer": "Peers: 0",
      "download": "Peers: 0",
      "uploader": "Uploader: Bowden",
      "url": "https://exhentai.org/torrent/2558455/1596023bf110828e6f0f9cd9b04e5733aa6b4539.torrent",
      "error": ""
    }
  ],
  "error": ""
}
```


====


## mycurl

TODO:
- ~~request though proxy~~
- ~~request with additional cookie~~ 
- request with specific IP address.

## eh-api
``` go
	resp, err := mycurl.Fetch("GET", BASE_URL+query,
		map[string]string{"Cookie": COOKIE},
		nil)
	if err != nil {
		index.Error = err.Error()
		return
	}
```
this can be trim. maybe one day.
### index page

input:
- query, string, which for search
PS：留给前端做，要选项之类的。https://exhentai.org/?f_cats=767&f_search=tsukiyo

output:
this page with summary information


### gallery page

return meta infomation for gallery

return img previews of this page


### image page

add manual later.

if no origin image, the original image's value will be empty string.

## fiber 


proxy cover


====


4.5
background-position:0px -21px;opacity:1
4.0
background-position:-16px -3px;opacity:1
3.5
background-position:-16px -21px;opacity:1


```
Preview:https://s.exhentai.org/t/d3/41/d341a2000e7f7aed8ea6dcc099948a527d687582-1647015-1790-3018-jpg_250.jpg 
Title:[Artist] Tsukiyo 
OriginTitle: 
Tags:[] 
Category:Non-H 
Rate:background-position:0px -21px;opacity:1 
Date:2023-05-20 03:47 
Pages:398 pages 
Seeds:[] 
Url:https://exhentai.org/g/2557923/cd8d42df44/ 
Error:

```