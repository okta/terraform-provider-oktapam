package pam

import "encoding/json"

type EncryptedString struct {
	Payload string
}

var _ json.Marshaler = new(EncryptedString)

func (e EncryptedString) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Payload)
}
