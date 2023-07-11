package kit

import (
	"regexp"
	"strings"
)

// IsEmailValid checks email format
func IsEmailValid(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}
	match, err := regexp.MatchString("^(((\\\\.)|[^\\s[:cntrl:]\\(\\)<>@,;:'\\\\\\\"\\.\\[\\]]|')+|(\"(\\\\\"|[^\"])*\"))(\\.(((\\\\.)|[^\\s[:cntrl:]\\(\\)<>@,;:'\\\\\\\"\\.\\[\\]]|')+|(\"(\\\\\"|[^\"])*\")))*@[a-zA-Z0-9а-яА-Я](?:[a-zA-Z0-9а-яА-Я-]{0,61}[a-zA-Z0-9а-яА-Я])?(?:\\.[a-zA-Z0-9а-яА-Я](?:[a-zA-Z0-9а-яА-Я-]{0,61}[a-zA-Z0-9а-яА-Я])?)*$", email)
	return match && err == nil
}

func IsUrlValid(url string) bool {
	match, err := regexp.MatchString("^(https:\\/\\/)?(http:\\/\\/)?(www\\.)?(([-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6})|(localhost:[0-9]{1,4}))\\b([-a-zA-Z0-9()@:%_\\+.~#?&\\/=]*)$", url)
	return match && err == nil
}

// IsIpV4Valid checks ip v4 format
func IsIpV4Valid(ip string) bool {
	match, err := regexp.MatchString("^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$", ip)
	return match && err == nil
}

// IsIpV6Valid checks ip v6 format
func IsIpV6Valid(ip string) bool {
	match, err := regexp.MatchString("^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$", ip)
	return match && err == nil
}

// IsPhoneValid checks phone format (with country code without special characters)
func IsPhoneValid(phone string) bool {
	match, err := regexp.MatchString("^\\d{2,14}$", phone)
	return match && err == nil
}

// IsRussianPhoneValid checks Russian phone format (with country code without special characters)
func IsRussianPhoneValid(phone string) bool {
	match, err := regexp.MatchString("^(7|8)\\d{10}$", phone)
	return match && err == nil
}

var cyrLatinMap = map[rune]rune{
	'\u0423': '\u0059',
	'\u0425': '\u0058',
	'\u0415': '\u0045',
	'\u041A': '\u004B',
	'\u041C': '\u004D',
	'\u0421': '\u0043',
	'\u0422': '\u0054',
	'\u0410': '\u0041',
	'\u0412': '\u0042',
	'\u041D': '\u0048',
	'\u041E': '\u004F',
	'\u0420': '\u0050',
}

// CarRegNumberSanitize sanitizes a car registration number
func CarRegNumberSanitize(regNum string) string {
	r := RemoveNonAlfaDigital(strings.ToUpper(regNum))
	runes := []rune(r)
	for i, r := range runes {
		if latin, ok := cyrLatinMap[r]; ok {
			runes[i] = latin
		}
	}
	return string(runes)
}

// IsCarRegNumberValid checks if passed a valid car registration number
func IsCarRegNumberValid(regNumber string) bool {
	if regNumber == "" {
		return false
	}
	if ok, _ := regexp.MatchString("^[A-Z0-9]{3,16}$", regNumber); !ok {
		return false
	}
	return true
}

// IsVINValid checks if passed valid VIN number
func IsVINValid(vin string) bool {
	if vin == "" {
		return false
	}
	if ok, _ := regexp.MatchString("^[(A-H|J-N|P|R-Z|0-9)]{17}$", vin); !ok {
		return false
	}
	return true
}

// IsCarCertificateNumberValid checks if car certificate number (STS) valid
func IsCarCertificateNumberValid(certNum string) bool {
	if certNum == "" {
		return false
	}
	if ok, _ := regexp.MatchString("^[0-9]{2}[А-ЯA-Z0-9]{2}(\\s)+[0-9]{6}$", certNum); !ok {
		return false
	}
	return true
}

// IsCarEDocPassportNumberValid checks if passed car edoc passport number (electronic PTS) valid
func IsCarEDocPassportNumberValid(num string) bool {
	if num == "" {
		return false
	}
	if ok, _ := regexp.MatchString("^[0-9]{15}$", num); !ok {
		return false
	}
	return true
}

// IsCarPaperPassportNumberValid checks if passed car paper passport number (PTS) valid
func IsCarPaperPassportNumberValid(num string) bool {
	if num == "" {
		return false
	}
	if ok, _ := regexp.MatchString("^[0-9]{2}[АВЕКМНОРСТУХЛ0]{2}(\\s)+[0-9]{6}$", num); !ok {
		return false
	}
	return true
}

func IsTelegramChannelValid(channel string) bool {
	if channel == "" {
		return false
	}
	ok, _ := regexp.MatchString("^(https?:\\/\\/)?(www[.])?(telegram|t)\\.me\\/([a-zA-Z0-9_-]*)\\/?$", channel)
	return ok
}

// IsCoordinateValid checks if coordinate valid
func IsCoordinateValid(c string) bool {
	ok, _ := regexp.MatchString(`^-?[0-9]{1,2}\.[0-9]{5,7}$`, c)
	return ok
}
