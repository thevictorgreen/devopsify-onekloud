//aws_deploy_vpc.go
package main
import (
  "fmt"
)

func aws_deploy_vpc(region string, environment string) {
  fmt.Println("Deploying AWS vpc in region: " + region + " for environment: " + environment)
}
