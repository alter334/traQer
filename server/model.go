package main

type User struct {
	UserName string `json:"name" db:"name"`
	UUID     string `json:"uuid" db:"id"`
	Attack   int    `json:"attack" db:"attack"`
	Rate     int    `json:"rate" db:"rate"`
}


