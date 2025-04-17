package actions

type Action string

const (
	downgraded    Action = "downgraded"
	initiated     Action = "initiated"
	reiterated    Action = "reiterated"
	targetLowered Action = "target lowered"
	targetRaised  Action = "target raised"
	targetSet     Action = "target set"
	upgraded      Action = "upgraded"
)

var stringToAction = map[string]Action{
	"downgraded":     downgraded,
	"initiated":      initiated,
	"reiterated":     reiterated,
	"target lowered": targetLowered,
	"target raised":  targetRaised,
	"target set":     targetSet,
	"upgraded":       upgraded,
}
