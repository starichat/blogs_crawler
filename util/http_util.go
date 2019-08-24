package util

import (
	"log"
	"net/http"
)

var userAgent  =  [...]string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.87 Safari/537.36"}


func GetRequest(url string) (*http.Response,error) {
	req, err := http.NewRequest("GET",url,nil)
	if err != nil{
		return nil,err
	}
	//req.Header.Set("User-Agent", "")
	log.Print("Success Set Header and start request \n")
	client := &http.Client{}
	return client.Do(req)
}