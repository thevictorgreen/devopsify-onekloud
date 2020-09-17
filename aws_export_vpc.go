//aws_export_vpc.go
package main
import (
  "fmt"
)

func aws_export_vpc(region string, environment string) {
  //fmt.Println("Exporting AWS vpc in region: " + region + " for environment: " + environment)

  if check_aws_initialized() == true {
    if get_environment_state(environment,region) != "initialized" {
      if check_aws_vpc_config(region,environment) == 0 {
        export_aws_vpc_settings(region,environment)
        //UPDATE ENVIRONMENT'S STATE
        update_environment_state("current_state:exported",environment,region)
        fmt.Println("AWS VPC " +region+ " " + environment + " environment exported")
      } else {
        fmt.Println("Preflight check failed")
        fmt.Println("Check the following")
        fmt.Println("  1: AWS/"+region+"/"+environment+"/networking/vpc-config")
        fmt.Println("  2: AWS/"+region+"/"+environment+"/networking/public-nacl-config")
        fmt.Println("  3: AWS/"+region+"/"+environment+"/networking/public-subnets-config")
        fmt.Println("  4: AWS/"+region+"/"+environment+"/networking/public-security-group-config")
      }
    } else {
      fmt.Println("AWS VPC for " + environment + " in region:" + region + " must be created. run: onecloud build --provider aws --item vpc --region "+region+" --environment "+environment+" --action create")
    }
  } else {
    fmt.Println("aws must be initialized first. run: onecloud build --provider aws --item cloud --action initialize")
  }
}
