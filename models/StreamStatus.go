package models

type StreamStatus int

const (
	StreamLive StreamStatus = 1 << iota
	StreamOffline
)
