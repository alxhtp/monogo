package jsonconvert

import (
	"encoding/json"
	"log"
)

// Serialize takes an interface and print json as string.
func Serialize(any interface{}) string {
	str, err := json.Marshal(any)
	if err != nil {
		log.Printf("Error marshalling %s", err)
	}

	return string(str)
}

// Deserialize takes an interface and convert it to the output wanted
func Deserialize(any interface{}, output interface{}) {
	bytes, _ := json.Marshal(any)
	err := json.Unmarshal(bytes, output)
	if err != nil {
		log.Printf("error unmarshalling attributes: %v", err)
	}
}
