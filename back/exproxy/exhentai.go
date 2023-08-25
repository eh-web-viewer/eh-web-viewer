package exproxy

import (
	"compress/gzip"
	"errors"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/andybalholm/brotli"
)

var re *regexp.Regexp = regexp.MustCompile(`Domain=[\w\.]*;`)

var client func() *http.Client

func SetClient(clientProviderFunc func() *http.Client) {
	client = clientProviderFunc
}

func getPlainTextReader(body io.ReadCloser, encoding string) (io.ReadCloser, error) {
	switch encoding {
	case "gzip":
		reader, err := gzip.NewReader(body)
		if err != nil {
			log.Println("error decoding gzip response", reader)
		}
		return reader, err
	case "br":
		reader := brotli.NewReader(body)
		var err error
		if reader == nil {
			log.Println("error decoding br response", reader)
			err = errors.New("error decoding br response")
		}
		return io.NopCloser(reader), err
	default:
		return body, nil
	}
}

// return the proxy function
func httpHandler(w http.ResponseWriter, r *http.Request) {
	newUrl := r.URL
	newUrl.Host = HOST
	newUrl.Scheme = "https"

	// if ip is not from CN then return 302
	if country := r.Header.Get("Cf-Ipcountry"); country != "CN" {
		http.Redirect(w, r, REDIRECT_URL, http.StatusFound)
		return
	}

	req, err := http.NewRequest(r.Method, newUrl.String(), r.Body)
	if err != nil {
		log.Println(`Error On NewRequest`, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for k, v := range r.Header {
		req.Header.Set(k, v[0])
	}
	// if COOKIE != "" {
	req.Header.Del("Cookie")
	req.Header.Set("Cookie", COOKIE)
	// }

	resp, err := client().Do(req)
	if err != nil {
		log.Println(`Error On Do Request`, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "text/html") ||
		strings.HasPrefix(contentType, "application/javascript") ||
		strings.HasPrefix(contentType, "text/css") {
		// if are text type
		body, err := getPlainTextReader(resp.Body, resp.Header.Get("Content-Encoding"))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		text, err := io.ReadAll(body)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		textReplaced := strings.Replace(string(text), URL_ORI, "", -1)
		textReplaced = strings.Replace(textReplaced, S_URL_ORI, S_URL_OW, -1)
		// if in image view, add buttom
		if strings.HasPrefix(r.URL.Path, `/s/`) {
			textReplaced = addWaterFallViewButton(textReplaced)
		} else if strings.HasPrefix(r.URL.Path, `/g/`) {
			textReplaced = addButtonToEHWV(textReplaced)
		}

		// process the header
		for k, v := range resp.Header {
			if k == "Content-Length" {
				continue
			}
			for _, vi := range v {
				if k == "Set-Cookie" {
					vi = string(re.ReplaceAll([]byte(vi), []byte{}))
				}
				w.Header().Add(k, vi)
			}
		}
		w.Header().Del("Content-Encoding")

		// w.WriteHeader(http.StatusOK)
		w.Write([]byte(textReplaced))
	} else {
		// process the header
		for k, v := range resp.Header {
			if k == "Content-Length" {
				continue
			}
			for _, vi := range v {
				if k == "Set-Cookie" {
					vi = string(re.ReplaceAll([]byte(vi), []byte{}))
				}
				w.Header().Add(k, vi)
			}
		}

		// w.WriteHeader(http.StatusOK)
		io.Copy(w, resp.Body)
	}
}

func Proxy(listen string) {

	handler := http.NewServeMux()
	handler.HandleFunc("/", httpHandler)
	server := &http.Server{Addr: listen, Handler: handler}

	err := server.ListenAndServe()

	log.Println(err)
}

func addButtonToEHWV(html string) string {
	return strings.Replace(html, "<body>", `<body>
	<div style="
		height: 60px;
		width: 100px;
		text-align: center;
		/* background-color: violet; */
		position: fixed;
		right: 20px; 
		top: 20px;
		z-index: 99;
		display: table-cell;
		vertical-align: middle;
		/* float: right; */
	">
		<a id="ehwv" href="">
			<button style="
				width: 100%;    
				height: 100%;
				font-size: x-large;
			">
				翻页式
			</button>
		</a>
	</div>
	<script>
	document.getElementById("ehwv").href = location.href.replace(location.host, "ehwv.moonchan.xyz");
	</script>`, 1)
}

func addWaterFallViewButton(html string) string {
	return strings.Replace(html, "<body>", `<body>
	<div style="
	  height: 60px;
	  width: 100px;
	  text-align: center;
	  /* background-color: violet; */
	  position: fixed;
	  right: 20px; 
	  top: 20px;
	  z-index: 99;
	  display: table-cell;
	  vertical-align: middle;
	  /* float: right; */
	">
	  <button id="waterfall" style="
			width: 100%;    
			height: 100%;
			font-size: x-large;
	  ">
			下拉式
	  </button>
		<a id="ehwv">
			<button style="
				width: 100%;    
				height: 100%;
				font-size: x-large;
			">
				翻页式
			</button>
		</a>
	</div>
  <script type="text/javascript">
	async function execWaterfall(){
		console.log('!');
		document.getElementById("waterfall").remove();
		let pn = document.createElement('div');
		let lp = location.href;
		let ln = location.href;
		const element = document.getElementById('i1');
		element.appendChild(pn);
		let hn = document.getElementById('next').href;
		while (hn != ln) {
		  let doc;
		  while(!doc) {
			doc = await fetch(hn).then(resp => resp.text())			
			  .then(data => {
			    console.log(data);
			    let parser = new DOMParser();
			    let doc = parser.parseFromString(data, "text/html");
			    return doc;
			  });
			}
		  console.log(doc);
		  let img = document.createElement('img');
		  let element = doc.getElementById('img');
		  if (element) {
			img.src = element.src;
			pn.appendChild(img);
			ln = hn;
			hn = doc.getElementById('next').href;
		  }
		}
		let p = document.createElement('p');
		p.innerHTML = hn;
	  }
	document.getElementById("waterfall").addEventListener("click", execWaterfall, false); 
	document.getElementById("ehwv").href = location.href.replace(location.host, "ehwv.moonchan.xyz");
	</script>`, 1)
}
