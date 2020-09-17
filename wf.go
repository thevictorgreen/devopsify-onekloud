package main
import (
  "fmt"
  "os"
)

func wf(content []string,path string) {

  //Write File
	f, err := os.Create(path)
  if err != nil {
		fmt.Println(err)
		f.Close()
		return
  }
	for _, v := range content {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
  }
	err = f.Close()
  if err != nil {
		fmt.Println(err)
		return
  }
}
