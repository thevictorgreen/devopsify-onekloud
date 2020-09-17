package main
import (
  "log"
  "os/user"
)

func update_deployment_state(state string,environment string,region string) {
  //GET .onecloud CONFIG DIRECTORY FROM USER HOME ~
  usr, err := user.Current()
  if err != nil {
    log.Fatal( err )
  }
  //SET PATH
  path := usr.HomeDir + "/.onecloud"
  //UPDATE STATE
  spec := []string {
    "echo "+state+" > "+path+"/states/"+region+"/"+environment+"/deployment_state",
  }
  es(spec)
}
