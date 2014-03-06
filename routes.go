package main

import 	(
	"path/filepath"
	"os"
)

type Route struct {
	path string
}

type Routes interface {
	Load()
	RouteList() []Route
}

type RoutesImpl struct {
	routelist []Route
}

func (routes *RoutesImpl) Load() {

	OrgPath := "./"
	Recurse := true
	walkFn := func(path string, info os.FileInfo, err error) error {
		stat, err := os.Stat(path)
		if err != nil {
			return err
		}

		if stat.IsDir() && path != OrgPath && !Recurse {
			return filepath.SkipDir
		}

		if err != nil {
			return err
		}
		if stat.IsDir() {
			routes.routelist = append(routes.routelist, Route{path})
		}

		return nil
	}

	filepath.Walk(OrgPath, walkFn)
}

func (routes *RoutesImpl) RouteList() []Route {
	return routes.routelist
}
