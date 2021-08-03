package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type appointment struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

type genericResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var appointments []appointment

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(appointments)
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

func find(id int) (appointment, error) {
	var selectedAppointment appointment
	for _, item := range appointments {
		if item.Id == id {
			selectedAppointment = item
			break
		}
	}

	if selectedAppointment.Id == 0 {
		err := fmt.Errorf("appointment with id %d not found", id)
		return selectedAppointment, err
	}

	return selectedAppointment, nil
}

func create(item appointment) error {
	_, findErr := find(item.Id)
	if findErr == nil {
		err := fmt.Errorf(fmt.Sprintf("an appointment with id %d already exists", item.Id))
		return err
	}

	appointments = append(appointments, item)

	return nil
}

func update(item appointment) error {
	_, findErr := find(item.Id)
	if findErr != nil {
		return findErr
	}

	appointmentsToPreserve := make([]appointment, 0)
	for _, singleAppointment := range appointments {
		if singleAppointment.Id != item.Id {
			appointmentsToPreserve = append(appointmentsToPreserve, singleAppointment)
		}
	}

	appointmentsToPreserve = append(appointmentsToPreserve, item)
	appointments = appointmentsToPreserve

	return nil
}

func delete(id int) (appointment, error) {
	appointmentToDelete, findErr := find(id)
	if findErr != nil {
		return appointmentToDelete, findErr
	}

	appointmentsToPreserve := make([]appointment, 0)
	for _, item := range appointments {
		if item.Id != appointmentToDelete.Id {
			appointmentsToPreserve = append(appointmentsToPreserve, item)
		}
	}
	appointments = appointmentsToPreserve

	return appointmentToDelete, nil
}

func getAppointmentFromRequestBody(r *http.Request) (appointment, error) {
	var bodyAppointment appointment
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
	errResponse := genericResponse{Code: statusCode, Message: message}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errResponse)
}

func main() {
	appointments = make([]appointment, 0)
	toDo1 := appointment{Id: 1, Title: "ToDo Task #1", IsDone: false}
	toDo2 := appointment{Id: 2, Title: "ToDo Task #2", IsDone: true}
	appointments = append(appointments, toDo1)
	appointments = append(appointments, toDo2)

	router := mux.NewRouter().StrictSlash(true)
	router.Use(jsonMiddleware)

	router.HandleFunc("/api", createHandler).Methods("POST")
	router.HandleFunc("/api", getAllHandler).Methods("GET")
	router.HandleFunc("/api/{id}", getOneHandler).Methods("GET")
	router.HandleFunc("/api/{id}", updateHandler).Methods("PATCH")
	router.HandleFunc("/api/{id}", deleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
