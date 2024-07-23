package goreload

import "strings"

func FixMyBack(input string)string{
	chr := []rune(input)
	newString := ""
	for i := 0; i < len(chr); i++ {
		if chr[i]== '(' {
			if (i-1) >= 0 && chr[i-1] != ' '{
				newString += " ("
			}
		} else if chr[i] == ')' {
			if (i+1) != len(chr) && chr[i+1] != ' '{
				newString += ") "
			}
		} else {
			newString += string(chr[i])
		}
	}
	return newString
}

func FixMyGuts (input string)string{
	nochange := input

	input = strings.ReplaceAll(input, " )", ")")
    input = strings.ReplaceAll(input, "( ", "(")

	if nochange == input {
		return input
	}
	return FixMyGuts(input)
}

func FixComma (input string)string {
	output := ""
    for i := 0; i < len(input); i++ {
		if i+1 < len(input){
			if input[i] == ',' && input[i+1] != ' ' {
				output += ", "
			} else {
				output += string(input[i])
			}
		} else {
			output += string(input[i])
		}
    }
    return output
}