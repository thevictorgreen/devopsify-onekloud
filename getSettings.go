package main
import (
    "log"
    "os/user"
    "strings"
)

func getSettings(setting string) string  {

  value := "setting-not-found"
  usr, err := user.Current()
  if err != nil {
    log.Fatal( err )
    value = "error"
  }
  path := usr.HomeDir + "/.onecloud/onecloud.cfg"
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
