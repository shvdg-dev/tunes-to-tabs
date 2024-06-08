package info

import (
	"github.com/shvdg-dev/base-pkg/sessions"
	consts "github.com/shvdg-dev/tunes-to-tabs/internal/constants"
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
	info := &Info{Path: request.URL.Path, IsAuthenticated: i.IsAuthenticated(request)}
	for _, opt := range opts {
		opt(info)
	}
	return info
}

// IsAuthenticated checks if the user is authenticated based on the session stored in the Informer object.
func (i *Informer) IsAuthenticated(request *http.Request) bool {
	isAuthenticated, ok := i.Sessions.Get(consts.ValueIsAuthenticated, request).(bool)
	if !ok {
		return false
	}
	return isAuthenticated
}

func (i *Informer) SetAuthenticated(request *http.Request) {
	i.Sessions.Store(consts.ValueIsAuthenticated, true, request)
}

func (i *Informer) SetUnauthenticated(request *http.Request) {
	i.Sessions.Store(consts.ValueIsAuthenticated, false, request)
}
