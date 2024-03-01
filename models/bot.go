package models

type Message struct {
	Author  Author
	Content string
}

type Author struct {
	ID       string
	Email    string
	Locale   string
	Username string
	Verified bool
}
