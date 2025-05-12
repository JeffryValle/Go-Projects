package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("**** Consumo de API ****")

	// 1. Hacer la solicitud
	url := "https://pokeapi.co/api/v2/pokemon?offset=10&limit=10"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al obtener la respuesta ", err)
		return
	}
	defer resp.Body.Close()

	// 2. Obtener la respuesta
	body, er := io.ReadAll(resp.Body)
	if er != nil {
		fmt.Println("Error al obtener los datos", er)
		return
	}

	// 3. Decodificar el json (decode json) Creacion de modelos
	var f PokemonResponse
	msg := json.Unmarshal(body, &f)
	if msg != nil {
		fmt.Println("No se pudo parsear los datos", msg)
	}
	// 4. Muestreo de la data
	for i, t := range f.Results {
		fmt.Println("Pokemon", i+1, t.Name)
	}

	//------------------------------------------------------------
	fmt.Println("****** Servidor ******")

	router := gin.Default()

	// Cargar las plantillas HTML desde la carpeta "templates"
	router.LoadHTMLGlob("templates/*")

	// Ruta que renderiza la plantilla HTML
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":     "Mi primera página con Gin",
			"message":   "¡Hola desde Gin!",
			"Pokemones": f.Results,
		})
	})

	// Inicia el servidor
	router.Run(":8080")
}

// Creación de modelos
type PokemonResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Pokemon `json:"results"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
