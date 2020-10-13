package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
)

func (a *AppType) PingResponse(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, a.Name)

}

func (a *AppType) FolderResponse(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, a.Folder)
}

func ExtractSessionInfo(req *http.Request) map[string]string {
	m := map[string]string{}
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return m
	}
	jsonparser.EachKey(b, func(i int, value []byte, vt jsonparser.ValueType, err error) {
		if err != nil {
			return
		}
		switch i {
		case 0:
			m["session_id"] = string(value)
		case 1:
			m["began"] = string(value)
		case 2:
			m["finished"] = string(value)
		case 3:
			m["task_name"] = string(value)
		case 4:
			m["brcd"] = string(value)
		case 5:
			m["printer"] = string(value)
		case 6:
			m["ref"] = string(value)
		case 7:
			m["content"] = string(value)
		case 8:
			m["queue_name"] = string(value)
		case 9:
			m["timeout"] = string(value)

		}
	}, []string{"session_id"},
		[]string{"began"},
		[]string{"finished"},
		[]string{"task_name"},
		[]string{"brcd"},
		[]string{"printer"},
		[]string{"ref"},
		[]string{"content"},
		[]string{"queue_name"},
		[]string{"timeout"})
	return m
}
