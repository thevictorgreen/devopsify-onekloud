//aws_destroy_vpc.go
package main
import (
  "fmt"
)

func aws_destroy_vpc(region string, environment string) {
  fmt.Println("Destroying AWS vpc in region: " + region + " for environment: " + environment)
}
