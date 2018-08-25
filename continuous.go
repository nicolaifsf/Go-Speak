package speech

import (
	"bytes"
	"fmt"
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
func ContinuousRecognition() error {
	var err error
	for {
		err = start()
		if err != nil {
			return err
		}
	}
}

func start() error {
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
	text, err := SendWitBuff(buf)
	if err != nil {
		return err
	}

	fmt.Println(text)
	return nil
}
