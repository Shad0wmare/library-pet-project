package model

type Authors struct {
	Id        int     `json:"id" db:"author_id"`
	FirstName *string `json:"firstName" db:"first_name"`
	LastName  *string `json:"lastName" db:"last_name"`
}

type CreateAuthor struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type UpdateAuthor struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}
