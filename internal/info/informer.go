package info

import (
	"github.com/shvdg-dev/base-pkg/sessions"
	"net/http"
)

// Informer provides information about a particular request.
type Informer struct {
	Sessions *sessions.Service
}

// NewInformer creates a new Informer instance.
func NewInformer(sessions *sessions.Service) *Informer {
	return &Informer{Sessions: sessions}
}

// NewInfo creates a new Info object based on the provided request and options.
func (i *Informer) NewInfo(request *http.Request, opts ...Option) *Info {
	info := &Info{Path: request.URL.Path, IsAuthenticated: false}
	for _, opt := range opts {
		opt(info)
	}
	return info
}
