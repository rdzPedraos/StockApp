package apiClient

import "fmt"

// SetHeader establece un header personalizado para las peticiones
func (c *APIClient) SetHeader(key, value string) {
	c.Headers[key] = value
}

// SetAuthToken establece un token de autorización en los headers
func (c *APIClient) SetAuthToken(token string) {
	c.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
}
