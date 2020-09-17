package main
import (
    "log"
    "strings"
)

func get_ext_settings(setting string) string  {

  value := "setting-not-found"
  /*usr, err := user.Current()
  if err != nil {
    log.Fatal( err )
    value = "error"
  }*/
  path := "external_repos.txt"
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
