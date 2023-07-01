# eh-web-viewer-back

the backend of eh-web-viewer.

convert eh webpage to api format

mainly contain these catalogsf
 
- index page
  - it contains the gallery facebooks, and the next page and previous page.-
- seed page
  - view only the seed info in the gallery.
- gallery page
  - view the infomation of the gallery
- view page
  - where to read the comic, in both water fall mode and book mode (maybe should add from left or from right)



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