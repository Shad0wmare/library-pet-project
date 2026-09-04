package model

type Authors struct {
	Id        int     `json:"id" db:"author_id"`
	FirstName *string `json:"firstName" db:"first_name"`
	LastName  *string `json:"lastName" db:"last_name"`
}
