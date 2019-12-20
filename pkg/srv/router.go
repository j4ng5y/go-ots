package srv

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// newRouter returns a new gorilla mux router with routes injected
func (S *Server) newRouter() http.Handler {
	R := mux.NewRouter()

	R.Handle("/api/v1/secret", handlers.MethodHandler{
		http.MethodGet: handlers.ContentTypeHandler(
			handlers.CompressHandler(
				handlers.CombinedLoggingHandler(
					os.Stdout,
					http.HandlerFunc(S.getHandler))),
			"application/json"),
		http.MethodDelete: handlers.ContentTypeHandler(
			handlers.CompressHandler(
				handlers.CombinedLoggingHandler(
					os.Stdout,
					http.HandlerFunc(S.deleteHandler))),
			"application/json"),
		http.MethodPost: handlers.ContentTypeHandler(
			handlers.CompressHandler(
				handlers.CombinedLoggingHandler(
					os.Stdout,
					http.HandlerFunc(S.createHandler))),
			"application/json"),
	})

	return handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodDelete}),
		handlers.AllowedOrigins([]string{"*"}))(R)
}
