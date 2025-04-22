package resources

// TypedResource es una interfaz genérica para recursos con tipo específico
type TypedResource[T any] interface {
	Format(data T) interface{}
}

func FormatWith[T any, R TypedResource[T]](resource R, data T) interface{} {
	return resource.Format(data)
}

func FormatArrayWith[T any, R TypedResource[T]](resource R, data []T) []interface{} {
	items := make([]interface{}, 0, len(data))

	for _, item := range data {
		newItem := resource.Format(item)
		items = append(items, newItem)
	}

	return items
}
