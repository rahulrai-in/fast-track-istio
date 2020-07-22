package models

import "time"

type Book struct {
	Id        int
	Date      time.Time
	Title     string
	Thumbnail string
	Author    string
	Genre     string
	Summary   string
}

type Movie struct {
	Id        int
	Date      time.Time
	Title     string
	Thumbnail string
	Director  string
	Genre     string
	Summary   string
}
