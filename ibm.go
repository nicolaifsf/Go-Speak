package speech

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
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
	Cookies         *cookiejar.Jar
}

/**
* Gets a session url and sets up an IBMSeshJson
*
**/
func GetSession() IbmSeshJson {
	url := "stream.watsonplatform.net/speech-to-text/api/v1/sessions"
	seshUrl := fmt.Sprintf("https://%s:%s@%s", ibmUsername, ibmPassword, url)
	jsonStr := []byte(`{}`)
	req, err := http.NewRequest("POST", seshUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	cookiesJar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Jar: cookiesJar,
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 201 {
		fmt.Println("WELL SHIT")
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
	sesh.Cookies = cookiesJar
	return sesh
}

func GetResults(sesh IbmSeshJson) {
	results := sesh.Observe_result
	client := &http.Client{
		Jar: sesh.Cookies,
	}
	req, err := http.NewRequest("GET", results, nil)

	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bod, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bod))
}

func SendIBMSpeech(fileRef string) {
	sesh := GetSession()
	seshURL := sesh.Recognize
	client := &http.Client{
		Jar: sesh.Cookies,
	}
	fmt.Println("seshURL = " + seshURL)
	reader, err := os.Open(fileRef)
	if err != nil {
		log.Fatal("Error reading file:\n%v\n", err)
	}

	toURL := fmt.Sprintf("%s", seshURL)
	fmt.Println("toURL : " + toURL)
	req, err := http.NewRequest("POST", toURL, reader)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s:%s", ibmUsername, ibmPassword))
	req.Header.Set("Content-Type", "audio/wav")
	req.Header.Set("Transfer-Encoding", "Chunked")
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bod, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bod))
}
func ReceiveSpeech(s *IbmSeshJson) {
	seshURL := s.Observe_result
	client := &http.Client{}
	req, err := http.NewRequest("GET", seshURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s:%s", ibmUsername, ibmPassword))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bod, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bod))
}
