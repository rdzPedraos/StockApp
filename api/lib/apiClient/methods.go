package apiClient

import (
	"log"
	"net/http"
	"net/url"
)

func (c *APIClient) Get(path string, result interface{}) error {
	return c.doRequest(http.MethodGet, path, nil, result)
}

func (c *APIClient) GetWithQueryParams(path string, result interface{}, queryParams map[string]string) error {
	path = c.addQueryParams(path, queryParams)
	return c.doRequest(http.MethodGet, path, nil, result)
}

func (c *APIClient) Post(path string, body interface{}, result interface{}) error {
	return c.doRequest(http.MethodPost, path, body, result)
}

func (c *APIClient) Put(path string, body interface{}, result interface{}) error {
	return c.doRequest(http.MethodPut, path, body, result)
}

func (c *APIClient) Delete(path string, result interface{}) error {
	return c.doRequest(http.MethodDelete, path, nil, result)
}

func (c *APIClient) addQueryParams(path string, queryParams map[string]string) string {
	if len(queryParams) == 0 {
		return path
	}

	parsedURL, err := url.Parse(path)
	if err != nil {
		log.Printf("Error al analizar la URL %v: %v", path, err)
		return path
	}

	query := parsedURL.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}

	parsedURL.RawQuery = query.Encode()
	return parsedURL.String()
}
