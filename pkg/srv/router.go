package srv

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// newRouter returns a new gorilla mux router with routes injected
func (S *Server) newRouter() *mux.Router {
	R := mux.NewRouter()

	R.Handle("/api/v1/secret", handlers.MethodHandler{
		http.MethodGet:    handlers.CompressHandler(handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(S.retrieveHandler))),
		http.MethodDelete: handlers.CompressHandler(handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(S.burnHandler))),
		http.MethodPost:   handlers.CompressHandler(handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(S.generateHandler))),
		http.MethodPut:    handlers.CompressHandler(handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(S.createHandler))),
	})
	R.Handle("/api/v1/metadata", handlers.MethodHandler{
		http.MethodGet: handlers.CompressHandler(handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(S.retrieveMetadataHandler))),
	})

	return R
}
