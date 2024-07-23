package goreload

import(
	"strconv"
	"regexp"
	"strings"
)

func Frohex(hexnum string) (int, error){
	dec, err := strconv.ParseInt(hexnum, 16, 64)
	if err != nil {
        return 0, err
    }
    return int(dec), nil
}

func Hex (s string) string{
	re := regexp.MustCompile(`(?:0[xX])?([0-9a-fA-F]+)\s?\(hex\)`)
	newString := re.ReplaceAllStringFunc(s, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}
		dec, err := Frohex(submatches[1])
		if err != nil {
			return match
		}
		return strconv.Itoa(dec)
	})
	if !strings.EqualFold(newString,s) {
        return Hex(newString)
    }
	return newString
}
