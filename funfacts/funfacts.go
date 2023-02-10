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

func GetFunFacts(sun, luna, terra string) []string {

	funfact := []FunFacts{
		{sun: ["hello"], luna: "ss"},
	}

	return
}
