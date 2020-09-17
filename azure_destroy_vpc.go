//azure_destroy_vpc.go
package main
import (
  "fmt"
)

func azure_destroy_vpc(region string, environment string) {
  fmt.Println("Destroying Azure vpc in region: " + region + " for environment: " + environment)
}
