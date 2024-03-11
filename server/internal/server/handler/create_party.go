package handler

import (
	"kimparty/internal/party"
	"kimparty/internal/server"
	"kimparty/pkg/convertion"
	"net/http"
)

type CreatePartyHandler struct {
	partyService *party.Service
}

func NewCreatePartyHandler(partyService *party.Service) server.Handler {
	return &CreatePartyHandler{
		partyService: partyService,
	}
}

func (h *CreatePartyHandler) Path() string {
	return "/party"
}

func (h *CreatePartyHandler) Methods() string {
	return "POST"
}

func (h *CreatePartyHandler) Handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		capacity := r.URL.Query().Get("capacity")

		partyJSON, err := h.partyService.CreateParty(url, capacity)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(convertion.StringToBytes(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(partyJSON)
	}
}
