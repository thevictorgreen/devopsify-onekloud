//check_aws_initialized.go
package main
import (
  "log"
  "os"
  "os/user"
)

func check_aws_initialized() bool {
  //CHECK FOR CONFIG FOLDER AND FILE.
  usr, err := user.Current()
  if err != nil {
    log.Fatal( err )
  }
  path := usr.HomeDir + "/.onecloud/onecloud.cfg"
  if _, err := os.Stat(path); os.IsNotExist(err) {
    return false
  } else {
    return true
  }
}
