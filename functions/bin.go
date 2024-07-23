package goreload

import(
	"strconv"
	"regexp"
	"strings"
)

func Frobin(BINNUM string) (int, error){
	dec, err := strconv.ParseInt(BINNUM, 2, 64)
	if err != nil {
        return 0, err
    }
    return int(dec), nil
}

func Bin (s string) string{
	re := regexp.MustCompile(`([1,0]+)\s?\(bin\)`)
	newString := re.ReplaceAllStringFunc(s, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}
		dec, err := Frobin(submatches[1])
		if err != nil {
			return match
		}
		return strconv.Itoa(dec)
	})
	if !strings.EqualFold(newString,s) {
        return Bin(newString)
    }
	return newString
}