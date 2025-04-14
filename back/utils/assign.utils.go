package utils

import (
	"reflect"
)

func validateValue[T any](value *T, bypassDefault bool) bool {
	valueReflect := reflect.ValueOf(value)
	kind := valueReflect.Kind()

	canBeNil := func(k reflect.Kind) bool {
		return k == reflect.Pointer || k == reflect.Slice ||
			k == reflect.Map || k == reflect.Interface ||
			k == reflect.Chan || k == reflect.Func
	}

	// Verifica si el valor es nil
	if canBeNil(kind) {
		if valueReflect.IsNil() {
			return false
		}
	}

	// Verifica si el valor es vacío (igual al valor cero)
	if bypassDefault && reflect.DeepEqual(value, new(T)) {
		return false
	}

	return true
}

func validateTrutyValue[T any](value *T) bool {
	return validateValue(value, true)
}

// Coalesce devuelve el primer valor que no esté vacio ni nulo de la lista de valores proporcionados.
// Esta función es similar al operador de coalescencia nula (??) en otros lenguajes.
func Coalesce[T any](values ...T) T {
	for _, value := range values {
		if validateTrutyValue(&value) {
			return value
		}
	}

	// Si todos son vacíos o nulos, devolvemos el último valor
	if len(values) > 0 {
		return values[len(values)-1]
	}

	return *new(T)
}
