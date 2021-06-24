package surl_engine

import (
	"fmt"
	"github.com/waspnesser/sURL/base64"
	"time"
	"math"
	"math/rand"
)

var url_map = map[base64.Base64]string{}
var Limit int64 = int64(math.Pow(2,48))

func get_random()base64.Base64{
	rand.Seed(time.Now().UnixNano())
	return base64.ToBase64(uint64(rand.Int63n(Limit)))
}

func get_code()base64.Base64{
	random_code := get_random()
	_,ok := url_map[random_code]

	for ;ok;{
		random_code = get_random()
		_,ok = url_map[random_code]
	}

	return random_code
}

func AddURL(url string) base64.Base64{
	code := get_code()
	url_map[code] = url
	return code
}

func GetURL(surl string) (string,bool){
	value,ok := url_map[base64.Base64{surl}]
	return value,ok
}

func main(){
	msg := "\n100 Shorten URL\n101 Get URL\n200 Show all shorten URLS\n500 Exit\nOperation code : "
	var response int
	var end_flag bool
	for {
		fmt.Print(msg)
		fmt.Scanln(&response)
		switch response{
			case 500:
				end_flag = true
			case 100:
				var url string
				fmt.Print("URL : ")
				fmt.Scanln(&url)
				fmt.Println("URL has been shortened to ",AddURL(url))
			case 101:
				var surl string
				fmt.Print("Shortened URL : ")
				fmt.Scanln(&surl)
				value,ok := GetURL(surl)
				if ok{
					fmt.Println("URL of ",surl," is ",value)
				} else{
					fmt.Println("Invalid shortened URL")
				}

			case 200:
				fmt.Println(url_map)
			default:
				fmt.Println("Unknown operation code.")

		}
		if end_flag{
			break
		}


	}

}
