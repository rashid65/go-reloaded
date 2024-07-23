package main

import (
	"fmt"
	goreload "goreload/functions"
)

func main() {
	fmt.Println("---------------------------------hex test-------------------------------")
	hex1 := "14(hex) stitches, fun turning -0x15 (hex) , error handeling 2Y (hex) "
	fmt.Println("text: ",hex1)
	modifiedText := goreload.Hex(hex1)
	fmt.Println(modifiedText)
	fmt.Println("---------------------------------bin test-------------------------------")
	bin1 := "10110 (bin) too young to be messed with, im only -11001(bin) idk anything "
	fmt.Println("text: ",bin1)
	modifiedText = goreload.Bin(bin1)
	fmt.Println(modifiedText)
	fmt.Println("---------------------------------up test-----------------------------")
	up2 := "say say(up), dont don't (up, 2) , go go (up, 2)"
	fmt.Println("text: ",up2)
	modifiedText = goreload.UP(up2)
	fmt.Println(modifiedText)
	fmt.Println("---------------------------------low test-----------------------------")
	low1 := "I SAID(low) I LOVE YOU (low,  2)"
	fmt.Println("text= ", low1)
	modifiedText = goreload.LOW(low1)
	fmt.Println(modifiedText)
	fmt.Println("---------------------------------cap test-----------------------------")
	cap1 := "hello (cap) how are you (cap, 3)"
	fmt.Println("text= ",cap1)
	modifiedText = goreload.CAP(cap1)
	fmt.Println(modifiedText)
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(goreload.FixMyBack("am(low).stupid hahah(will)never changenayef(nawaf)7amar"))
	fmt.Println(goreload.FixMyGuts("am( low) stupid hahah (  will ) never changenayef (  nawaf   ) 7amar"))
	FixMe := "HELLO (low) am(low,4).stupid hahah(low, 4)never change, nayef(low,  4)7amar"
	Fixed := (goreload.FixComma(FixMe))
	fmt.Println(Fixed)
}
