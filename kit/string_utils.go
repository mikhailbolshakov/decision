package kit

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"encoding/hex"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var (
	encoding     = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")
	digitsRegExp = regexp.MustCompile(`^\d+$`)
)

// NewRandString generates a unique string
func NewRandString() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	_, _ = encoder.Write(uuid.NewV4().Bytes())
	_ = encoder.Close()
	b.Truncate(26)
	return b.String()
}

// NewId generates unique Id
// use this function for Id generation
func NewId() string {
	return uuid.NewV4().String()
}

// UUID generates UUID
func UUID(size int) string {
	u := make([]byte, size)
	_, _ = io.ReadFull(rand.Reader, u)
	return hex.EncodeToString(u)
}

// ValidateUUIDs check UUID format nad return error if at least one UUID isn't in a correct format
func ValidateUUIDs(uuids ...string) error {
	for _, u := range uuids {
		if _, err := uuid.FromString(u); err != nil {
			return err
		}
	}
	return nil
}

// Nil returns nil UUID
func Nil() string {
	return uuid.Nil.String()
}

// TODO: remove
func ToJson(v interface{}) (string, error) {
	if v != nil {
		var b, err = json.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return "", nil
}

func MustJson(v interface{}) string {
	s, _ := ToJson(v)
	return s
}

func Json(i interface{}) string {
	r, _ := json.Marshal(i)
	return string(r)
}

// Strings represents slice of strings
type Strings []string

// Distinct returns distinct slice
func (s Strings) Distinct() Strings {
	var res []string
	m := make(map[string]struct{}, len(s))
	for _, i := range s {
		if _, ok := m[i]; !ok {
			m[i] = struct{}{}
			res = append(res, i)
		}
	}
	return res
}

// Contains check if a strings is in slice
func (s Strings) Contains(str string) bool {
	for _, i := range s {
		if i == str {
			return true
		}
	}
	return false
}

func (s Strings) Intersect(r Strings) Strings {
	res := Strings{}
	rDistinct := r.Distinct()
	for _, i := range s.Distinct() {
		for _, j := range rDistinct {
			if i == j {
				res = append(res, i)
			}
		}
	}
	return res
}

func StrToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func RemoveNonAlfaDigital(str string) string {
	reg := regexp.MustCompile(`[^0-9a-zA-ZА-Яа-я]|\^|\_`)
	return reg.ReplaceAllString(str, "")
}

func Digits(s string) bool {
	if s == "" {
		return false
	}
	return digitsRegExp.MatchString(s)
}

func YesNoToBool(s string) *bool {
	t, f := true, false
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "нет", "no":
		return &f
	case "да", "yes":
		return &t
	}
	return nil
}
