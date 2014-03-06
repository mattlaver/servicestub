package main

import	(
    "fmt"
	"github.com/codegangsta/martini"
	"net/http"
	"strings"
)

type Service interface {
	InitRoutes(routes []Route)
	Start(settings Settings)
}

type ServiceImpl struct {
	routes []Route
}

func (service *ServiceImpl) InitRoutes(routes []Route) {
	service.routes = routes
}

func (service *ServiceImpl) Start(settings Settings) {
	m := martini.New()
	m.Use(ParseRoute(service))
	http.ListenAndServe(":3000", m)
}

func ParseRoute(service *ServiceImpl) martini.Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		for _, b := range service.routes {

			path := strings.TrimPrefix(r.URL.Path, "/")
			if b.path == path {

				if r.Method == "GET" {
					http.ServeFile(w, r, b.path + "/GET")
				}

				// serve up the file
				fmt.Println(r.Method)
				fmt.Fprintf(w, "Path Found")
			}
		}

		fmt.Println(r.URL.Path)
		r.URL.Path = ""
	}
}

