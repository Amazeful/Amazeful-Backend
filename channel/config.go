package channel

import (
	"github.com/Amazeful/Amazeful-Backend/rest"
	"github.com/Amazeful/dataful/store"
)

type ChannelConfig struct {
	responseWriter rest.ResponseWriter
	channelStore   store.ChannelStore
}

func NewChannelConfig(responseWriter rest.ResponseWriter, channelStore store.ChannelStore) *ChannelConfig {
	return &ChannelConfig{
		responseWriter: responseWriter,
		channelStore:   channelStore,
	}
}

// func (cc *ChannelConfig) HandleGetChannel(rw http.ResponseWriter, req *http.Request) {
// 	http.Redirect(rw, req, ta.oauth2.AuthCodeURL(ta.twitchConfig.State), http.StatusTemporaryRedirect)
// }
