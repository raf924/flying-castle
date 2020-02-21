package model

import "time"

type StorageKey struct {
	Key       []byte
	CreatedAt time.Time
}

type StorageKeyBusiness interface {
	Create() error
	GetLatest()
}
