package handlers

import (
	"github.com/shvdg-dev/tunes-to-tabs/internal/constants"
	inf "github.com/shvdg-dev/tunes-to-tabs/internal/info"
	"net/http"
)

// JobsPage handles the request for the home page.
func (j *Jobs) JobsPage(writer http.ResponseWriter, request *http.Request) {
	info := j.Context.Informer.NewInfo(request, inf.WithTitle(constants.BundleJobs))
	page := j.Views.Jobs.CreateJobsPage()
	j.Renderer.Render(writer, request, info, page)
}
