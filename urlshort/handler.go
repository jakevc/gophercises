package urlshort

import (
	"fmt"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// handler redirects shortened urls to long ones from a map.
	return func(w http.ResponseWriter, r *http.Request) {
		pth := r.URL.Path[0:]
		fmt.Println(pth)
		// execute redirect if path is in map
		if s, ok := pathsToUrls[pth]; ok {
			fmt.Println(s)
			http.Redirect(w, r, s, 301)
		}
		// fallback to the provided handler
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return nil, nil
}
