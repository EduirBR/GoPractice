package main

import (
	"fmt"
	"log"
	"net/http"
	"practice/config"
	"practice/pkg/dbconect"
	"practice/pkg/handlers"
	"practice/pkg/models"
)

func main() {

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalln("error load settings ")
	}
	db := dbconect.ConnectionMain()
	fmt.Println(dbconect.CreateTable(models.PersonScheme))
	db.Close()

	fmt.Println("API Personas")
	fmt.Println("server online Port:", conf.Port)
	http.ListenAndServe(":"+conf.Port, handlers.GetRoutes())

}
