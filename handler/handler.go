package handler

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type Url struct {
	URL  string `yaml:"url"`
	Path string `yaml:"path"`
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func HandleMap(UrlPath map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if des, ok := UrlPath[path]; ok {
			http.Redirect(w, r, des, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func HandleYAML(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	UrlPaths, err := parseYAML(yamlBytes)
	Check(err)
	paths := makeMap(UrlPaths)
	return HandleMap(paths, fallback), nil
}

func makeMap(UrlPaths []Url) map[string]string {
	path := make(map[string]string)
	for _, p := range UrlPaths {
		path[p.Path] = p.URL
	}
	return path
}

func parseYAML(data []byte) ([]Url, error) {
	var urls []Url
	err := yaml.Unmarshal(data, &urls)
	Check(err)
	return urls, nil
}
