package steno

import (
	"encoding/json"
	"fmt"
)

type JsonCodec struct {
}

func NewJsonCodec() Codec {
	return new(JsonCodec)
}

func (j *JsonCodec) EncodeRecord(record *Record) string {
	hash := map[string]string {
		"timestamp" : record.timestamp.String(),
		"message"   : record.message,
		"log_level" : record.level.name,
	}

	bytes, err := json.Marshal(hash)
	message := string(bytes)
	if err != nil {
		message = fmt.Sprintf("Error: Encoding JsonCodec, record: (%s)", err)
	}
	return message + "\n"
}