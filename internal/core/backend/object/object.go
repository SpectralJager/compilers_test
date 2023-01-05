package object

import (
	"bytes"
	"encoding/gob"
)

type Object struct {
	Type ObjectType
	Data []byte
}

func NewObject(tp ObjectType, data []byte) *Object {
	return &Object{Type: tp, Data: data}
}

func EncodeData[T float64 | bool | string](data T) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func DecodeData[T float64 | bool | string](encodedData []byte) T {
	var buf bytes.Buffer
	var decodedData T
	_, err := buf.Write(encodedData)
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(&decodedData)
	if err != nil {
		panic(err)
	}
	return decodedData
}
