package websocket

import (
	"github.com/lxzan/gws"
)

func NewUpgrader() *gws.Upgrader {
	handler := new(WebSocket)

	return gws.NewUpgrader(handler, &gws.ServerOption{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,

		PermessageDeflate: gws.PermessageDeflate{
			Enabled:               true,
			ServerContextTakeover: false,
			ClientContextTakeover: false,
		},
	})
}
