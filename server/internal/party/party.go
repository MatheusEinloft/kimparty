package party

import (
	"errors"
	"time"

	"kimparty/config"
	"kimparty/internal/cmap"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/lxzan/gws"
)

var (
	MaxPartyCapacity = config.GetPartyMaxCapacity()
	ErrPartyFull     = errors.New("party is full")
)

type Party struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	Capacity    uint8    `json:"capacity"`
	connections cmap.ConcurrentMap[*gws.Conn]
}

func New(url string, capacity uint8) *Party {
	if capacity < 2 || capacity > MaxPartyCapacity {
		panic("invalid party capacity")
	}

	return &Party{
		ID:          uuid.New().String(),
		URL:         url,
		Capacity:    capacity,
		connections: cmap.New[*gws.Conn](),
		CreatedAt:   time.Now().Format(time.RFC3339),
	}
}

func (pt *Party) AddConn(k string, conn *gws.Conn) error {
	if pt.CountConns() >= int(pt.Capacity) {
		return ErrPartyFull
	}

	pt.connections.Set(k, conn)
	return nil
}

func (pt *Party) RemoveConn(k string) {
	pt.connections.Remove(k)
}

func (pt *Party) CountConns() int {
	return pt.connections.Count()
}

func (pt *Party) IterConns() <-chan *gws.Conn {
	return pt.connections.Iter()
}

func (pt *Party) ToJSON() ([]byte, error) {
	return jsoniter.ConfigFastest.Marshal(pt)
}
