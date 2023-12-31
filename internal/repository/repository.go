package repository

type DatabaseRepo interface {
	AllUsers() bool
	
	InsertReservation(res models.Reservation) (int, error)
}
