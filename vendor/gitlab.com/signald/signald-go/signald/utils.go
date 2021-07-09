package signald

import (
	"math/rand"

	"gitlab.com/signald/signald-go/signald/client-protocol/v0"
)

const idsize = 10

var charset = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// GenerateID is a helper function to generate random request IDs.
func GenerateID() string {
	id := make([]rune, idsize)
	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
	}
	return string(id)
}

func GetLegacyResponse(c chan v0.LegacyResponse, id string) v0.LegacyResponse {
	for {
		message := <-c
		if message.ID == id {
			return message
		}
	}
}
