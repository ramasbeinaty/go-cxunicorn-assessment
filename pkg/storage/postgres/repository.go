package postgres

import (
	"clinicapp/pkg/listing"
	"errors"
	"fmt"
	"log"

	"database/sql"

	"github.com/blockloop/scan"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

// start and returns a new DB
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	s.DB, err = sql.Open("postgres", "postgres://postgres:123456@localhost:5432/clinic_app2?sslmode=disable")

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// defer s.DB.Close()

	s.DB.SetConnMaxLifetime(0)
	s.DB.SetMaxOpenConns(3)

	return s, nil
}

// list doctor with given id
func (s *Storage) GetDoctor(id string) (listing.Doctor, error) {
	// if row := s.DB.QueryRow(`
	// SELECT first_name, last_name, email, work_shift, specialization, doctors.id from users, staffs, doctors
	// WHERE users.id = $1 AND staffs.id = $1 AND doctors.id = $1`, id); row != nil {
	// 	if err := row.Scan(&doctor.FirstName, &doctor.LastName, &doctor.Email,
	// 		&doctor.WorkShift, &doctor.Specialization, &doctor.ID); err != nil {
	// 		return doctor, errors.New(fmt.Sprintln("ERROR: GetDoctor - ", err))
	// 	}
	// 	if row.Err() == sql.ErrNoRows {
	// 		return doctor, errors.New(fmt.Sprintln("ERROR: GetDoctor - ", listing.ErrIdNotFound))
	// 	}
	// 	return doctor, errors.New(fmt.Sprintln("ERROR: GetDoctor - ", row))
	// }

	// row := s.DB.QueryRow(`
	// SELECT * FROM users, staffs, doctors
	// WHERE users.id = $1 AND staffs.id = $1 AND doctors.id = $1`, id)

	// scanning := structscanner.Select(s.DB, &d, "",
	// 	`SELECT * FROM users, staffs, doctors
	// 	WHERE users.id = 1 AND staffs.id = 1 AND doctors.id = 1`)

	var doctor listing.Doctor

	row, _ := s.DB.Query(`SELECT * FROM users, staffs, doctors 
							WHERE users.id = $1 AND staffs.id = $1 AND doctors.id = $1`, id)

	if err := scan.RowStrict(&doctor, row); err != nil {
		if err == sql.ErrNoRows {
			return doctor, errors.New(fmt.Sprintln("ERROR: GetDoctor - ", listing.ErrIdNotFound))
		}
		return doctor, errors.New(fmt.Sprintln("ERROR: GetDoctor - ", err))

	}

	return doctor, nil
}

func (s *Storage) GetAllDoctors() []listing.Doctor {
	var doctors []listing.Doctor = []listing.Doctor{}

	// rows, _ := s.DB.Query(`SELECT * FROM doctors, staffs, users
	// 					   WHERE staffs.id = doctors.id AND users.id = doctors.id
	// 						`)

	rows, _ := s.DB.Query(`SELECT * FROM doctors
						 	JOIN staffs ON doctors.id = staffs.id
							JOIN users ON doctors.id = users.id
							`)

	if err := scan.RowsStrict(&doctors, rows); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Warning: GetAllDoctors", listing.ErrEmpty)
			return doctors
		}

		fmt.Println("ERROR: GetAllDoctors - ", err)
		return doctors
	}

	return doctors

}
