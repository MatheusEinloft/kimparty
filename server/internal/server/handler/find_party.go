package handler

import (
	"kimparty/internal/party"
	"kimparty/internal/server"
	"kimparty/pkg/convertion"
	"net/http"
)

type FindPartyHandler struct {
	partyService *party.Service
}

func NewFindPartyHandler(partyService *party.Service) server.Handler {
	return &FindPartyHandler{
		partyService: partyService,
	}
}

func (h *FindPartyHandler) Path() string {
	return "/party/find"
}

func (h *FindPartyHandler) Methods() string {
	return "GET"
}

func (h *FindPartyHandler) Handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(convertion.StringToBytes("id is required"))
			return
		}

		pt, err := h.partyService.FindParty(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(convertion.StringToBytes(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)

		ptJSON, err := pt.ToJSON()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(convertion.StringToBytes(err.Error()))
			return
		}

		w.Write(ptJSON)
	}
}
