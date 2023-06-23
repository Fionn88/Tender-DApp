package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*   Route  */
type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("GET", "/InitLedger", InitLedger, nil)
	register("POST", "/CreateData", CreateData, nil)
	register("GET", "/ReadData", ReadData, nil)
	register("GET", "/GetHistory", GetHistory, nil)
	register("POST", "/UpdateData", UpdateData, nil)
	// register("GET", "/DeleteData", DeleteData, nil)

}

// NewRouter returns router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
