package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/desarrollogj/mug_golang_cha_3_3/model"
	"github.com/desarrollogj/mug_golang_cha_3_3/repository"
	"github.com/gorilla/mux"
)

var appointmentRepository repository.Repository

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(findAll())
}

func getOneHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequestParams(r)
	if err != nil {
		writeErrorMessage(w, http.StatusBadRequest, "id parameter has not a valid value")
		return
	}

	selectedAppointment, err := find(id)
	if err != nil {
		writeErrorMessage(w, http.StatusNotFound, err.Error())
		return
	}

	json.NewEncoder(w).Encode(selectedAppointment)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	newAppointment, reqErr := getAppointmentFromRequestBody(r)
	if reqErr != nil {
		writeErrorMessage(w, http.StatusBadRequest, "request body is not valid")
		return
	}

	createErr := create(newAppointment)
	if createErr != nil {
		writeErrorMessage(w, http.StatusBadRequest, createErr.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAppointment)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	id, getIdErr := getIdFromRequestParams(r)
	if getIdErr != nil {
		writeErrorMessage(w, http.StatusBadRequest, "id parameter has not a valid value")
		return
	}

	updatedAppointment, bodyErr := getAppointmentFromRequestBody(r)
	if bodyErr != nil {
		writeErrorMessage(w, http.StatusBadRequest, "request body is not valid")
		return
	}

	if id != updatedAppointment.Id {
		writeErrorMessage(w, http.StatusBadRequest, "param id and body id must be the same")
		return
	}

	updateErr := update(updatedAppointment)
	if updateErr != nil {
		writeErrorMessage(w, http.StatusNotFound, updateErr.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedAppointment)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, getIdErr := getIdFromRequestParams(r)
	if getIdErr != nil {
		writeErrorMessage(w, http.StatusBadRequest, "id parameter has not a valid value")
		return
	}

	appointmentToDelete, deleteErr := delete(id)
	if deleteErr != nil {
		writeErrorMessage(w, http.StatusNotFound, deleteErr.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(appointmentToDelete)
}

func findAll() []model.Appointment {
	return appointmentRepository.FindAll()
}

func find(id int) (model.Appointment, error) {
	selectedAppointment := appointmentRepository.Find(id)

	if selectedAppointment.Id == 0 {
		err := fmt.Errorf("appointment with id %d not found", id)
		return selectedAppointment, err
	}

	return selectedAppointment, nil
}

func create(item model.Appointment) error {
	_, findErr := find(item.Id)
	if findErr == nil {
		return fmt.Errorf("an appointment with id %d already exists", item.Id)
	}

	appointmentRepository.Create(item)

	return nil
}

func update(item model.Appointment) error {
	_, findErr := find(item.Id)
	if findErr != nil {
		return findErr
	}

	appointmentRepository.Update(item)

	return nil
}

func delete(id int) (model.Appointment, error) {
	appointmentToDelete, findErr := find(id)
	if findErr != nil {
		return appointmentToDelete, findErr
	}

	appointmentRepository.Delete(id)

	return appointmentToDelete, nil
}

func getAppointmentFromRequestBody(r *http.Request) (model.Appointment, error) {
	var bodyAppointment model.Appointment
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return bodyAppointment, err
	}

	err = json.Unmarshal(requestBody, &bodyAppointment)
	if err != nil {
		return bodyAppointment, err
	}

	return bodyAppointment, nil
}

func getIdFromRequestParams(r *http.Request) (int, error) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return 0, err
	}
	return id, nil
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func writeErrorMessage(w http.ResponseWriter, statusCode int, message string) {
	errResponse := model.GenericResponse{Code: statusCode, Message: message}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errResponse)
}

func main() {
	//appointmentRepository = repository.MemoryRepository{}
	appointmentRepository = repository.MongoRepository{}

	router := mux.NewRouter().StrictSlash(true)
	router.Use(jsonMiddleware)

	router.HandleFunc("/api", createHandler).Methods("POST")
	router.HandleFunc("/api", getAllHandler).Methods("GET")
	router.HandleFunc("/api/{id}", getOneHandler).Methods("GET")
	router.HandleFunc("/api/{id}", updateHandler).Methods("PATCH")
	router.HandleFunc("/api/{id}", deleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
