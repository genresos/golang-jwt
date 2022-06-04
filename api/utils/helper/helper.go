package helper

import (
	"regexp"
	"strings"
)

func ExtractValue(body string, key string) string {

	// usage : helper.ExtractValue(string(body), "key")
	
    keystr := "\"" + key + "\":[^,;\\]}]*"
    r, _ := regexp.Compile(keystr)
    match := r.FindString(body)
    keyValMatch := strings.Split(match, ":")
    return strings.ReplaceAll(keyValMatch[1], "\"", "")
}