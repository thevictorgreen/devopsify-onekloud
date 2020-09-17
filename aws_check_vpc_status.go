//aws_check_vpc_status.go
package main
import (
  "fmt"
)

func aws_check_vpc_status(region string,environment string)  {
  if check_aws_initialized() == true {
    fmt.Println( get_environment_state(environment,region) )
  } else {
    fmt.Println("Provider AWS Not Initialized")
  }
}
