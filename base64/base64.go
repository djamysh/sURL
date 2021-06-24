package base64

import (
	"fmt"
	"math"
)

var Base64_chars string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+="
var Base64_map = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'a': 10,
	'b': 11,
	'c': 12,
	'd': 13,
	'e': 14,
	'f': 15,
	'g': 16,
	'h': 17,
	'i': 18,
	'j': 19,
	'k': 20,
	'l': 21,
	'm': 22,
	'n': 23,
	'o': 24,
	'p': 25,
	'q': 26,
	'r': 27,
	's': 28,
	't': 29,
	'u': 30,
	'v': 31,
	'w': 32,
	'x': 33,
	'y': 34,
	'z': 35,
	'A': 36,
	'B': 37,
	'C': 38,
	'D': 39,
	'E': 40,
	'F': 41,
	'G': 42,
	'H': 43,
	'I': 44,
	'J': 45,
	'K': 46,
	'L': 47,
	'M': 48,
	'N': 49,
	'O': 50,
	'P': 51,
	'Q': 52,
	'R': 53,
	'S': 54,
	'T': 55,
	'U': 56,
	'V': 57,
	'W': 58,
	'X': 59,
	'Y': 60,
	'Z': 61,
	'+': 62,
	'=': 63,
}

type Base64_interface interface{
	toBase10() uint64
}

type Base64 struct{
	Number string
}

func (b Base64) String() string{
	return fmt.Sprintf("%s",b.Number)
}

func (b Base64) ToBase10() uint64{
	length := len(b.Number)
	var result uint64 = 0
	var digit int
	for i:=length-1;i>=0;i--{
		digit  = Base64_map[b.Number[i]]
		result += uint64(digit)*uint64(math.Pow(64,float64(length-1-i)))
	}

	return result
}

func ToBase64(Number uint64) Base64{
	var result string = ""
	//var digit_count int = math.Ceil(math.Log(Number)/math.Log(64))
	for ;Number>0;{
		digit := Number %64
		result =  string(Base64_chars[digit]) + result
		Number = (Number - digit)/64
	}

	return Base64{result}
}
