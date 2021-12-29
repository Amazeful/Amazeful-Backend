package models

import (
	"testing"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {

	type args struct {
		signKey    string
		expiration time.Time
		sessionId  string
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
	}{
		{"valid", false, args{"test", time.Now(), "123"}},
		{"valid2", false, args{"test", time.Now(), "321"}},
		{"invalid sign key", true, args{"", time.Now(), "321"}},
		{"invalid time", true, args{"", time.Now().Add(-time.Hour), "321"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			j := NewJWT([]byte(test.args.signKey), jwa.HS256)

			tokenString, err := j.Encode(test.args.sessionId, test.args.expiration)
			if test.wantErr {
				assert.Error(t, err)
				assert.Empty(t, tokenString)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, tokenString)

				token, err := jwt.Parse([]byte(tokenString))
				assert.NoError(t, err)
				assert.Equal(t, test.args.expiration.Unix(), token.Expiration().Unix())
				sid, ok := token.Get(SessionIdKey)
				assert.True(t, ok)
				assert.Equal(t, test.args.sessionId, sid)
			}

		})
	}

}
