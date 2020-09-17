//onecloud.go
package main
import (
  "fmt"
  "log"
  "os"
  "github.com/urfave/cli"
  "strings"
)

var app = cli.NewApp()

func info() {
  app.Name = "onecloud"
  app.Usage = "One Interface. All Clouds"
  app.Author = "Victor Green"
  app.Version = "0.1.0"
}

func commands() {
  app.Commands = []cli.Command {
    {
      Name: "build",
      Usage: "Build Out Cloud",
      Flags: []cli.Flag {
        cli.StringFlag {
          Name: "provider,p",
          Usage: "Accepted Values: [aws,azure,gcloud]",
        },
        cli.StringFlag {
          Name: "item,i",
          Usage: "Accepted Values: [cloud,vpc]",
        },
        cli.StringFlag {
          Name: "region,r",
          Usage: "Accepted Values: [us-east-1,us-east-2,us-east-3,us-west-1,us-west-2]",
        },
        cli.StringFlag {
          Name: "environment,e",
          Usage: "Accepted Values: [management,development,staging,production]",
        },
        cli.StringFlag {
          Name: "domain,d",
          Usage: "Primary domain: [e.g. mycompany.com]",
        },
        cli.StringFlag {
          Name: "bucket,b",
          Usage: "AWS S3 bucket Prefix: [e.g. mycompany]",
        },
        cli.StringFlag {
          Name: "keypair,k",
          Usage: "AWS EC2 keypair: [e.g. my-kp]",
        },
        cli.StringFlag {
          Name: "action,a",
          Usage: "Accepted Values:[initialize,create,export,deploy,update,destroy,status]",
        },
      },
      Action: func(c *cli.Context) {
        //Amazon Web Services Provider
        if c.String("provider") == "aws" {
          //Cloud Item
          if c.String("item") == "cloud" {
            //Initialize action
            if c.String("action") == "initialize" {
              aws_initilize_cloud()
            } else
            //Action Not specified
            if c.String("action") == "" {
              fmt.Println("--action required: [initialize]")
            } else {
              //Invalid action
              fmt.Println("Invalid --action: [initialize]")
            }
          } else
          //VPC Item
          if c.String("item") == "vpc" {
            //Region
            if c.String("region") == "us-east-1" || c.String("region") == "us-east-2" || c.String("region") == "us-west-1" || c.String("region") == "us-west-2" {
              region := c.String("region")
              //Environment
              if c.String("environment") == "management" || c.String("environment") == "development" || c.String("environment") == "staging" || c.String("environment") == "production" {
                environment := c.String("environment")
                //Action
                if c.String("action") == "create" {
                  //Domain
                  if len(strings.TrimSpace(c.String("domain"))) != 0 {
                    domain := c.String("domain")
                    //Bucket
                    if len(strings.TrimSpace(c.String("bucket"))) != 0 {
                      bucket := c.String("bucket")
                      //Keypair
                      if len(strings.TrimSpace(c.String("keypair"))) != 0 {
                        keypair := c.String("keypair")
                        //fmt.Println(region+","+environment+","+domain+","+bucket+","+keypair)
                        aws_create_vpc(region,environment,domain,bucket,keypair)
                      } else {
                        fmt.Println("--keypair required: [e.g my-keypair]")
                      }
                    } else {
                      fmt.Println("--bucket required: [e.g mycompany]")
                    }
                  } else {
                    fmt.Println("--domain required: [e.g mycompany.com]")
                  }
                  //aws_create_vpc(region,environment)
                } else
                if c.String("action") == "export" {
                  aws_export_vpc(region,environment)
                } else
                if c.String("action") == "deploy" {
                  aws_deploy_vpc(region,environment)
                } else
                if c.String("action") == "update" {
                  aws_update_vpc(region,environment)
                } else
                if c.String("action") == "destroy" {
                  aws_destroy_vpc(region,environment)
                } else
                if c.String("action") == "status" {
                  aws_status_vpc(region,environment)
                } else
                //Action not specified
                if c.String("action") == "" {
                  fmt.Println("--action required: [create,export,deploy,update,destroy,status]")
                } else {
                  fmt.Println("Invalid --action: [create,export,deploy,update,destroy,status]")
                }
              } else
              //Environment not specified
              if c.String("environment") == "" {
                fmt.Println("--environment required: [management,development,staging,production]")
              } else {
                //Invalid environment
                fmt.Println("Invalid --environment: [management,development,staging,production]")
              }
            } else
            // Region Not Specified
            if c.String("region") == "" {
              fmt.Println("--region required: [us-east-1,us-east-2,us-west-1,us-west-2]")
            } else {
              //Invalid Region
              fmt.Println("Invalid --region: [us-east-1,us-east-2,us-west-1,us-west-2]")
            }
          } else
          //Item not specified
          if c.String("item") == "" {
            fmt.Println("--item required: [cloud,vpc]")
          } else {
            //Invalid Item
            fmt.Println("Invalid --item: [cloud,vpc]")
          }
        } else
        //Microsoft Azure Provider
        if c.String("provider") == "azure" {
          //Cloud Item
          if c.String("item") == "cloud" {
            //Initialize action
            if c.String("action") == "initialize" {
              azure_initilize_cloud()
            } else
            //Action Not specified
            if c.String("action") == "" {
              fmt.Println("--action required: [initialize]")
            } else {
              //Invalid action
              fmt.Println("Invalid --action: [initialize]")
            }
          } else
          //VPC Item
          if c.String("item") == "vpc" {
            //Region
            if c.String("region") == "us-east-1" || c.String("region") == "us-east-2" || c.String("region") == "us-west-1" || c.String("region") == "us-west-2" {
              region := c.String("region")
              //Environment
              if c.String("environment") == "management" || c.String("environment") == "development" || c.String("environment") == "staging" || c.String("environment") == "production" {
                environment := c.String("environment")
                //Action
                if c.String("action") == "create" {
                  azure_create_vpc(region,environment)
                } else
                if c.String("action") == "export" {
                  azure_export_vpc(region,environment)
                } else
                if c.String("action") == "deploy" {
                  azure_deploy_vpc(region,environment)
                } else
                if c.String("action") == "update" {
                  azure_update_vpc(region,environment)
                } else
                if c.String("action") == "destroy" {
                  azure_destroy_vpc(region,environment)
                } else
                if c.String("action") == "status" {
                  azure_status_vpc(region,environment)
                } else
                //Action not specified
                if c.String("action") == "" {
                  fmt.Println("--action required: [create,export,deploy,update,destroy,status]")
                } else {
                  fmt.Println("Invalid --action: [create,export,deploy,update,destroy,status]")
                }
              } else
              //Environment not specified
              if c.String("environment") == "" {
                fmt.Println("--environment required: [management,development,staging,production]")
              } else {
                //Invalid environment
                fmt.Println("Invalid --environment: [management,development,staging,production]")
              }
            } else
            // Region Not Specified
            if c.String("region") == "" {
              fmt.Println("--region required: [us-east-1,us-east-2,us-west-1,us-west-2]")
            } else {
              //Invalid Region
              fmt.Println("Invalid --region: [us-east-1,us-east-2,us-west-1,us-west-2]")
            }
          } else
          //Item not specified
          if c.String("item") == "" {
            fmt.Println("--item required: [cloud,vpc]")
          } else {
            //Invalid Item
            fmt.Println("Invalid --item: [cloud,vpc]")
          }
        } else
        //Google Cloud Provider
        if c.String("provider") == "gcloud" {
          //Cloud Item
          if c.String("item") == "cloud" {
            //Initialize action
            if c.String("action") == "initialize" {
              gcloud_initilize_cloud()
            } else
            //Action Not specified
            if c.String("action") == "" {
              fmt.Println("--action required: [initialize]")
            } else {
              //Invalid action
              fmt.Println("Invalid --action: [initialize]")
            }
          } else
          //VPC Item
          if c.String("item") == "vpc" {
            //Region
            if c.String("region") == "us-east-1" || c.String("region") == "us-east-2" || c.String("region") == "us-west-1" || c.String("region") == "us-west-2" {
              region := c.String("region")
              //Environment
              if c.String("environment") == "management" || c.String("environment") == "development" || c.String("environment") == "staging" || c.String("environment") == "production" {
                environment := c.String("environment")
                //Action
                if c.String("action") == "create" {
                  gcloud_create_vpc(region,environment)
                } else
                if c.String("action") == "export" {
                  gcloud_export_vpc(region,environment)
                } else
                if c.String("action") == "deploy" {
                  gcloud_deploy_vpc(region,environment)
                } else
                if c.String("action") == "update" {
                  gcloud_update_vpc(region,environment)
                } else
                if c.String("action") == "destroy" {
                  gcloud_destroy_vpc(region,environment)
                } else
                if c.String("action") == "status" {
                  gcloud_status_vpc(region,environment)
                } else
                //Action not specified
                if c.String("action") == "" {
                  fmt.Println("--action required: [create,export,deploy,update,destroy,status]")
                } else {
                  fmt.Println("Invalid --action: [create,export,deploy,update,destroy,status]")
                }
              } else
              //Environment not specified
              if c.String("environment") == "" {
                fmt.Println("--environment required: [management,development,staging,production]")
              } else {
                //Invalid environment
                fmt.Println("Invalid --environment: [management,development,staging,production]")
              }
            } else
            // Region Not Specified
            if c.String("region") == "" {
              fmt.Println("--region required: [us-east-1,us-east-2,us-west-1,us-west-2]")
            } else {
              //Invalid Region
              fmt.Println("Invalid --region: [us-east-1,us-east-2,us-west-1,us-west-2]")
            }
          } else
          //Item not specified
          if c.String("item") == "" {
            fmt.Println("--item required: [cloud,vpc]")
          } else {
            //Invalid Item
            fmt.Println("Invalid --item: [cloud,vpc]")
          }
        } else
        //Provider not specified
        if c.String("provider") == "" {
          fmt.Println("--provider required: [aws,azure,gcloud]")
        } else {
          //Invalid provider
          fmt.Println("Invalid provider: [aws,azure,gcloud]")
        }
      },
    },
  }
}

func main() {
  info()
  commands()
  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
