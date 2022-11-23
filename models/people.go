package models

type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

var People []Person
