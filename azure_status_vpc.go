//azure_status_vpc.go
package main
import (
  "fmt"
)

func azure_status_vpc(region string, environment string) {
  fmt.Println("Azure vpc Status in region: " + region + " for environment: " + environment)
}
