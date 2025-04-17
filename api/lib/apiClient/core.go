package apiClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
