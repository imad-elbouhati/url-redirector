package urlredirect

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler returns an http.HandlerFunc that attempts to map paths to their corresponding URLs.
// If the requested path is found in the pathsToUrls map, it redirects the client to the corresponding URL.
// If the path is not found, it calls the fallback http.Handler.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestedPath := r.URL.Path

		if url, ok := pathsToUrls[requestedPath]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

// PathToURL represents a YAML structure containing path-url mappings.
type PathToURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// YAMLHandler returns an http.HandlerFunc that reads path-url mappings from YAML data and attempts to map paths to their corresponding URLs.
// If the requested path is found in the parsed YAML, it redirects the client to the corresponding URL.
// If the path is not found, it calls the fallback http.Handler.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(yml)
	if err != nil {
		fmt.Print(err.Error())
	}

	pathMap := buildMap(parsedYAML)

	for k, v := range pathMap {
		fmt.Printf("Key: %v, Value: %v \n", k, v)
	}

	return MapHandler(pathMap, fallback), nil
}

// parseYAML parses the YAML data and returns a slice of PathToURL objects.
func parseYAML(yml []byte) ([]PathToURL, error) {
	var data []PathToURL

	err := yaml.Unmarshal(yml, &data)
	if err != nil {
		return nil, errors.New("cannot unmarshal data: " + err.Error())
	}

	return data, nil
}

// buildMap builds a map of path-url mappings from the parsed YAML data.
func buildMap(parsedYAML []PathToURL) map[string]string {
	var pathMap = make(map[string]string)

	for _, data := range parsedYAML {
		pathMap[data.Path] = data.URL
	}

	return pathMap
}
