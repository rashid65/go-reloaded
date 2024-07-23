package goreload

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func LOWAlone(s string) string {
    re := regexp.MustCompile(`(\w+\'?\w?)\(low\)`)
    newString := re.ReplaceAllStringFunc(s, func(match string) string {
        submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
            return match
        }
        word := submatches[1]
        word = strings.ToLower(word)
		return word
    })
	if !strings.EqualFold(newString,s) {
        return LOWAlone(newString)
    }
    return newString
}

func LOW(input string)string{
    //solves all the cases of (low) 
    input = LOWAlone(input)
    words := strings.Fields(input)
    for i:=0; i < len(words); i++{
        if words[i] == "(low," && i+1 < len(words){
            if strings.HasSuffix(words[i+1], ")"){
                number,err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
                if err != nil { //checks the number if its right
                    fmt.Println("error: please provide a correct number", number)
                    return strings.Join(words," ")
                }
                // check number if it was accesable or not then do the required coding\
                if number < 1 {
                    fmt.Println("error: you can't enter a number less than 1")
                    return strings.Join(words, " ")
                }
                if number > i {
                    fmt.Println("error: number is larger than the words insereted")
                    return strings.Join(words," ")
                }
                for k := i-number ; k < i ; k ++ {
                    if k >= 0 && k < len(words) {
                        words[k] = strings.ToLower(words[k])
                    }
                } 
                words = append(words[:i], words[i+2:]...)
			    i-- 
            }
        }
    }
    return strings.Join(words, " ")
}