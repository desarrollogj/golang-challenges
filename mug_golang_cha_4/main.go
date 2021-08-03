package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/desarrollogj/mug_golang_cha_4/model"
	"github.com/desarrollogj/mug_golang_cha_4/repository"
	"github.com/gin-gonic/gin"
)

var appointmentRepository repository.Repository

func getAllHandler(c *gin.Context) {
	c.JSON(http.StatusOK, findAll())
}

func getOneHandler(c *gin.Context) {
	id, err := getIdFromRequestParams(c)
	if err != nil {
		writeErrorMessage(c, http.StatusBadRequest, "id parameter has not a valid value")
		return
	}

	selectedAppointment, err := find(id)
	if err != nil {
		writeErrorMessage(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, selectedAppointment)
}

func createHandler(c *gin.Context) {
	newAppointment, reqErr := getAppointmentFromRequestBody(c)
	if reqErr != nil {
		writeErrorMessage(c, http.StatusBadRequest, "request body is not valid")
		return
	}

	createErr := create(newAppointment)
	if createErr != nil {
		writeErrorMessage(c, http.StatusBadRequest, createErr.Error())
		return
	}

	c.JSON(http.StatusCreated, newAppointment)
}

func updateHandler(c *gin.Context) {
	id, getIdErr := getIdFromRequestParams(c)
	if getIdErr != nil {
		writeErrorMessage(c, http.StatusBadRequest, "id parameter has not a valid value")
		return
	}

	updatedAppointment, bodyErr := getAppointmentFromRequestBody(c)
	if bodyErr != nil {
		writeErrorMessage(c, http.StatusBadRequest, "request body is not valid")
		return
	}

	if id != updatedAppointment.Id {
		writeErrorMessage(c, http.StatusBadRequest, "param id and body id must be the same")
		return
	}

	updateErr := update(updatedAppointment)
	if updateErr != nil {
		writeErrorMessage(c, http.StatusNotFound, updateErr.Error())
		return
	}

	c.JSON(http.StatusOK, updatedAppointment)
}

func deleteHandler(c *gin.Context) {
	id, getIdErr := getIdFromRequestParams(c)
	if getIdErr != nil {
		writeErrorMessage(c, http.StatusBadRequest, "id parameter has not a valid value")
		return
	}

	appointmentToDelete, deleteErr := delete(id)
	if deleteErr != nil {
		writeErrorMessage(c, http.StatusNotFound, deleteErr.Error())
		return
	}

	c.JSON(http.StatusOK, appointmentToDelete)
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

func getAppointmentFromRequestBody(c *gin.Context) (model.Appointment, error) {
	var bodyAppointment model.Appointment
	err := c.BindJSON(&bodyAppointment)
	if err != nil {
		return bodyAppointment, err
	}

	return bodyAppointment, nil
}

func getIdFromRequestParams(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func jsonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
	}
}

func writeErrorMessage(c *gin.Context, statusCode int, message string) {
	c.JSON(http.StatusCreated, model.GenericResponse{Code: statusCode, Message: message})
}

func main() {
	appointmentRepository = repository.MemoryRepository{}
	//appointmentRepository = repository.MongoRepository{}

	router := setupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("HTTP_PLATFORM_PORT")
		if port == "" {
			port = "8080"
			log.Printf("Defaulting to port %s", port)
		}
	}

	router.Run(fmt.Sprintf(":%s", port))
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(jsonMiddleware())

	api := router.Group("/api")
	api.GET("/", getAllHandler)
	api.GET("/:id", getOneHandler)
	api.POST("/", createHandler)
	api.PATCH("/:id", updateHandler)
	api.DELETE("/:id", deleteHandler)

	return router
}
