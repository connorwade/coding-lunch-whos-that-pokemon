package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Operation struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
}

type PokemonNameData struct {
	Data struct {
		Pokemon []struct {
			Name string `json:"name"`
		} `json:"pokemon_v2_pokemon"`
	} `json:"data"`
}

type PokemonDetailsData struct {
	Details PokemonDetails `json:"data"`
}

type PokemonDetails struct {
	Species []Species `json:"pokemon_v2_pokemonspecies"`
}

type Species struct {
	FlavorTexts []FlavorText `json:"flavorText"`
	Pokemon     Pokemon      `json:"pokemon"`
}

type Pokemon struct {
	Nodes []PokemonNodes `json:"nodes"`
}

type PokemonNodes struct {
	Height int     `json:"height"`
	Name   string  `json:"name"`
	Weight int     `json:"weight"`
	Stats  []Stats `json:"stats"`
	Types  []Types `json:"types"`
}

type Stats struct {
	Base int  `json:"base_stat"`
	Stat Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type Types struct {
	Slot     int      `json:"slot"`
	TypeName TypeName `json:"pokemon_v2_type"`
}

type TypeName struct {
	Name string `json:"name"`
}

type FlavorText struct {
	Text string `json:"flavor_text"`
}

func getAllPokemonNames() PokemonNameData {
	if _, err := os.Stat("names.json"); err == nil {
		//file exists
		content, err := os.ReadFile("names.json")
		if err != nil {
			log.Fatal(err)
		}

		var names PokemonNameData
		err = json.Unmarshal(content, &names)
		if err != nil {
			log.Fatal(err)
		}

		return names
	}
	pokemonNames := Operation{
		OperationName: "pokemon_names",
		Variables:     nil,
		Query: `query pokemon_names {
			pokemon_v2_pokemon(limit: 251) {
			  name
			}
		  }`,
	}

	var names PokemonNameData
	body := graphqlCall(pokemonNames)

	err := json.Unmarshal(body, &names)
	if err != nil {
		log.Fatal(err)
	}

	file, err := json.MarshalIndent(names, "", "")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("names.json", file, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return names
}

func getPokemonById(id int) PokemonDetailsData {
	pokemonDetails := Operation{
		OperationName: "pokemon_details",
		Variables: map[string]interface{}{
			"id": id,
		},
		Query: `query pokemon_details($id: Int) {
			pokemon_v2_pokemonspecies(where: {id: {_eq: $id}}) {
				flavorText: pokemon_v2_pokemonspeciesflavortexts(limit: 1) {
					flavor_text
				  }
			 pokemon: pokemon_v2_pokemons_aggregate(limit: 1){
			  nodes{
				height
				name
				weight
				stats: pokemon_v2_pokemonstats{
				  base_stat
				  stat: pokemon_v2_stat{
					name
				  }
				}
				types: pokemon_v2_pokemontypes{
					slot
					pokemon_v2_type{
					  name
					}
				  }
				
			  }
			}
			}
		  }`}

	// var pokemon Pokemon
	body := graphqlCall(pokemonDetails)

	var details PokemonDetailsData

	err := json.Unmarshal(body, &details)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(details)

	return details
}

func graphqlCall(operation Operation) []byte {
	url := "https://beta.pokeapi.co/graphql/v1beta"
	body, err := json.Marshal(operation)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(url, "", bytes.NewReader(body))
	if resp.StatusCode > 299 {
		log.Fatal(resp)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
