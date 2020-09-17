package main
import (
  "log"
  "os/user"
  "strings"
)

func get_environment_state(environment string,region string) string {
  value := "setting-not-found"
  setting := "current_state"
  usr, err := user.Current()
  if err != nil {
    log.Fatal( err )
    value = "error"
  }
  path := usr.HomeDir + "/.onecloud/states/"+region+"/"+environment+"/state"
  fileSlice,err := readLines(path)
  if err != nil {
    log.Fatal(err)
    value = "error"
  } else {
    for i := 0; i < len(fileSlice); i++ {
      if strings.Contains(fileSlice[i],setting) {
        fields := strings.Split(fileSlice[i],":")
        value = fields[1]
      }
    }
  }
  return value
}
