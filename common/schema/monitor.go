package schema

import (
	"encoding/json"
)

type AllInfo struct {
	ServerAddr string
	Clients []*ClientInfo
}

func (ai *AllInfo) Marshal() ([]byte, error){
	return json.Marshal(ai)
}

func (ai *AllInfo) Unmarshal(data []byte) error {
	return json.Unmarshal(data, ai)
}

type ClientInfo struct {
	ClientId ClientId
	ClientAddr string
	Port int
	ConnectionNumber int
	UploadSpeed int
	DownloadSpeed int
}

func (ci *ClientInfo) Marshal() ([]byte, error){
	return json.Marshal(ci)
}

func (ci *ClientInfo) Unmarshal(data []byte) error {
	return json.Unmarshal(data, ci)
}