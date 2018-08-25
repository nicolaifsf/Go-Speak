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

// ContinuousRecognition starts recording audio and sends the buffer to Wit.ai
func ContinuousRecognition() {
	for {
		start()
	}
}

func start() {
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
		log.Fatalf("Error getting StdoutPipe: %v", err)
	}

	err = cmdExec.Start()
	if err != nil {
		log.Fatalf("Error executing command %s: %v", cmd2, err)
	}

	buf.ReadFrom(stdout)
	fmt.Println(SendWitBuff(buf))
}
