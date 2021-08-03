package repository

import "github.com/desarrollogj/mug_golang_cha_3_3/model"

type Repository interface {
	FindAll() []model.Appointment
	Find(id int) model.Appointment
	Create(item model.Appointment) model.Appointment
	Update(item model.Appointment) model.Appointment
	Delete(id int)
}
