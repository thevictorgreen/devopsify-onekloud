//gcloud_status_vpc.go
package main
import (
  "fmt"
)

func gcloud_status_vpc(region string, environment string) {
  fmt.Println("GCloud vpc Status in region: " + region + " for environment: " + environment)
}
