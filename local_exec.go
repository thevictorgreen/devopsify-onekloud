package main
import (
  "fmt"
  "log"
  "os"
  "os/exec"
)

func local_exec(content []string) {
  //Write File
  cmdLines := String(30)
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
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	//Delete File
	err = os.Remove(cmdLines)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}
