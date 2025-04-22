package apiClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type BinaryResult struct {
	Data        []byte
	ContentType string
}

func (c *APIClient) doRequest(method, path string, body interface{}, result interface{}) error {
	req, err := c.generateRequest(method, path, body)

	if err != nil {
		return err
	}

	return c.executeRequest(req, func(resp *http.Response) error {
		if result != nil {
			if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
				return fmt.Errorf("error al decodificar respuesta: %w", err)
			}
		}

		return nil
	})
}

// doRequestBinary ejecuta la petición HTTP y devuelve los bytes de la respuesta y su tipo de contenido
func (c *APIClient) doRequestBinary(method, path string, body interface{}, result *BinaryResult) error {
	req, err := c.generateRequest(method, path, body)

	if err != nil {
		return err
	}

	return c.executeRequest(req, func(resp *http.Response) error {
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		*result = BinaryResult{
			Data:        bodyBytes,
			ContentType: resp.Header.Get("Content-Type"),
		}

		return nil
	})
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

func (c *APIClient) generateRequest(method, path string, body interface{}) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)

	var reqBody io.Reader

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error al serializar body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error al crear la petición: %w", err)
	}

	// Agregar headers
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (c *APIClient) executeRequest(request *http.Request, callback func(resp *http.Response) error) error {
	resp, err := c.HTTPClient.Do(request)

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

	return callback(resp)
}
