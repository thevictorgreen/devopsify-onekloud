package main
import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "bytes"
)


func es(content []string) string {

  //Write File
  cmdLines := string(30)
	f, err := os.Create(cmdLines)
  if err != nil {
		fmt.Println(err)
		f.Close()
		log.Fatalf("%s\n", err)
  }
	for _, v := range content {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			log.Fatalf("%s\n", err)
		}
  }
	err = f.Close()
  if err != nil {
		fmt.Println(err)
		log.Fatalf("%s\n", err)
  }
	//Execute File
	cmd := exec.Command("/bin/bash", cmdLines)
  var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout//os.Stdout
	cmd.Stderr = &stderr//os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	//Delete File
	err = os.Remove(cmdLines)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
  //Return Data
  outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
  if errStr != "" {
    return errStr
  } else {
    return outStr
  }
}
