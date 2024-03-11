package party

import (
	"errors"
	"kimparty/internal/cmap"
	"kimparty/internal/party/user"
	"log"
	"strconv"
)

var (
	ErrInvalidURL      = errors.New("invalid url")
	ErrPartyNotFound   = errors.New("party not found")
	ErrInvalidCapacity = errors.New("invalid capacity")
)

type Service struct {
	parties cmap.ConcurrentMap[*Party]
}

func NewService() *Service {
	return &Service{
		parties: cmap.New[*Party](),
	}
}

func (s *Service) CreateParty(url string, capacity string) ([]byte, error) {
	if url == "" {
		return nil, ErrInvalidURL
	}

	cap, err := handlePartyCapacity(capacity)

	if err != nil {
		return nil, err
	}

	pt := New(url, cap)
	s.parties.Set(pt.ID, pt)

	ptJSON, err := pt.ToJSON()

	if err != nil {
		return nil, err
	}

	log.Printf("Party [%s] created", pt.ID)
	return ptJSON, nil
}

func (s *Service) FindParty(id string) (*Party, error) {
	if id == "" {
		return nil, ErrPartyNotFound
	}

	pt, ok := s.parties.Get(id)

	if !ok {
		return nil, ErrPartyNotFound
	}

	return pt, nil
}

func (s *Service) PrepareForEntry(partyID string, username string) (*Party, *user.User, error) {
	if partyID == "" {
		return nil, nil, ErrPartyNotFound
	}

	pt, err := s.FindParty(partyID)

	if err != nil {
		return nil, nil, err
	}

	newUser := user.New(username, pt.ID)
	return pt, newUser, nil
}

func (s *Service) RemovePartyIfEmpty(party *Party) {
	if party.CountConns() == 0 {
		s.parties.Remove(party.ID)
		log.Printf("Party [%s] removed", party.ID)
	}
}

func handlePartyCapacity(capacity string) (uint8, error) {
	var (
		parsedCap uint64
		err       error
	)

	var cap uint8 = 2

	if capacity == "" {
		return cap, nil
	}

	parsedCap, err = strconv.ParseUint(capacity, 10, 8)
	cap = uint8(parsedCap)

	if err != nil || cap < 2 || cap > MaxPartyCapacity {
		return 0, ErrInvalidCapacity
	}

	return cap, nil
}
