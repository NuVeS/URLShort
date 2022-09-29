package storage

import (
	"github.com/NuVeS/URLShort/cmd/models"
)

var users = make([]*models.User, 10)

var tokens = make(map[string]*models.User)

var linkList = make(models.LinkList)

type StorageAPI interface {
	CreateUser(name string, password string) bool
	DeleteUser(user *models.User) bool
	GetUser(name string) (bool, *models.User)
	ListLinks(token string) []*models.Link
	AddLink(token string, url string, shortLink string) bool
	DeleteLink(token string, shortLink string) bool
	GetUserByToken(token string) *models.User
	SetToken(token string, user *models.User) bool
	IsAvailable(link string) bool
	GetURL(shortLink string) string
}

type StorageDB struct{}

var Storage = StorageDB{}

func (strg *StorageDB) CreateUser(name string, password string) bool {
	newUser := &models.User{Name: name, Password: password}
	users = append(users, newUser)
	return true
}

func (strg *StorageDB) DeleteUser(user *models.User) bool {
	user = nil
	return true
}

func (strg *StorageDB) GetUser(name string) (bool, *models.User) {
	for _, tempUser := range users {
		if name == tempUser.Name {
			return true, tempUser
		}
	}
	return false, nil
}

func (strg *StorageDB) ListLinks(token string) []*models.Link {
	return linkList[token]
}

func (strg *StorageDB) AddLink(token string, url string, shortLink string) bool {
	link := &models.Link{Url: url, Short: shortLink}
	linkList[token] = append(linkList[token], link)
	return true
}

func (strg *StorageDB) DeleteLink(token string, shortLink string) bool {
	userLinks := linkList[token]
	for i, item := range userLinks {
		if item.Short == shortLink {
			remove(userLinks, i)
			break
		}
	}
	return true
}

func (strg *StorageDB) GetUserByToken(token string) *models.User {
	return tokens[token]
}

func (strg *StorageDB) SetToken(token string, user *models.User) bool {
	tokens[token] = user
	return true
}

func (strg *StorageDB) IsAvailable(link string) bool {
	_, occupied := linkList[link]
	return !occupied
}

func (strg *StorageDB) GetURL(shortLink string) string {
	for _, value := range linkList {
		for _, item := range value {
			if item.Short == shortLink {
				return item.Url
			}
		}
	}

	return ""
}

func remove(slice []*models.Link, s int) []*models.Link {
	return append(slice[:s], slice[s+1:]...)
}
