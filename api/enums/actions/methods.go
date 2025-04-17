package actions

import (
	"fmt"
	"strings"
)

func Parse(value string) (Action, error) {
	value = strings.Replace(value, " by", "", 1)
	action, ok := stringToAction[value]

	if !ok {
		return "", fmt.Errorf("action no válida: %s", value)
	}
	return action, nil
}
