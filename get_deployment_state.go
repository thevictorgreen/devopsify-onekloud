package main
import (
  "log"
  "os/user"
  "strings"
)

func get_deployment_state(environment string,region string) string {
  value := "setting-not-found"
  setting := "has_been_deployed"
  usr, err := user.Current()
  if err != nil {
    log.Fatal( err )
    value = "error"
  }
  path := usr.HomeDir + "/.onecloud/states/"+region+"/"+environment+"/deployment_state"
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
