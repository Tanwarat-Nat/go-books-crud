package main

type Book struct {
	ID      string   `json:"id"`
	Isbn    string   `json:"isbn"`
	Title   string   `json:"title"`
	Authors *Authors `json:"authors"`
}

type Authors struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
