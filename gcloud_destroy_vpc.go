//gcloud_destroy_vpc.go
package main
import (
  "fmt"
)

func gcloud_destroy_vpc(region string, environment string) {
  fmt.Println("Destroying GCloud vpc in region: " + region + " for environment: " + environment)
}
