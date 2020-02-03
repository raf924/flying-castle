package model

import "time"

type FileKind uint8

const (
	RegularFile FileKind = 0
	Directory   FileKind = 1
)

type DataHolder struct {
	Data []byte
}

type File struct {
	FileSystemEntity
	DataHolder
	Size               uint64
	Parent             *File
	AccessTime         time.Time
	MetadataChangeTime time.Time
	DataChangeTime     time.Time
	Owner              *User
	Group              *Group
	Kind               FileKind
	UserPermissions    []EntityPermission
	GroupPermissions   []EntityPermission
}
