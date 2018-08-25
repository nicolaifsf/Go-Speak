package speech

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

var witKey string

// SetWitKey witKey must be set prior to executing any wit commands
func SetWitKey(key string) string {
	witKey = key
	return witKey
}

// PrintWitKey Returns the current wit key if set, otherwise returns nil
func PrintWitKey() string {
	return witKey
}

// convert converts a message with spaces into one suitable to passing to wit
func convert(message string) string {
	arrString := strings.Split(message, " ")
	var ret string
	for x := 0; x < len(arrString); x++ {
		ret += arrString[x] + "%20"
	}

	return ret
}

// SendWitMessage sends the given message to wit api
func SendWitMessage(message string) (string, error) {
	url := "https://api.wit.ai/message?v=20160225&q=" + convert(message)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+witKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

// SendWitVoice Sends an audio file to wit.ai, wit key must have been set prior to calling
//  - @param filename the full path to the file that is to be sent
//  - @return a string with the json data received
func SendWitVoice(fileRef string) (string, error) {
	audio, err := ioutil.ReadFile(fileRef)
	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(audio)

	url := "https://api.wit.ai/speech?v=20141022"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+witKey)
	req.Header.Set("Content-Type", "audio/wav")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// SendWitBuff sends an audio buffer to wit.ai
func SendWitBuff(buffer *bytes.Buffer) (string, error) {
	url := "https://api.wit.ai/speech?v=20141022"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+witKey)
	req.Header.Set("Content-Type", "audio/wav")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
