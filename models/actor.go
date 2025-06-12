package models

type Actor struct {
	ID        int
	FirstName string
	LastName  string
	ImageURL  *string // nullable
	Name      string  // full name, derived from FirstName and LastName
}
