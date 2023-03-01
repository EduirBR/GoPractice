package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practice/pkg/dbconect"
	"practice/pkg/models"
)

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

func deletePerson(rw http.ResponseWriter, r *http.Request) {

}
func putPerson(rw http.ResponseWriter, r *http.Request) {

}
func patchPerson(rw http.ResponseWriter, r *http.Request) {

}

func SendResponse(rw http.ResponseWriter, codeStatus int, data []byte) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(codeStatus)
	rw.Write(data)
}
