package funfacts

func GetFunFacts(about string) []string {
	funFacts := new(FunFacts)

	funFacts.Sun = []string{"Temperatur i Solens kjerne 15000000 C", "Temperatur på ytre lag av Solen 5778 K"}
	funFacts.Luna = []string{"Temperatur på Månens overflate om natten -183 C", "Temperatur på Månens overflate om dagen 106 C"}
	funFacts.Terra = []string{"Høyeste temperatur målt på Jordens overflate er 56.7 C", "Laveste temperatur målt på Jordens overflate -89.4 C", "Temperatur i Jordens indre kjerne 9392 K"}

	if about == "sun" {
		return funFacts.Sun
	} else if about == "luna" {
		return funFacts.Luna
	} else {
		return funFacts.Terra
	}
}
