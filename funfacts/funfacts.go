package conv

type FunFacts struct {
	Sun   []string
	Moon  []string
	Earth []string
}

func GetFunFacts(about string) []string {
	funFacts := &FunFacts{
		Sun:   []string{"The sun is the closest star to Earth."},
		Moon:  []string{"The moon is about 4.5 billion years old."},
		Earth: []string{"Earth is the third planet from the sun."},
	}

	switch about {
	case "sun":
		return funFacts.Sun
	case "moon":
		return funFacts.Moon
	default:
		return funFacts.Earth
	}
}
