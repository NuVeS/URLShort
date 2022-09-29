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

type NewLinkRequest struct {
	Url string `json:"url"`

	Beauty string `json:"beauty,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Name string `json:"name"`

	Password string `json:"password"`
}

type LinkResponse struct {
	Url string `json:"url"`

	ShortUrl string `json:"shortUrl"`
}

type LinkRequest struct {
	Url string `json:"url"`
}

type LinkListResponse struct {
	List []LinkResponse `json:"list"`
}
