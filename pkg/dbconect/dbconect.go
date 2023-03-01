package dbconect

import (
	"database/sql"
	"fmt"
	"log"
	"practice/config"

	_ "github.com/lib/pq"
)

//BASE DE DATOS PRINCIPAL
//POSTGRES
func ConnectionMain() (db *sql.DB) {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalln("error al cargar los datos ")
	}
	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.User, conf.Password, conf.HostDB, conf.PortPosgresql, conf.DBName)
	connectionMain, err := sql.Open("postgres", uri)
	if err != nil {
		fmt.Println("error al conectar a la base de datos")
	}
	return connectionMain
}

func CreateTable(schm string) string{
	db := ConnectionMain()
	_,err := db.Exec(schm)
	if err!= nil{
		return err.Error()
	}
	return "ok"
}