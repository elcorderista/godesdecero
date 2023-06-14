package mapas

import "fmt"

func MostrarMapas() {
	campeonatos := make(map[string]string)
	fmt.Println(campeonatos)

	paises := map[string]string{
		"Afganistan":     "Kabul",
		"Arabia Saudita": "Riad",
		"Armenia":        "Erevan",
		"Azerbaiyan":     "Baku",
		"Bangladesh":     "Dacca",
		"Barein":         "Manama",
		"Birmania":       "Naipyido",
	}

	for clave, valor := range paises {
		fmt.Println(clave + " - " + valor)
	}

	//Validate how many countries
	fmt.Printf("Cantidad de Paises: %d \n", len(paises))
	fmt.Println(paises)
	//Delete a component
	delete(paises, "Afganistan")

	fmt.Printf("Cantidad de Paises: %d \n", len(paises))
	fmt.Println(paises)

	//Validate if a country exist
	capital, exist := paises["Mexico"]
	fmt.Printf("Existe la capital %s en tu lista: %t \n", capital, exist)

	//Add a new elemnt
	paises["Mexico"] = "Ciudad de Mexico"
	fmt.Printf("Cantidad de Paises: %d \n", len(paises))
	fmt.Println(paises)
	capital, exist = paises["Mexico"]
	fmt.Printf("Existe la capital %s en tu lista: %t \n", capital, exist)
}
