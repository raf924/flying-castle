package model

type Group struct {
	FileSystemEntity
	Users []*User
}
