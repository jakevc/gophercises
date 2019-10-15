package urlshort

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"net/http"
)

// struct to parse yaml string
type T []struct {
	P string `yaml:"path"`
	U string `yaml:"url"`
}

// parse yaml into slice of structs
func parseYaml(yml []byte) T {
	t := T{}
	err := yaml.Unmarshal([]byte(yml), &t)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return t
}

// make map from slice of structs
func makeMap(yamlStruct T) map[string]string {
	urlMap := make(map[string]string)
	for _, s := range yamlStruct {
		urlMap[s.P] = s.U
	}
	return urlMap
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// handler redirects shortened urls to long ones from a map.
	return func(w http.ResponseWriter, r *http.Request) {
		pth := r.URL.Path[0:]
		// execute redirect if path is in map
		if s, ok := pathsToUrls[pth]; ok {
			http.Redirect(w, r, s, 301)
		}
		// fallback to the provided handler
		fallback.ServeHTTP(w, r)
	}
}

// handle the yaml map using the mapHandler
func YAMLHandler(yml []byte, fallback http.Handler) http.HandlerFunc {
	yamS := parseYaml(yml)
	yamM := makeMap(yamS)
	return MapHandler(yamM, fallback)
}
