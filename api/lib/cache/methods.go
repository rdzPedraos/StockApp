package cache

import (
	"encoding/json"
	"time"
)

func Ping() error {
	return cache.rdb.Ping(cache.ctx).Err()
}

// Get obtiene un valor de la caché y lo deserializa al tipo especificado a través del puntero result
func Get[T any](key string, result *T) error {
	// Obtener el valor crudo de Redis
	val, err := cache.rdb.Get(cache.ctx, key).Result()
	if err != nil {
		return err
	}

	// Deserializar el valor de string a T
	return json.Unmarshal([]byte(val), result)
}

// Set serializa y guarda un valor en la caché
func Set[T any](key string, value T, expiration time.Duration) error {
	// Serializar el valor a JSON
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// Guardar en caché
	return cache.rdb.Set(cache.ctx, key, string(jsonBytes), expiration).Err()
}

func Delete(key string) error {
	return cache.rdb.Del(cache.ctx, key).Err()
}

// Once es una versión tipada que almacena el resultado directamente en el puntero proporcionado
func Once[T any](key string, expiration time.Duration, result *T, fn func() (*T, error)) error {
	// Intentamos obtener el valor desde la caché
	err := Get(key, result)
	if err == nil {
		return nil
	}

	// Si no hay valor en caché o hubo error al deserializar, ejecutamos la función
	tempResult, err := fn()
	if err != nil {
		return err
	}

	// Copiamos el resultado al puntero proporcionado
	*result = *tempResult

	// Guardamos en caché para futuras peticiones
	return Set(key, *result, expiration)
}
