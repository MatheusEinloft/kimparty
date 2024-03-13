package message

func NewLeavedMessage(partyID string, userID string) ([]byte, error) {
	msg := New(Params{
		Content: "leaved",
		PartyID: partyID,
		UserID:  userID,
		Type:    Event,
	})

	return msg.ToJSON()
}

func NewFromUserMessage(partyID string, userID string, msgBytes []byte) ([]byte, error) {
	parsedMsg, err := MessageFromJSON(msgBytes)

	if err != nil {
		return nil, err
	}

	parsedMsg.PartyID = partyID
	parsedMsg.UserID = userID

	return parsedMsg.ToJSON()
}
