package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Headers    map[string]string
}

// NewAPIClient crea una nueva instancia de APIClient con configuración por defecto
func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

// SetHeader establece un header personalizado para las peticiones
func (c *APIClient) SetHeader(key, value string) {
	c.Headers[key] = value
}

// SetAuthToken establece un token de autorización en los headers
func (c *APIClient) SetAuthToken(token string) {
	c.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
}

// Get realiza una petición GET a la API
func (c *APIClient) Get(path string, result interface{}) error {
	return c.doRequest(http.MethodGet, path, nil, result)
}

// Post realiza una petición POST a la API
func (c *APIClient) Post(path string, body interface{}, result interface{}) error {
	return c.doRequest(http.MethodPost, path, body, result)
}

// Put realiza una petición PUT a la API
func (c *APIClient) Put(path string, body interface{}, result interface{}) error {
	return c.doRequest(http.MethodPut, path, body, result)
}

// Delete realiza una petición DELETE a la API
func (c *APIClient) Delete(path string, result interface{}) error {
	return c.doRequest(http.MethodDelete, path, nil, result)
}

// doRequest ejecuta la petición HTTP con los parámetros proporcionados
func (c *APIClient) doRequest(method, path string, body interface{}, result interface{}) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("error al serializar body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("error al crear la petición: %w", err)
	}

	// Agregar headers
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	// Ejecutar la petición
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error al ejecutar la petición: %w", err)
	}
	defer resp.Body.Close()

	// Manejar errores HTTP
	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error en la petición HTTP: código %d, respuesta: %s",
			resp.StatusCode, string(bodyBytes))
	}

	// Decodificar respuesta JSON si se proporcionó un destino
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("error al decodificar respuesta: %w", err)
		}
	}

	return nil
}
