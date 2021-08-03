package repository

import (
	"github.com/desarrollogj/mug_golang_cha_3_3/model"
)

var appointments = []model.Appointment{}

type MemoryRepository struct {
}

func (m MemoryRepository) FindAll() []model.Appointment {
	return appointments
}

func (m MemoryRepository) Find(id int) model.Appointment {
	var selectedAppointment model.Appointment
	for _, item := range appointments {
		if item.Id == id {
			selectedAppointment = item
			break
		}
	}

	return selectedAppointment
}

func (m MemoryRepository) Create(item model.Appointment) model.Appointment {
	appointments = append(appointments, item)

	return item
}

func (m MemoryRepository) Update(item model.Appointment) model.Appointment {
	appointmentsToPreserve := make([]model.Appointment, 0)
	for _, singleAppointment := range appointments {
		if singleAppointment.Id != item.Id {
			appointmentsToPreserve = append(appointmentsToPreserve, singleAppointment)
		}
	}

	appointmentsToPreserve = append(appointmentsToPreserve, item)
	appointments = appointmentsToPreserve

	return item
}

func (m MemoryRepository) Delete(id int) {
	appointmentsToPreserve := make([]model.Appointment, 0)
	for _, item := range appointments {
		if item.Id != id {
			appointmentsToPreserve = append(appointmentsToPreserve, item)
		} 
	}

	appointments = appointmentsToPreserve
}
