package model

type Chunk struct {
	Id        int64
	Key       string
	NextChunk *int64
}
