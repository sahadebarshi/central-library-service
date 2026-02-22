package model

type BookUser struct {
	Memberid        int64
	Fname           string
	Lname           string
	MembershipType  string `binding:"required"`
	Email           string
	ContactNo       string
	PrefferedGenres string
	Status          string
	Startdate       CustomDate `json:"start_date"`
	Enddate         CustomDate `json:"end_date"`
}
