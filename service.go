package main

import	(
    "fmt"
	"io"
	"io/ioutil"
	"github.com/codegangsta/martini"
	"net/http"
    "net/url"
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

					fmt.Println(r.URL.RawQuery)
					m, _ := url.ParseQuery(r.URL.RawQuery)

					// need to wrap the file in a function if jsonp
					content, _ := ioutil.ReadFile(b.path + "/GET")

					// check to see if this is a jsonp call
					if value, ok := m["callback"]; ok {
						s := fmt.Sprintf("%s(%s)", value[0], string(content))
						io.WriteString(w, s)
					} else {
				        io.WriteString(w, string(content))
				    }
				}
			}
		}

		fmt.Println(r.URL.Path)
		r.URL.Path = ""
	}
}

