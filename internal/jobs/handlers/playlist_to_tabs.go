package handlers

import (
	"log"
	"net/http"
	"tabs/internal/constants"
)

// PlaylistToTabs handles the request for the home page.
func (j *Jobs) PlaylistToTabs(writer http.ResponseWriter, request *http.Request) {
	playlist := request.FormValue(constants.ValuePlaylistLink)
	log.Println(playlist)
}
