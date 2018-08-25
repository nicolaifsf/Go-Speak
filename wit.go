package speech

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

// WitHandler handles all the wit.ai APIs
type WitHandler struct {
	witKey string
}

// NewWitHandler returns a new WitHandler instance
func NewWitHandler(key string) *WitHandler {
	return &WitHandler{key}
}

// GetWitKey Returns the current wit key if set, otherwise returns nil
func (wh *WitHandler) GetWitKey() string {
	return wh.witKey
}

// SendWitMessage sends the given message to wit api
func (wh *WitHandler) SendWitMessage(message string) (string, error) {
	url := "https://api.wit.ai/message?v=20160225&q=" + convert(message)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+wh.witKey)
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
func (wh *WitHandler) SendWitVoice(fileRef string) (string, error) {
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

	req.Header.Set("Authorization", "Bearer "+wh.witKey)
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
func (wh *WitHandler) SendWitBuff(buffer *bytes.Buffer) (string, error) {
	url := "https://api.wit.ai/speech?v=20141022"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+wh.witKey)
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

// ContinuousRecognition starts recording audio and sends the buffer to Wit.ai
func (wh *WitHandler) ContinuousRecognition() error {
	var err error
	for {
		err = wh.start()
		if err != nil {
			return err
		}
	}
}

func (wh *WitHandler) start() error {
	cmd2 := "rec"
	arg2 := []string{
		"-t", "wav", "-",
		"rate", "32k", "silence",
		"1", "0.1", "2%",
		"1", "3.0", "0.25%",
	}

	var byteArr []byte
	buf := bytes.NewBuffer(byteArr)

	cmdExec := exec.Command(cmd2, arg2...)
	stdout, err := cmdExec.StdoutPipe()
	if err != nil {
		return err
	}

	err = cmdExec.Start()
	if err != nil {
		return err
	}

	buf.ReadFrom(stdout)
	text, err := wh.SendWitBuff(buf)
	if err != nil {
		return err
	}

	fmt.Println(text)
	return nil
}

func convArgs(strArray []string) string {
	res := ""
	for x := 0; x < len(strArray); x++ {
		res += strArray[x]
	}
	return res
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
