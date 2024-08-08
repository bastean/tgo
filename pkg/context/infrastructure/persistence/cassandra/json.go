package cassandra

import (
	"bytes"
	"fmt"
	"regexp"
)

func AddExtraKeyQuotes(json []byte) []byte {
	key := regexp.MustCompile(`"[^"]+":`)
	name := regexp.MustCompile(`[^:]+`)

	keys := key.FindAll(json, -1)
	var quoted []byte

	for _, key := range keys {
		quoted = []byte(fmt.Sprintf("%q:", name.Find(key)))
		json = bytes.Replace(json, key, quoted, 1)
	}

	return json
}

func RemoveExtraKeyQuotes(json string) []byte {
	quote := regexp.MustCompile(`\\"{1}`)
	return []byte(quote.ReplaceAllString(json, ""))
}
