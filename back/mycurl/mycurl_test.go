package mycurl

import (
	"io"
	"log"
	"net/http"
)

func printResp(resp *http.Response) {
	respText, _ := io.ReadAll(resp.Body)
	log.Println(string(respText))
}

// cookie jar is perfect.
// it works as addons on the given cookie
// and here shows how to add a new cookie to a jar.
// because client refer the jar by pointer, change the jar in the public domain works for every clinet.

// func Test_CookieJar(t *testing.T) {
// 	SetClient("")
// 	print(client, vanillaClient)
// 	resp, _ := fetch("GET", DST_URL,
// 		map[string]string{"Cookie": "nani"},
// 		nil)
// 	printResp(resp)

// 	SetClient("proxy")
// 	print(client, proxyClient)
// 	resp, _ = fetch("GET", DST_URL,
// 		map[string]string{"Cookie": "nani"},
// 		nil)
// 	printResp(resp)

// 	// it works
// 	cookie := &http.Cookie{
// 		Name:   "M_WEIBOCN_PARAMS",
// 		Value:  "rl%3D1",
// 		Path:   "/",
// 		Domain: ".weibo.com",
// 	}
// 	cookies := []*http.Cookie{cookie}
// 	u, _ := url.Parse("https://weibo.com")
// 	jar.SetCookies(u, cookies)
// 	log.Println(*jar) // 这里能看见。

// 	req, _ := generateRequest("GET", DST_URL,
// 		// map[string]string{"Cookie": "nani"},
// 		nil,
// 		nil)
// 	resp, _ = proxyClient.Do(req)
// 	printResp(resp)

// 	req, _ = generateRequest("GET", DST_URL,
// 		map[string]string{"Cookie": "nani"},
// 		// nil,
// 		nil)
// 	resp, _ = proxyClient.Do(req)
// 	printResp(resp)

// }

// func Test_spefic(t *testing.T) {
// 	SetClient("specificIP")
// 	print(client, specificIPClient)
// 	resp, _ := fetch("GET", "https://6.ipw.cn",
// 		map[string]string{"Cookie": "nani"},
// 		nil)
// 	printResp(resp)

// }
