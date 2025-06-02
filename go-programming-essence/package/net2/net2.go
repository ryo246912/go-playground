package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type routerParam map[string]string

type routerFunc func(routerParam, http.ResponseWriter, *http.Request)

type routerItem struct {
	method  string
	matcher *regexp.Regexp
	fuc     routerFunc
}
type router struct {
	items []routerItem
}

func (rt *router) GET(prefix string, fuc routerFunc) {
	rt.items = append(rt.items, routerItem{
		method:  http.MethodGet,
		matcher: regexp.MustCompile(prefix),
		fuc:     fuc,
	})
}

func (rt *router) POST(prefix string, fuc routerFunc) {
	rt.items = append(rt.items, routerItem{
		method:  http.MethodPost,
		matcher: regexp.MustCompile(prefix),
		fuc:     fuc,
	})
}

func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, item := range rt.items {
		if item.method == r.Method && item.matcher.MatchString(r.URL.Path) {
			params := make(routerParam)
			matches := item.matcher.FindStringSubmatch(r.URL.Path)
			if len(matches) > 1 {
				for i, name := range item.matcher.SubexpNames() {
					if i > 0 && name != "" {
						params[name] = matches[i]
					}
				}
			}
			item.fuc(params, w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	rt := &router{}

	rt.GET(`/hello/(?P<name>\w+)`, func(params routerParam, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello", params["name"])
	})

	log.Fatal(http.ListenAndServe(":8080", rt))
}
