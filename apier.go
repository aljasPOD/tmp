package main

import (
	"ap/generic/apul"
	"fmt"
	"io"
	"net/http"
	"time"
)

var clnt = &http.Client{}

func getTorrents(l apul.Logger) bool {
	fmt.Printf(".")
	req, err := http.NewRequest("GET", "http://127.0.0.1:8081/api/v2/torrents/info", nil)
	if err != nil {
		l.Error("Failed to get torrent: %v", err)
		return false
	}

	req.Header.Set("User-Agent", "aP qbtControl")
	req.Header.Set("Connection", "close")

	resp, err := clnt.Do(req)
	if err != nil {
		l.Error("Failed to get response: %v", err)
		return false
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		l.Error("Failed to read response: %v", err)
		return false
	}

	if resp.StatusCode != 200 {
		l.Error("Error code %d returned for torrents: %s", resp.StatusCode, string(data))
		return false
	}
	return true
}

func main() {
	l := apul.SetupConsole(apul.LGTest)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				if !getTorrents(l) {
					return
				}
			}
		}()
	}
	for {
		time.Sleep(time.Minute)
	}
}
