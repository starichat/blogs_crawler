package util

import (
	"log"
	"net/http"
)

const userAgent  []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.87 Safari/537.36"
}

func WithRequest(url string) (*http.Response,error) {
	req, err := http.NewRequest("GET",url,nil)
	if err != nil{
		return nil,err
	}
	req.Header.Set("User-Agent", userAgent[0])
	log.Print("Success Set Header")
	return http.DefaultClient.Do(req)
}