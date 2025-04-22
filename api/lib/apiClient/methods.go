package apiClient

import (
	"net/http"
)

func (c *APIClient) Get(path string, result interface{}) error {
	return c.doRequest(http.MethodGet, path, nil, result)
}

func (c *APIClient) GetWithQueryParams(path string, result interface{}, queryParams map[string]string) error {
	path = c.addQueryParams(path, queryParams)
	return c.Get(path, result)
}

// GetBinary realiza una petición GET y devuelve la respuesta como datos binarios
func (c *APIClient) GetBinary(path string, result *BinaryResult) error {
	return c.doRequestBinary(http.MethodGet, path, nil, result)
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
