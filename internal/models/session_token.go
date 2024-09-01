package models

type SessionToken struct {
	UserEmail string `json:"email"`
	FullName  string `json:"full_name"`
}
