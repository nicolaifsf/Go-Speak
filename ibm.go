package speech

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var ibmUsername string
var ibmPassword string

/**
* Set your IBM Credentials in order to successfully communicate with api
* @param username is your ibm api username
* @param password is your ibm api password
**/
func SetIBMCredentials(username string, password string) {
	ibmUsername = username
	ibmPassword = password
}

/**
*Sends an audio file to ibm, ibm username and password must have been set prior to calling
* otherwise error is thrown
*@param filename the full path to the file that is to be sent
*@return a string with the json data received
**/
func SendIBMVoice(fileRef string) string {
	//currentDir, err := os.Getwd()
	//fileRef := currentDir + "/test.wav"
	/*
		if err != nil {
			log.Fatal(err)
		}
	*/
	audio, err := ioutil.ReadFile(fileRef)

	if err != nil {
		log.Fatal("Error reading file:\n%v\n", err)

	}

	reader := bytes.NewReader(audio)

	url := "https://api.wit.ai/speech?v=20141022"
	client := &http.Client{}
	/*
		f, err := os.Open(fileRef)
		if err != nil {
			log.Fatal(err)
			fmt.Println(f)
		}*/
	//req, err := http.NewRequest("POST", url, f)
	req, err := http.NewRequest("POST", url, reader)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+witKey)
	req.Header.Set("Content-Type", "audio/wav")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(body))
	return string(body)
}

type IbmSeshJson struct {
	Recognize       string `json: "recognize"`
	RecognizeWS     string `json: "recognizeWS"`
	Observe_result  string `json: "observe_result"`
	Session_id      string `json: "session_id"`
	New_session_uri string `json: "new_session_uri"`
}

/**
* Gets a session url and sets up an IBMSeshJson
*
**/
func GetSession() IbmSeshJson {
	url := "stream.watsonplatform.net/speech-to-text/api/v1/sessions"
	seshUrl := fmt.Sprintf("https://%s:%s@%s", ibmUsername, ibmPassword, url)
	req, err := http.NewRequest("POST", seshUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bod, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	inString := string(bod)
	data := []byte(inString)
	var sesh IbmSeshJson
	err = json.Unmarshal(data, &sesh)
	if err != nil {
		fmt.Println("this err")
		log.Fatal(err)
	}
	return sesh
}
