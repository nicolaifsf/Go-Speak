package speech

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var googleKey string

func SendGoogleRequest(filePath string) {
	//language := "en-US"
	//url := "http://www.google.com/speech-api/v2/recognize?" + language + "&key=" + googleKey
	url := "http://www.google.com/speech-api/v2/recognize?"
	client := &http.Client{}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "audio/wav")
	req.Header.Set("output", "json")
	req.Header.Set("lang", "en-us")
	req.Header.Set("key", googleKey)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

//func GetGoogleCreds() (string, string) {
func GetGoogleCreds() string {
	return googleKey
}
func SetGoogleKey(key string) {
	googleKey = key
}
