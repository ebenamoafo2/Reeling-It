package models

type Actor struct {
	ID        int
	FirstName string
	LastName  string
	ImageURL  *string // Optional field for actor's image URL
}
