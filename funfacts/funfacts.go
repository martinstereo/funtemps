package funfacts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra
*/

type FunFacts struct {
	sun   []string
	luna  []string
	terra []string
}

func GetFunFacts(about string) []string {

	funf := FunFacts{
		sun: []string{
			"Temperaturen i Solens kjerne er",
			"Temperatur på ytre lag av Solen er",
		},
		luna: []string{
			"Temperatur på Månens overflate om natten er",
			"Temperatur på Månens overflate om dagen er",
		},
		terra: []string{
			"Høyeste temperatur målt på Jordens overflate er",
			"Laveste temperatur målt på Jordens overflate er",
		},
	}

	if about == "sun" {
		return funf.sun
	}
	if about == "luna" {
		return funf.luna
	}
	if about == "terra" {
		return funf.terra
	} else {
		return []string{} // return nothing if none match
	}
}
