package statTracker

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    100 * time.Millisecond,
		WriteTimeout:   200 * time.Millisecond,
	}

	return s.httpServer.ListenAndServe()
}
