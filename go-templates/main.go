package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Answer struct {
	Guess string `json:"guess"`
}

type Response struct {
	WinLose string `json:"winlose"`
}

type ServedPokemon struct {
	Name        string    `json:"name"`
	FlavorText  string    `json:"flavorText"`
	Weight      int       `json:"weight"`
	Height      int       `json:"height"`
	HasTwoTypes bool      `json:"hasTwoTypes"`
	Types       []string  `json:"types"`
	Stats       LongStats `json:"stats"`
}

type LongStats struct {
	HP    int `json:"hp"`
	ATK   int `json:"atk"`
	DEF   int `json:"def"`
	SPATK int `json:"spAtk"`
	SPDEF int `json:"spDef"`
	SPD   int `json:"spd"`
}

var (
	servedPokemon ServedPokemon
)

func main() {

	id := randomInt()
	pokemon := getPokemonById(id)

	servedPokemon.Name = pokemon.Details.Species[0].Pokemon.Nodes[0].Name
	servedPokemon.FlavorText = pokemon.Details.Species[0].FlavorTexts[0].Text
	servedPokemon.Height = pokemon.Details.Species[0].Pokemon.Nodes[0].Height
	servedPokemon.Weight = pokemon.Details.Species[0].Pokemon.Nodes[0].Weight
	for _, t := range pokemon.Details.Species[0].Pokemon.Nodes[0].Types {
		servedPokemon.Types = append(servedPokemon.Types, t.TypeName.Name)
	}
	servedPokemon.Stats.HP = pokemon.Details.Species[0].Pokemon.Nodes[0].Stats[0].Base
	servedPokemon.Stats.ATK = pokemon.Details.Species[0].Pokemon.Nodes[0].Stats[1].Base
	servedPokemon.Stats.DEF = pokemon.Details.Species[0].Pokemon.Nodes[0].Stats[2].Base
	servedPokemon.Stats.SPATK = pokemon.Details.Species[0].Pokemon.Nodes[0].Stats[3].Base
	servedPokemon.Stats.SPDEF = pokemon.Details.Species[0].Pokemon.Nodes[0].Stats[4].Base
	servedPokemon.Stats.SPD = pokemon.Details.Species[0].Pokemon.Nodes[0].Stats[5].Base

	fmt.Println(pokemon.Details.Species[0].Pokemon.Nodes[0].Types, servedPokemon)

	r := gin.Default()

	r.Use(cors.Default())
	r.Static("/js", "./js")

	r.LoadHTMLGlob("templates/*")
	Router(r)
	log.Println("Server Started")
	r.Run()
}

func Router(r *gin.Engine) {
	r.GET("/", app)
	r.POST("/answer", name)
}

func name(c *gin.Context) {
	var newGuess Answer
	var response Response

	err := c.BindJSON(&newGuess)
	if err != nil {
		log.Fatalln("error setting json in name ", err)
	}

	if newGuess.Guess == servedPokemon.Name {
		response.WinLose = "win"
		c.JSON(http.StatusOK, response)
	} else {
		response.WinLose = "lose"
		c.JSON(http.StatusOK, response)
	}
}

func app(c *gin.Context) {
	names := getAllPokemonNames()
	var n []string

	for _, val := range names.Data.Pokemon {
		n = append(n, val.Name)
	}
	c.HTML(
		http.StatusOK,
		"app.html",
		gin.H{
			"names":      n,
			"flavorText": servedPokemon.FlavorText,
			"height":     servedPokemon.Height,
			"weight":     servedPokemon.Weight,
			"types":      servedPokemon.Types,
			"hp":         servedPokemon.Stats.HP,
			"atk":        servedPokemon.Stats.ATK,
			"def":        servedPokemon.Stats.DEF,
			"spAtk":      servedPokemon.Stats.SPATK,
			"spDef":      servedPokemon.Stats.SPDEF,
			"spd":        servedPokemon.Stats.SPD,
		},
	)
}

func randomInt() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 251
	randInt := rand.Intn(max-min+1) + min
	return randInt
}
