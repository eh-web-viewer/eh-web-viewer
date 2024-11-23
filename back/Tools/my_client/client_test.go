package myclient

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func TestClient(t *testing.T) {
	tests := []struct {
		name string
		ip   string
	}{
		{"IP1", "240e:38c:8e9b:ce00:21a9:f386:d35f:d47d"},
		{"IP2", "240e:38c:8e9b:ce00:52a:23a5:6298:7a06"},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// ip := net.ParseIP()
			ip := net.ParseIP(tt.ip)

			jar, _ := cookiejar.New(nil)

			client := NewClientWithLocalAddress(ip, jar)

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

			if string(body) != tt.ip {
				t.Fatalf("Expected ip %v but got %v", tt.ip, string(body))
			}

			fmt.Println("Response body:", string(body))
		})
	}
}
