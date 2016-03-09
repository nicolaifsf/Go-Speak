package speech

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	key string = "AIzaSyBOti4mM-6x9WDnZIjIeyEU21OpBXqWBgw"
	url string = "https://www.google.com/speech-api/v2/recognize?output=json&lang=en-us&key=" + key
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

func GetGoogleCreds() string {
	return googleKey
}
func SetGoogleKey(key string) {
	googleKey = key
}

func SendGoogleMessage(file string) string {
	//stream, err := ioutil.ReadFile(file)
	stream, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	info, err := stream.Stat()
	if err != nil {
		panic(err)
	}
	var size int64 = info.Size()
	b := make([]byte, size)
	buffer := bufio.NewReader(stream)
	_, err = buffer.Read(b)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
	//	req.Header.Set("Content-Type", "audio/l16; rate=44100;")
	req.Header.Set("Content-Type", "audio/l16; rate=16000")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bod, err := ioutil.ReadAll(res.Body)
	return string(bod)

}
