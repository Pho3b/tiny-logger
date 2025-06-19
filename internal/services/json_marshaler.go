package services

import (
	"bytes"
	"fmt"
	"strconv"
)

type jsonLogEntry struct {
	Level    string `json:"level,omitempty"`
	Date     string `json:"date,omitempty"`
	Time     string `json:"time,omitempty"`
	DateTime string `json:"datetime,omitempty"`
	Message  string `json:"msg"`
	Extras   []any  `json:"extras,omitempty"`
}

type JsonMarshaler struct {
}

func (j *JsonMarshaler) Marshal(entry *jsonLogEntry) ([]byte, error) {
	var res bytes.Buffer
	res.Grow(250)

	res.WriteByte('{')

	if entry.Level != "" {
		res.WriteString("\"level\":\"")
		res.WriteString(entry.Level)
		res.WriteByte('"')
		res.WriteByte(',')
	}

	if entry.Date != "" {
		res.WriteString("\"date\":\"")
		res.WriteString(entry.Date)
		res.WriteByte('"')
		res.WriteByte(',')
	}

	if entry.Time != "" {
		res.WriteString("\"time\":\"")
		res.WriteString(entry.Time)
		res.WriteByte('"')
		res.WriteByte(',')
	}

	if entry.DateTime != "" {
		res.WriteString("\"datetime\":\"")
		res.WriteString(entry.DateTime)
		res.WriteByte('"')
		res.WriteByte(',')
	}

	res.WriteString("\"msg\":\"")
	res.WriteString(entry.Message)
	res.WriteByte('"')

	if entry.Extras != nil {
		res.WriteString(",\"extras\":{")

		for i, extra := range entry.Extras {
			switch v := extra.(type) {
			case string:
				res.WriteString(v)
			case rune:
				res.WriteRune(v)
			case int:
				res.WriteString(strconv.Itoa(v))
			case int64:
				res.WriteString(strconv.FormatInt(v, 10))
			case float64:
				res.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
			case bool:
				res.WriteString(strconv.FormatBool(v))
			default:
				// Using the slower fmt.Sprint only for unknown types
				res.WriteString(fmt.Sprint(v))
			}

			if i%2 == 0 && i > 0 {
				res.WriteString("\":")
			} else {
				res.WriteByte(',')
			}
		}

		if len(entry.Extras)%2 != 0 {
			res.WriteString("null")
		}

		res.WriteByte('}')
	}

	res.WriteByte('}')
	return res.Bytes(), nil
}

//// buildExtraMessages constructs a map from a variadic list of key-value pairs.
//// It expects an even number of arguments, where even indices (0, 2, 4, ...) are keys
//// and odd indices (1, 3, 5, ...) are values. If an odd number of arguments is passed,
//// the last key will be assigned a `nil` value.
////
//// Example Usage:
////
////	extra := b.buildExtraMessages("user", "alice", "ip", "192.168.1.1")
////	// Result: map[string]interface{}{"user": "alice", "ip": "192.168.1.1"}
//func (j *JsonMarshaler) buildExtraMessages(keyAndValuePairs ...interface{}) map[string]interface{} {
//	keyAndValuePairsLen := len(keyAndValuePairs)
//	if keyAndValuePairsLen == 0 {
//		return nil
//	}
//
//	resMap := make(map[string]interface{}, keyAndValuePairsLen/2)
//
//	for i := 0; i < keyAndValuePairsLen-1; i += 2 {
//		key := fmt.Sprint(keyAndValuePairs[i])
//		value := keyAndValuePairs[i+1]
//		resMap[key] = value
//	}
//
//	if keyAndValuePairsLen%2 != 0 {
//		lastKey := fmt.Sprint(keyAndValuePairs[keyAndValuePairsLen-1])
//		resMap[lastKey] = nil
//	}
//
//	return resMap
//}
