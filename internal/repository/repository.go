package repository

import "bed_brkfst/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	
	InsertReservation(res models.Reservation) error
	// InsertRoomRestriction(r models.RoomRestriction) error
}
