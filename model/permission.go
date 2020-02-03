package model

type FileSystemEntity struct {
	Id   uint64
	Name string
}

type Permission uint8

const (
	Read      Permission = 0b00000010
	Write     Permission = 0b00000100
	ReadWrite Permission = Read & Write
)

type EntityPermission struct {
	Entity     *FileSystemEntity
	Permission *Permission
}
