package ratings

import (
	"fmt"
	"slices"
	"strings"
)

func Parse(value string) (Rating, error) {
	value = strings.Replace(value, " by", "", 1)
	rating, ok := stringToRating[value]

	if !ok {
		return "", fmt.Errorf("rating no válido: %s", value)
	}
	return rating, nil
}

func (r Rating) Label() string {
	switch {
	case slices.Contains(positiveRatings, r):
		return "Positive"
	case slices.Contains(cautiousRatings, r):
		return "Cautious"
	case slices.Contains(negativeRatings, r):
		return "Negative"
	default:
		return "Neutral"
	}
}
