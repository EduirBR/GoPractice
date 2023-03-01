package dbconect

import (
	"errors"
	"fmt"
	"practice/pkg/models"
)

func PostPerson(person models.Person) error {

	sqline := `INSERT INTO persons 
			(name, last_name, dni, profession) VALUES ($1,$2,$3,$4)`

	db := ConnectionMain()
	defer db.Close()

	stmt, err := db.Prepare(sqline)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(person.Name, person.LastName, person.Dni, person.Profession)
	if err != nil {
		return err
	}
	i, _ := result.RowsAffected()
	if i != 1 {
		mens := fmt.Sprintf("se esperaba solo una fila afectada y no %d", i)
		return errors.New(mens)
	}
	return nil
}

func GetPerson() ([]models.Person, error) {

	list := []models.Person{}
	db := ConnectionMain()
	defer db.Close()
	sqline := "SELECT * FROM persons"
	row, err := db.Query(sqline)
	if err != nil {
		return list, err
	} else {
		for row.Next() {
			i := models.Person{}
			row.Scan(&i.Id, &i.Name, &i.LastName, &i.Dni, &i.Profession)
			list = append(list, i)
		}
	}
	return list, err
}
