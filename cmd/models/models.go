package models

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	// Links    map[string]string `json:"links"`
}

type Link struct {
	Url   string
	Short string
}

type LinkList = map[string][]*Link
