package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Amazeful/Amazeful-Backend/rest"
	"github.com/Amazeful/dataful/models"
	"github.com/Amazeful/dataful/store"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

type JWTAuthenticator struct {
	sessionStore   store.SessionStore
	responseWriter rest.ResponseWriter
	signKey        []byte
	algorithm      jwa.SignatureAlgorithm
	expiry         time.Time
	cookieName     string
	sessionKey     string
}

func NewJWTAuthenticator(sessionStore store.SessionStore, signKey []byte) *JWTAuthenticator {
	return &JWTAuthenticator{
		sessionStore: sessionStore,
		signKey:      signKey,
		expiry:       time.Now().Add(24 * time.Hour),
		algorithm:    jwa.HS256,
		cookieName:   "tokenful",
		sessionKey:   "sid",
	}
}

//CreateToken creates a new JWT and writes the cookie.
func (a *JWTAuthenticator) CreateToken(ctx context.Context, rw http.ResponseWriter, userId, channelId string) error {
	session, err := a.setSession(ctx, userId, channelId)
	if err != nil {
		return err
	}

	//Make a new jwt for token
	tokenString, err := a.encode(session.ID)
	if err != nil {
		return err
	}

	//Add in the token to cookie
	cookie := &http.Cookie{
		Name:     a.cookieName,
		Value:    tokenString,
		Expires:  a.expiry,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	a.responseWriter.WriteCookie(rw, cookie)

	return nil
}

//Authenticate middleware authenticates requests
func (a *JWTAuthenticator) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//Get the token from cookie
		cookie, err := req.Cookie(a.cookieName)
		if err != nil {
			a.responseWriter.WriteError(rw, http.StatusUnauthorized, err, rest.MessageUnauthorized)
			return
		}

		sessionId, err := a.verify(cookie.Value)
		if err != nil {
			a.responseWriter.WriteError(rw, http.StatusUnauthorized, err, rest.MessageUnauthorized)
			return
		}

		session, err := a.sessionStore.GetBySessionId(req.Context(), sessionId)
		if err != nil {
			a.responseWriter.WriteError(rw, http.StatusInternalServerError, err, rest.MessageInternalServerError)
			return
		}

		ctx := context.WithValue(req.Context(), rest.CtxSession, session)

		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}

//encode creates and encodes a new jwt, signs the token and returns token string.
func (a *JWTAuthenticator) encode(sessionId string) (string, error) {
	t, err := jwt.NewBuilder().
		IssuedAt(time.Now().UTC()).
		Expiration(a.expiry).
		Claim(a.sessionKey, sessionId).
		Build()
	if err != nil {
		return "", err
	}
	payload, err := jwt.Sign(t, a.algorithm, a.signKey)
	if err != nil {
		return "", err
	}

	return string(payload), nil
}

//setSession creates a new session and stores it in the storage.
func (a *JWTAuthenticator) setSession(ctx context.Context, userId, channelId string) (*models.Session, error) {

	//Make a new session for user
	session := models.NewSession()
	session.User = userId
	session.SelectedChannel = channelId

	//set the new session
	err := a.sessionStore.Create(ctx, session)
	return session, err
}

//verify parses a token using provided token string, validates the token and returns session id from the token.
func (a *JWTAuthenticator) verify(tokenString string) (string, error) {
	t, err := jwt.Parse([]byte(tokenString), jwt.WithVerify(a.algorithm, a.signKey))
	if err != nil {
		return "", err
	}

	if t == nil {
		return "", errors.New("failed to parse token")
	}

	if err := jwt.Validate(t); err != nil {
		return "", err
	}

	sid, ok := t.Get(a.sessionKey)
	if !ok {
		return "", errors.New("failed to parse token")

	}

	sessionId, ok := sid.(string)
	if !ok {
		return "", errors.New("failed to cast sid to string")

	}

	return sessionId, nil
}
