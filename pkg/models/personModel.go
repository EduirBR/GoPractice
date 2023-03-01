package models

type Person struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Dni        int    `json:"dni"`
	Profession string `json:"profession"`
}

const PersonScheme = `CREATE TABLE IF NOT EXISTS persons(
	id SERIAL,
	name VARCHAR(15) not null,
	last_name VARCHAR(15) not null,
	dni INTEGER not null,
	profession VARCHAR(15) not null,
	PRIMARY KEY(id),
	UNIQUE (dni)	
)`