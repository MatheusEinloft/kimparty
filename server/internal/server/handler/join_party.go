package handler

import (
	"kimparty/internal/party"
	"kimparty/internal/server"
	"kimparty/pkg/convertion"
	"net/http"

	"github.com/lxzan/gws"
)

type JoinPartyHandler struct {
	partyService *party.Service
	upgrader     *gws.Upgrader
}

func NewJoinPartyHandler(partyService *party.Service, upgrader *gws.Upgrader) server.Handler {
	return &JoinPartyHandler{
		partyService: partyService,
        upgrader:     upgrader,
	}
}

func (h *JoinPartyHandler) Path() string {
	return "/ws/join"
}

func (h *JoinPartyHandler) Methods() string {
	return "GET"
}

func (h *JoinPartyHandler) Handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("party_id")
		username := r.URL.Query().Get("username")

        pt, newUser, err := h.partyService.PrepareForEntry(id, username)

        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write(convertion.StringToBytes(err.Error()))
            return
        }

        socket, err := h.upgrader.Upgrade(w, r)
        pt.AddConn(newUser.ID, socket)

        socket.Session().Store("party", pt)
        socket.Session().Store("user", newUser)

        socket.ReadLoop()

        h.partyService.RemovePartyIfEmpty(pt)
	}
}
