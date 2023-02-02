package crypto

import (
	"encoding/base64"
	"encoding/json"
)

type JsonMsgSerializer struct {
}

type JsonMsgMetadata struct {
	Rails struct {
		Message string          `json:"message"`
		Exp     json.RawMessage `json:"exp"`
		Pur     json.RawMessage `json:"pur"`
	} `json:"_rails"`
}

func (s JsonMsgSerializer) Serialize(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (s JsonMsgSerializer) Unserialize(data string, v interface{}) error {
	bs := []byte(data)
	meta := JsonMsgMetadata{}
	err := json.Unmarshal(bs, &meta)
	if err == nil && len(meta.Rails.Message) > 0 {
		if bs, err = base64.StdEncoding.DecodeString(meta.Rails.Message); err != nil {
			return err
		}
	}

	return json.Unmarshal(bs, v)
}
