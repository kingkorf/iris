package ws

import (
	"encoding/json"
	"strconv"
)

type messageType byte

// The same values must be declared to the client side script also don't forget kataras!
const (
	_                             = iota
	stringMessageType messageType = 1
	intMessageType
	boolMessageType
	bytesMessageType
	jsonMessageType
)

var (
	prefix    = []byte("iris-websocket-message:")
	separator = []byte(";")
)

func appendBytes(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // (if) needed reallocate
		// allocation
		newSlice := make([]byte, (l+len(data))*2)
		// copy
		for i, c := range slice {
			newSlice[i] = c
		}
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

// _msg this will be used to control the flow between the client side iris websockets and server side iris websockets
// at the same time user must be allowed to use iris websockets with native browser's websocket without the iris' client side (type)javascript.
func _msg(event string, messageType messageType, dataMessage []byte) []byte {
	m := prefix
	m = appendBytes(m, []byte(event))
	m = appendBytes(m, separator)
	m = append(m, byte(messageType))
	m = appendBytes(m, separator)
	m = appendBytes(m, dataMessage)
	return m
	// ex: iris-websocket-message;user;json;themarshaledstringfromajsonstruct
	// and client side can parse it to find the event ( first ';' +1 until next ';') or after first separator until second
	// the type, the pre-last ';'+1 until next ';' or after second separator until third
	// the actual message, after the third ';' or after the third separator
}

// Supported data are: string, int, bool, bytes and JSON.
func encodeMessage(event string, data interface{}) (message []byte, err error) {

	if s, ok := data.(string); ok {
		message = _msg(event, stringMessageType, []byte(s))
	} else if i, ok := data.(int); ok {
		message = _msg(event, intMessageType, []byte(strconv.Itoa(i)))
	} else if b, ok := data.(bool); ok {
		message = _msg(event, boolMessageType, []byte(strconv.FormatBool(b)))
	} else if by, ok := data.([]byte); ok {
		message = _msg(event, bytesMessageType, by)
	} else {
		//we suppose is json
		result, err := json.Marshal(data)
		if err != nil {
			return message, err
		}
		message = _msg(event, jsonMessageType, result)
	}
	return message, nil
}
