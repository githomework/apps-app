package app

import (
	"fmt"
	"net/http"
)

func (a *AppType) PingResponse(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, a.Name)

}

func (a *AppType) FolderResponse(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, a.Folder)
}
