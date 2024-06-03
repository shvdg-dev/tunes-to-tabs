package handlers

import (
	"net/http"
	"tabs/internal/constants"
	inf "tabs/internal/info"
)

// JobsPage handles the request for the home page.
func (j *Jobs) JobsPage(writer http.ResponseWriter, request *http.Request) {
	info := j.Context.Informer.NewInfo(request, inf.WithTitle(constants.BundleJobs))
	page := j.Views.Jobs.CreateJobsPage()
	j.Renderer.Render(writer, request, info, page)
}
