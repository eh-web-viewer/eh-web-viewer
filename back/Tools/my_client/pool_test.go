package myclient

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
)

func TestO(t *testing.T) {
	TestClient := func(client *http.Client) {

		// Use a known URL for testing.
		url := "https://api64.ipify.org"

		resp, err := client.Get(url)
		if err != nil {
			t.Fatalf("Failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status OK but got %v", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		// if string(body) != tt.ip {
		// 	t.Fatalf("Expected ip %v but got %v", tt.ip, string(body))
		// }

		fmt.Println("Response body:", string(body))
	}

	jar, _ := cookiejar.New(nil)
	pool := NewClientPool()

	TestClient(pool.Get())
	TestClient(pool.Get())
	fmt.Println("=====")

	{
		ip := net.ParseIP("240e:38c:8e9b:ce00:52a:23a5:6298:7a06")
		client := NewClientWithLocalAddress(ip, jar)
		pool.Add(client)
	}

	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	fmt.Println("=====")

	{
		ip := net.ParseIP("240e:38c:8e9b:ce00:21a9:f386:d35f:d47d")
		client := NewClientWithLocalAddress(ip, jar)
		pool.Add(client)
	}

	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	fmt.Println("=====")

	{
		client := NewInsecureClient(jar)
		pool.Add(client)
	}

	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	fmt.Println("=====")

	{
		proxyURL, _ := url.Parse("http://localhost:10809")
		client := NewProxyClient(proxyURL, jar)
		pool.Add(client)
	}

	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	TestClient(pool.Get())
	fmt.Println("=====")

}
