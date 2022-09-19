package shortener

import (
	"salif.eu/go/hasher"
)

func MakeShort(url string) string {
	hash, _, _ := hasher.Hash(url)
	return hash
}
