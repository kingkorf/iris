package ws

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris/utils"
)

type messageType uint8

func (m messageType) String() string {
	return strconv.Itoa(int(m))
}

func (m messageType) Name() string {
	if m == stringMessageType {
		return "string"
	} else if m == intMessageType {
		return "int"
	} else if m == boolMessageType {
		return "bool"
	} else if m == bytesMessageType {
		return "[]byte"
	} else if m == jsonMessageType {
		return "json"
	}

	return "Invalid(" + m.String() + ")"

}

// The same values are exists on client side also
const (
	stringMessageType messageType = iota
	intMessageType
	boolMessageType
	bytesMessageType
	jsonMessageType
)

const (
	prefix          = "iris-websocket-message:"
	separator       = ";"
	prefixLen       = len(prefix)
	separatorLen    = len(separator)
	prefixAndSepIdx = prefixLen + separatorLen - 1
	prefixIdx       = prefixLen - 1
	separatorIdx    = separatorLen - 1
)

var (
	separatorByte = separator[0]
	buf           = utils.NewBufferPool(256)
)

// encodeMessage encodes a custom websocket message from server to be delivered to the client
// returns the  string form of the message
// Supported data are: string, int, bool, bytes and JSON.
func encodeMessage(event string, data interface{}) (string, error) {
	var msgType messageType
	var dataMessage string

	if s, ok := data.(string); ok {
		msgType = stringMessageType
		dataMessage = s
	} else if i, ok := data.(int); ok {
		msgType = intMessageType
		dataMessage = strconv.Itoa(i)
	} else if b, ok := data.(bool); ok {
		msgType = boolMessageType
		dataMessage = strconv.FormatBool(b)
	} else if by, ok := data.([]byte); ok {
		msgType = bytesMessageType
		dataMessage = string(by)
	} else {
		//we suppose is json
		res, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		msgType = jsonMessageType
		dataMessage = string(res)
	}

	b := buf.Get()
	b.WriteString(prefix)
	b.WriteString(event)
	b.WriteString(separator)
	b.WriteString(msgType.String())
	b.WriteString(separator)
	b.WriteString(dataMessage)
	dataMessage = b.String()
	buf.Put(b)

	return dataMessage, nil

}

// decodeMessage decodes a custom websocket message from the client
// ex: iris-websocket-message;user;4;themarshaledstringfromajsonstruct will return themarshaledstringfromajsonstruct as JSON
// Supported data are: string, int, bool, bytes and JSON.
func decodeMessage(event string, websocketMessage string) (message interface{}, err error) {
	t, formaterr := strconv.Atoi(websocketMessage[prefixAndSepIdx+len(event)+1 : prefixAndSepIdx+len(event)+2]) // in order to iris-websocket-message;user;-> 4
	if formaterr != nil {
		return nil, formaterr
	}
	_type := messageType(t)
	_message := websocketMessage[prefixAndSepIdx+len(event)+3:] // in order to iris-websocket-message;user;4; -> themarshaledstringfromajsonstruct

	if _type == stringMessageType {
		message = string(_message)
	} else if _type == intMessageType {
		message, err = strconv.Atoi(_message)
	} else if _type == boolMessageType {
		message, err = strconv.ParseBool(_message)
	} else if _type == bytesMessageType {
		message = []byte(_message)
	} else if _type == jsonMessageType {
		err = json.Unmarshal([]byte(_message), message)
	} else {
		return nil, fmt.Errorf("Type %s is invalid for message: %s", _type.Name(), websocketMessage)
	}

	return
}

// return empty string when the websocketMessage is native message
func getCustomEvent(websocketMessage string) string {
	if len(websocketMessage) < prefixAndSepIdx {
		return ""
	}
	s := websocketMessage[prefixAndSepIdx:]
	evt := s[:strings.IndexByte(s, separatorByte)]
	return evt
}
