package schema

import (
	"encoding/json"
)

type RegisterRequest struct {
	Port int
}

type RegisterResponse struct {
	ClientId ClientId
	Code ErrorCode
}

func (registerResponse *RegisterResponse) Marshal() ([]byte, error) {
	return json.Marshal(registerResponse)
}

func (registerResponse *RegisterResponse) Unmarshal(data []byte) error {
	return json.Unmarshal(data, registerResponse)
}

type UnregisterRequest struct {
	ClientId ClientId
}

type UnregisterResponse struct {
	ClientId ClientId
	Code ErrorCode
}