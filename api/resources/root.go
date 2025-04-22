package resources

import (
	"sync"
)

// TypedResource es una interfaz genérica para recursos con tipo específico
type TypedResource[T any] interface {
	Format(data T) interface{}
}

func FormatWith[T any, R TypedResource[T]](resource R, data T) interface{} {
	return resource.Format(data)
}

func FormatArrayWith[T any, R TypedResource[T]](resource R, data []T) []interface{} {
	if len(data) == 0 {
		return []interface{}{}
	}

	items := make([]interface{}, len(data))
	var wg sync.WaitGroup

	for i, item := range data {
		wg.Add(1)
		go func(idx int, elem T) {
			defer wg.Done()
			items[idx] = resource.Format(elem)
		}(i, item)
	}

	wg.Wait()
	return items
}
