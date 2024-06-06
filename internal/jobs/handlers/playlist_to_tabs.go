package handlers

import (
	"github.com/shvdg-dev/tunes-to-tabs/internal/constants"
	"log"
	"net/http"
)

// PlaylistToTabs handles the request for the home page.
func (j *Jobs) PlaylistToTabs(writer http.ResponseWriter, request *http.Request) {
	playlist := request.FormValue(constants.ValuePlaylistLink)
	log.Println(playlist)
}
