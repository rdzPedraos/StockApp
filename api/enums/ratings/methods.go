package ratings

import (
	"log"
	"slices"
	"strings"
)

func Parse(value string) Rating {
	value = strings.Replace(value, " by", "", 1)
	rating, ok := stringToRating[value]

	if !ok {
		log.Printf("Advertencia: enum.Rating no válido: %s", value)
		return ""
	}

	return rating
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
