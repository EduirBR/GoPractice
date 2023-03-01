package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"practice/pkg/dbconect"
	"practice/pkg/models"
	"practice/pkg/services"
)

func welcome(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	endpoint := fmt.Sprintln(`
	Estos son los endpoints de la API
	"/" inicio metodo GET
	"/persons"  agregar registro metodo POST
	"/persons"  optener registros metodo GET
	"/api/report" optener registros de personas en excel metodo GET`)
	fmt.Fprintln(rw, endpoint)

}
func postPerson(rw http.ResponseWriter, r *http.Request) {
	newPerson := models.Person{}
	json.NewDecoder(r.Body).Decode(&newPerson)

	if newPerson.Name != "" && newPerson.LastName != "" && newPerson.Dni != 0 && newPerson.Profession != "" {
		err := dbconect.PostPerson(newPerson)
		print(err)
		if err != nil {
			SendResponse(rw, http.StatusBadRequest, []byte(fmt.Sprintf(`"message":"%s"`, err.Error())))
			return
		}
		data, _ := json.MarshalIndent(newPerson, "", " ")
		SendResponse(rw, http.StatusCreated, data)
	} else {
		SendResponse(rw, http.StatusBadRequest, []byte(`{"message":"no puedes enviar campos vacios"}`))
	}

}

func getPerson(rw http.ResponseWriter, r *http.Request) {
	list, err := dbconect.GetPerson()
	if err != nil {
		SendResponse(rw, http.StatusInternalServerError, []byte(fmt.Sprintf(`"message":"%s"`, err.Error())))
		return
	}
	data, _ := json.Marshal(list)
	SendResponse(rw, http.StatusOK, data)
}

func rephandler(rw http.ResponseWriter, r *http.Request) {
	file_name := services.Repots()
	file, err := os.ReadFile(file_name)

	if err != nil {

		log.Fatal(err)
	}

	rw.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	rw.Write(file)

}

func SendResponse(rw http.ResponseWriter, codeStatus int, data []byte) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(codeStatus)
	rw.Write(data)
}
