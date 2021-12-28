package models

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	SessionId       string             `json:"sessionId"`
	User            primitive.ObjectID `json:"user"`
	SelectedChannel primitive.ObjectID `json:"channel"`

	rc util.IRedis
}

func NewSession(rc util.IRedis) *Session {
	return &Session{rc: rc}
}

//GenerateSessionId generates a new session uid
func (s *Session) GenerateSessionId() {
	s.SessionId = xid.New().String()
}

//SetSession adds session to db
func (s *Session) SetSession(ctx context.Context, expiry time.Duration) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	status := s.rc.Set(ctx, s.key(), data, expiry)
	return status.Err()
}

//GetSession gets the session from db
func (s *Session) GetSession(ctx context.Context) error {
	bytes, err := s.rc.Get(ctx, s.key()).Bytes()
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) key() string {
	return "session-" + s.SessionId
}
