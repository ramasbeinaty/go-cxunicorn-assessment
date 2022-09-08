package listing

import (
	"clinicapp/pkg/storage/postgres"
	"errors"
	"time"
)

var ErrIdNotFound = errors.New("doctor with given id not found")
var ErrEmpty = errors.New("no doctor was found")

// provide access to the doctor storage
type Repository interface {
	// returns a doctor with given id
	GetDoctor(int) (postgres.Doctor, error)

	// returns all doctors in storage
	GetAllDoctors() []postgres.Doctor

	// returns all the appointments of a specific doctor
	GetAllAppointmentsOfDoctor(int, time.Time) []postgres.Appointment
}

// provide listing operations for struct doctor
type Service interface {
	GetDoctor(int) (Doctor, error)
	GetAllDoctors() []Doctor
	GetAllAppointmentsOfDoctor(int, time.Time) []Appointment
}

type service struct {
	repo Repository
}

// creates a listing service with the necessary dependencies
func NewService(repo Repository) Service {
	return &service{repo}
}

// implement service methods
func (s *service) GetDoctor(id int) (Doctor, error) {
	var d postgres.Doctor
	var doctor Doctor
	var err error

	d, err = s.repo.GetDoctor(id)

	doctor.ID = d.ID
	doctor.Email = d.Email
	doctor.FirstName = d.FirstName
	doctor.LastName = d.LastName
	doctor.Specialization = d.Specialization

	if err != nil {
		return doctor, errors.New("GetDoctor - " + err.Error())
	}

	return doctor, nil
}

func (s *service) GetAllDoctors() []Doctor {
	var _doctors []postgres.Doctor
	var doctors []Doctor = []Doctor{}

	_doctors = s.repo.GetAllDoctors()

	for _, _doctor := range _doctors {
		var doctor Doctor

		doctor.ID = _doctor.ID
		doctor.Email = _doctor.Email
		doctor.FirstName = _doctor.FirstName
		doctor.LastName = _doctor.LastName
		doctor.Specialization = _doctor.Specialization

		doctors = append(doctors, doctor)
	}

	return doctors
}

func (s *service) GetAllAppointmentsOfDoctor(doctor_id int, date time.Time) []Appointment {
	var _appointments []postgres.Appointment
	var appointments []Appointment = []Appointment{}

	_appointments = s.repo.GetAllAppointmentsOfDoctor(doctor_id, date)

	for _, _appointment := range _appointments {
		var appointment Appointment

		appointment.ID = _appointment.ID
		appointment.DoctorID = _appointment.DoctorID
		appointment.CreatedBy = _appointment.CreatedBy
		appointment.CreatedAt = _appointment.CreatedAt
		appointment.StartDatetime = _appointment.StartDatetime
		appointment.EndDatetime = _appointment.EndDatetime
	}

	return appointments
}
