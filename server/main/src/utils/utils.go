package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var ouiDatabase map[string]string

func init() {
	ouiBytes, _ := ioutil.ReadFile("static/oui.json")
	json.Unmarshal(ouiBytes, &ouiDatabase)
}

func GetVendorFromOUI(s string) (string, error) {
	if strings.Count(s, ":") == 5 {
		ouiHeader := strings.ToUpper(strings.Replace(strings.TrimPrefix(s, "wifi-"), ":", "", -1))
		if v, ok := ouiDatabase[ouiHeader[:6]]; ok {
			return strings.Split(v, "\n")[0], nil
		}
		return "?", nil
	}
	return "", errors.New("not a mac address")
}

// src is seeds the random generator for generating random strings
var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyz012345789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandomString prints a random string
func RandomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// IsMacRandomized takes a mac address like "wifi-60:57:18:3d:b8:14"
// or "60:57:18:3d:b8:14" and pulls the first hex digit "60" and computes
// whether or not it is randomized.
// Randomized = Second-least-significant bit of first hex is 0
// (https://en.wikipedia.org/wiki/MAC_address#Universal_vs._local)
func IsMacRandomized(mac string) bool {
	mac = strings.TrimPrefix(mac, "wifi-")
	hexes := strings.Split(mac, ":")
	if len(hexes) != 6 {
		hexes = strings.Split(mac, "-")
		if len(hexes) != 6 {
			return false
		}
	}
	v, _ := strconv.ParseUint(hexes[0], 16, 8)
	return fmt.Sprintf("%08b", v)[6] == byte(49)
}
