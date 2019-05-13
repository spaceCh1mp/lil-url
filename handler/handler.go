package handler

import "net/http"

type Url struct {
	URL  string `yaml:"url"`
	Path string `yaml:"path"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func HandleMap(UrlPath map[string]string, fallback http.Handler) http.HandleFunc {
	// TODO:
	return nil
}

func HandleYAML(yamlBytes []byte, fallback http.Handler) (http.HandleFunc, error) {
	UrlPaths, err := parseYAML(yamlBytes)
	check(err)
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
	urls := new(Url)
	err := yaml.Unmarshal(data, &urls)
	check(err)
	return urls, nil
}
