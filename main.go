package main

import (
	"assignment_2_3/database"
	"assignment_2_3/routers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Status struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	port := os.Getenv("PORT")
	database.StartDB()
	r := routers.StartApp()

	r.GET("/random", random) // Assignment 3

	r.Run(fmt.Sprintf(":%s", port))
}

func random(c *gin.Context) {
	for {
		time.Sleep(1 * time.Second)
		jsonFile, err := os.Open("status.json")
		if err != nil {
			fmt.Println(err)
		}

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var x Status
		json.Unmarshal(byteValue, &x)

		x.Water = rand.Intn(100)
		x.Wind = rand.Intn(100)

		if x.Water < 5 {
			x.WaterStatus = "Aman"
		} else if x.Water <= 8 && x.Water >= 6 {
			x.WaterStatus = "Siaga"
		} else if x.Water > 8 {
			x.WaterStatus = "Bahaya"
		}

		if x.Wind < 6 {
			x.WindStatus = "Aman"
		} else if x.Wind <= 15 && x.Wind >= 7 {
			x.WindStatus = "Siaga"
		} else if x.Wind > 15 {
			x.WindStatus = "Bahaya"
		}

		tpl, err := template.ParseFiles("index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		tpl.Execute(c.Writer, x)
	}
}
