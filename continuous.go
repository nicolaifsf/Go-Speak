package speech

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func convArgs(strArray []string) string {
	res := ""
	for x := 0; x < len(strArray); x++ {
		res += strArray[x]
	}
	return res
}

func ContinuousRecognition() {
	for {
		start()
	}
}

func start() {
	/*
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	*/
	//cmd := "sox"
	cmd2 := "rec"
	arg2 := []string{
		//rec test.wav rate 32k silence 1 0.1 3% 1 3.0 3%
		"-t", "wav", "-",
		"rate", "32k",
		/*"rate", "16000", "channels", "1",*/
		"silence", "1", "0.1", "2%", "1", "3.0", "0.25%"}
	/*"silence", "1", "0.1", "0.1%", "1", "1.0", "0.1%"}*/
	/*args := []string{
	"-q",
	"-b", "16",
	"-d", "-t", "flac", "-",
	"rate", "16000", "channels", "1",
	//"silence", "1", "0.1", (ops.threshold || "0.1") + '%', "1", "1.0", (ops.threshold || "0.1") + '%'}
	"silence", "1", "0.1", "0.1" + "%", "1", "1.0", "0.1" + "%"}
	*/
	var byteArr []byte
	buf := bytes.NewBuffer(byteArr)
	cmdExec := exec.Command(cmd2, arg2...)
	stdout, err := cmdExec.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmdExec.Start()
	if err != nil {
		log.Fatal(err)
	}
	buf.ReadFrom(stdout)
	fmt.Println(SendWitBuff(buf))
}
