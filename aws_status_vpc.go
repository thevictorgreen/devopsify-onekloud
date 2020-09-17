//aws_status_vpc.go
package main
import (
  "fmt"
)

/*func aws_status_vpc(region string, environment string) {
  fmt.Println("AWS vpc Status in region: " + region + " for environment: " + environment)
}*/

func aws_status_vpc(region string,environment string)  {
  if check_aws_initialized() == true {
    fmt.Println( get_environment_state(environment,region) )
  } else {
    fmt.Println("Provider AWS Not Initialized")
  }
}
