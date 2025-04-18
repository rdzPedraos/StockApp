package actions

import (
	"log"
	"strings"
)

func Parse(value string) Action {
	value = strings.Replace(value, " by", "", 1)
	action, ok := stringToAction[value]

	if !ok {
		log.Printf("Advertencia: enum.Action no válida: %s", value)
		return ""
	}
	return action
}
