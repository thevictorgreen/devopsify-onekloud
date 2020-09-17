//aws_update_vpc.go
package main
import (
  "fmt"
)

func aws_update_vpc(region string, environment string) {
  fmt.Println("Updating AWS vpc in region: " + region + " for environment: " + environment)
}
