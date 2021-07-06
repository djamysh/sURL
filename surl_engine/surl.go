package surl_engine

import (
	"fmt"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/waspnesser/sURL/base64"
	"time"
	"math"
	"math/rand"
)


var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
})

var Limit int64 = int64(math.Pow(2,48))

func get_random()base64.Base64{
	rand.Seed(time.Now().UnixNano())
	return base64.ToBase64(uint64(rand.Int63n(Limit)))
}

func AddURL(url string,expiration int) base64.Base64{

	var code base64.Base64

	for xstat:=false;!xstat; {

		code = get_random()
		st,err := rdb.SetNX(ctx, code.String(), url, time.Duration(expiration)*time.Second).Result()
		xstat = st

		if err != nil {
			panic(err)
		}
	}

	return code
}

func GetURL(surl string) (string,bool){
	get := rdb.Get(ctx,surl)
	// if get.Err 
	// == nil	: means everything okay
	// == redis.Nil : specified key not found 
	// else 	: panic!!!

	var err bool

	if get.Err() != nil{
		err = true
	} else {
		err = false
	}
	return get.Val(), err
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
				fmt.Println("URL has been shortened to ",AddURL(url,0))
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
				fmt.Println("Not now.")

				//fmt.Println(url_map)
			default:
				fmt.Println("Unknown operation code.")

		}
		if end_flag{
			break
		}


	}

}
