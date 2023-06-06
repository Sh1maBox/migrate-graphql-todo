package gqltodo

import "github.com/google/uuid"

func MakeUUID(data []byte) string {
	uuidSpace, _ := uuid.NewUUID()
	uuid := uuid.NewSHA1(uuidSpace, data).String()
	return uuid
}
