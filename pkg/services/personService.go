package services

import (
	"fmt"
	"log"
	"practice/config"
	"practice/pkg/dbconect"
	"time"

	"github.com/xuri/excelize/v2"
)

func Repots() string {

	list, _ := dbconect.GetPerson()
	t := time.Now()
	file_name := ""
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "Person Table")
	f.MergeCell("Sheet1", "A1", "B2")
	f.SetSheetRow("Sheet1", "A3", &[]interface{}{"", "Name", "Last Name", "DNI", "Profession"})
	for i, person := range list {
		cell, _ := excelize.CoordinatesToCellName(1, i+4)
		f.SetSheetRow("Sheet1", cell, &[]interface{}{"", person.Name, person.LastName, person.Dni, person.Profession})
	}
	conf, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return file_name
	}
	f.SetDocProps(&excelize.DocProperties{Creator: conf.Creator})

	file_name = fmt.Sprintf("%s-%d%d%d-%d%d%d.xlsx", conf.FileBaseName, t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
	if err := f.SaveAs(file_name); err != nil {
		log.Fatal(err)
	}
	return file_name

}
