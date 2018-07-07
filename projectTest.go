package main

import "fmt"

// enumStrings 枚举字符串
var enumStrings string = "3456789TJQKA2"
var enumStringsDouble string = "33445566778899TTJJQQKKAA22mmMM"
var enumStringsTriple string = "333444555666777888999TTTJJJQQQKKKAAA222"

var StringToCard map[rune]byte = make(map[rune]byte)

func StringToValues(s string) []byte {
	values := make([]byte, 0)
	for _, s := range s {
		values = append(values, StringToCard[s])
	}
	return values
}

func main() {
	StringToCard['3'] = 0x03
	StringToCard['4'] = 0x04
	StringToCard['5'] = 0x05
	StringToCard['6'] = 0x06
	StringToCard['7'] = 0x07
	StringToCard['8'] = 0x08
	StringToCard['9'] = 0x09
	StringToCard['T'] = 0x0A
	StringToCard['J'] = 0x0B
	StringToCard['Q'] = 0x0C
	StringToCard['K'] = 0x0D
	StringToCard['A'] = 0x0E
	StringToCard['2'] = 0x02
	fmt.Printf("%#v\n", StringToCard)

	values := StringToValues(enumStrings)
	fmt.Printf("%#v\n", values)


}
