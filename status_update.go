package app

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Harcoded port 7713 for processing status updates
func UpdateStatusBegan7713(session string) {
	var jsonStr strings.Builder
	jsonStr.Grow(100)
	jsonStr.WriteString(`{"session_id":"`)
	jsonStr.WriteString(session)
	jsonStr.WriteString(`","began":"`)
	jsonStr.WriteString(time.Now().Format("2006-01-02 15:04:05.000"))
	jsonStr.WriteString(`"}`)
	req, err := http.NewRequest("POST", "http://127.0.0.1:7713/updateStatusBegan", strings.NewReader(jsonStr.String()))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)

}

// Harcoded port 7713 for processing status updates
func UpdateStatusFinished7713(session string) {
	var jsonStr strings.Builder
	jsonStr.Grow(100)
	jsonStr.WriteString(`{"session_id":"`)
	jsonStr.WriteString(session)
	jsonStr.WriteString(`","finished":"`)
	jsonStr.WriteString(time.Now().Format("2006-01-02 15:04:05.000"))
	jsonStr.WriteString(`"}`)
	req, err := http.NewRequest("POST", "http://127.0.0.1:7713/updateStatusFinished", strings.NewReader(jsonStr.String()))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)

}
