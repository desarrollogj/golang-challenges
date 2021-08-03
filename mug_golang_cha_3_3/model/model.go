package model

type Appointment struct {
	Id     int    `json:"id,omitempty" bson:"id,omitempty"`
	Title  string `json:"title,omitempty" bson:"title,omitempty"`
	IsDone bool   `json:"isDone" bson:"isDone"`
}

type GenericResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
