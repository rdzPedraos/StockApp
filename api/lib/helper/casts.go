package helper

import (
	"log"
	"strconv"
	"strings"
)

func ParseFloat(s string) float32 {
	if s == "" {
		return 0
	}

	s = strings.TrimSpace(s)
	s = strings.Replace(s, "$", "", -1)
	s = strings.Replace(s, ",", "", -1)

	value, err := strconv.ParseFloat(s, 32)

	if err != nil {
		log.Printf("Advertencia: No se pudo convertir '%s' a float: %v", s, err)
		return 0
	}

	return float32(value)
}
