//aws_initilize_cloud.go
package main
import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "os/user"
)

func aws_initilize_cloud() {

  cmd := exec.Command("/bin/bash", "-c", "mkdir AWS")
  cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
  if err != nil {
    log.Fatalf("%s\n", err)
  } else {
    //CHECK FOR CONFIG FOLDER AND FILE. IF NOT THERE CREATE.
    usr, err := user.Current()
    if err != nil {
      log.Fatal( err )
    }
    path := usr.HomeDir + "/.onecloud"
    if _, err := os.Stat(path); os.IsNotExist(err) {
      os.Mkdir(path, 0755)

      spec := []string {
        "curl -O https://raw.githubusercontent.com/thevictorgreen/devopsify-external-repo-versions/develop/external_repos.txt",
      }
      es(spec)

      ansible_root_project_ver := get_ext_settings("ansible_root_project_ver")
      terraform_ec2_ver := get_ext_settings("terraform_ec2_ver")
      us_east_1_vpc_ver := get_ext_settings("us-east-1_vpc_ver")

      spec = []string {
        "## ONECLOUD Configuation Options",
        "ansible_root_project_ver:"+ansible_root_project_ver,
        "terraform_ec2_ver:"+terraform_ec2_ver,
        "us-east-1_vpc_ver:"+us_east_1_vpc_ver,
        "ansible_root_project_src:thevictorgreen/devopsify-ansible-bootstrap/tar.gz",
        "us-east-1_vpc_src:thevictorgreen/devopsify-terraform-impl-aws-us-east-1-vpc/tar.gz",
        "us-east-1_region_id:001useast1",
        "terraform_ec2_src:thevictorgreen/devopsify-terraform-impl-aws-ec2-bootstrap/tar.gz",

        "us-east-1_management_cidrblock:10.10.0.0",
        "us-east-1_management_cidrslash:16",

        "us-east-1_development_cidrblock:10.20.0.0",
        "us-east-1_development_cidrslash:16",

        "us-east-1_staging_cidrblock:10.30.0.0",
        "us-east-1_staging_cidrslash:16",

        "us-east-1_production_cidrblock:10.40.0.0",
        "us-east-1_production_cidrslash:16",

        "us-east-2_vpc_src:thevictorgreen/devopsify-terraform-aws-us-east-2-vpc/tar.gz",
        "us-east-2_vpc_ver:0.1.1",
        "us-east-2_region_id:002useast2",

        "us-east-2_management_cidrblock:10.50.0.0",
        "us-east-2_management_cidrslash:16",

        "us-east-2_development_cidrblock:10.60.0.0",
        "us-east-2_development_cidrslash:16",

        "us-east-2_staging_cidrblock:10.70.0.0",
        "us-east-2_staging_cidrslash:16",

        "us-east-2_production_cidrblock:10.80.0.0",
        "us-east-2_production_cidrslash:16",

        "us-west-1_vpc_src:thevictorgreen/devopsify-terraform-aws-us-west-1-vpc/tar.gz",
        "us-west-1_vpc_ver:0.1.1",
        "us-west-1_region_id:001uswest1",

        "us-west-1_management_cidrblock:10.90.0.0",
        "us-west-1_management_cidrslash:16",

        "us-west-1_development_cidrblock:10.100.0.0",
        "us-west-1_development_cidrslash:16",

        "us-west-1_staging_cidrblock:10.110.0.0",
        "us-west-1_staging_cidrslash:16",

        "us-west-1_production_cidrblock:10.120.0.0",
        "us-west-1_production_cidrslash:16",

        "us-west-2_vpc_src:thevictorgreen/devopsify-terraform-aws-us-west-2-vpc/tar.gz",
        "us-west-2_vpc_ver:0.1.0",
        "us-west-2_region_id:002uswest2",

        "us-west-2_management_cidrblock:10.130.0.0",
        "us-west-2_management_cidrslash:16",

        "us-west-2_development_cidrblock:10.140.0.0",
        "us-west-2_development_cidrslash:16",

        "us-west-2_staging_cidrblock:10.150.0.0",
        "us-west-2_staging_cidrslash:16",

        "us-west-2_production_cidrblock:10.160.0.0",
        "us-west-2_production_cidrslash:16",

      }
      wf(spec,path+"/onecloud.cfg")

      spec = []string {
        "rm -rf external_repos.txt",
      }
      es(spec)

      spec = []string {
        "mkdir -p "+path+"/states/us-east-1/management/networking",
        "mkdir -p "+path+"/states/us-east-1/development/networking",
        "mkdir -p "+path+"/states/us-east-1/staging/networking",
        "mkdir -p "+path+"/states/us-east-1/production/networking",

        "mkdir -p "+path+"/states/us-east-2/management/networking",
        "mkdir -p "+path+"/states/us-east-2/development/networking",
        "mkdir -p "+path+"/states/us-east-2/staging/networking",
        "mkdir -p "+path+"/states/us-east-2/production/networking",

        "mkdir -p "+path+"/states/us-west-1/management/networking",
        "mkdir -p "+path+"/states/us-west-1/development/networking",
        "mkdir -p "+path+"/states/us-west-1/staging/networking",
        "mkdir -p "+path+"/states/us-west-1/production/networking",

        "mkdir -p "+path+"/states/us-west-2/management/networking",
        "mkdir -p "+path+"/states/us-west-2/development/networking",
        "mkdir -p "+path+"/states/us-west-2/staging/networking",
        "mkdir -p "+path+"/states/us-west-2/production/networking",
      }
      es(spec)
      spec = []string {
        "touch "+path+"/states/state",
        "echo current_state:initialized > "+path+"/states/state",
      }
      es(spec)
      // us-east-1 management public
      spec = []string {
        "subnet-count:6",

        "sub0-cidr:10.10.0.0",
        "sub0-slash:24",
        "sub0-zone:us-east-1a",

        "sub1-cidr:10.10.1.0",
        "sub1-slash:24",
        "sub1-zone:us-east-1b",

        "sub2-cidr:10.10.2.0",
        "sub2-slash:24",
        "sub2-zone:us-east-1c",

        "sub3-cidr:10.10.3.0",
        "sub3-slash:24",
        "sub3-zone:us-east-1d",

        "sub4-cidr:10.10.4.0",
        "sub4-slash:24",
        "sub4-zone:us-east-1e",

        "sub5-cidr:10.10.5.0",
        "sub5-slash:24",
        "sub5-zone:us-east-1f",
      }
      wf(spec,path+"/states/us-east-1/management/networking/public-subnets-config")
      // us-east-1 management private
      spec = []string {
        "subnet-count:6",
        "sub0-cidr:10.10.6.0",
        "sub0-slash:24",
        "sub0-zone:us-east-1a",

        "sub1-cidr:10.10.7.0",
        "sub1-slash:24",
        "sub1-zone:us-east-1b",

        "sub2-cidr:10.10.8.0",
        "sub2-slash:24",
        "sub2-zone:us-east-1c",

        "sub3-cidr:10.10.9.0",
        "sub3-slash:24",
        "sub3-zone:us-east-1d",

        "sub4-cidr:10.10.10.0",
        "sub4-slash:24",
        "sub4-zone:us-east-1e",

        "sub5-cidr:10.10.11.0",
        "sub5-slash:24",
        "sub5-zone:us-east-1f",
      }
      wf(spec,path+"/states/us-east-1/management/networking/private-subnets-config")
      // us-east-1 development public
      spec = []string {
        "subnet-count:6",

        "sub0-cidr:10.20.0.0",
        "sub0-slash:24",
        "sub0-zone:us-east-1a",

        "sub1-cidr:10.20.1.0/",
        "sub1-slash:24",
        "sub1-zone:us-east-1b",

        "sub2-cidr:10.20.2.0",
        "sub2-slash:24",
        "sub2-zone:us-east-1c",

        "sub3-cidr:10.20.3.0",
        "sub3-slash:24",
        "sub3-zone:us-east-1d",

        "sub4-cidr:10.20.4.0",
        "sub4-slash:24",
        "sub4-zone:us-east-1e",

        "sub5-cidr:10.20.5.0",
        "sub5-slash:24",
        "sub5-zone:us-east-1f",
      }
      wf(spec,path+"/states/us-east-1/development/networking/public-subnets-config")
      // us-east-1 development private
      spec = []string {
        "subnet-count:6",

        "sub0-cidr:10.20.6.0",
        "sub0-slash:24",
        "sub0-zone:us-east-1a",

        "sub1-cidr:10.20.7.0",
        "sub1-slash:24",
        "sub1-zone:us-east-1b",

        "sub2-cidr:10.20.8.0",
        "sub2-slash:24",
        "sub2-zone:us-east-1c",

        "sub3-cidr:10.20.9.0",
        "sub3-slash:24",
        "sub3-zone:us-east-1d",

        "sub4-cidr:10.20.10.0",
        "sub4-slash:24",
        "sub4-zone:us-east-1e",

        "sub5-cidr:10.20.11.0",
        "sub5-slash:24",
        "sub5-zone:us-east-1f",
      }
      wf(spec,path+"/states/us-east-1/development/networking/private-subnets-config")
      // us-east-1 staging public
      spec = []string {
        "subnet-count:6",

        "sub0-cidr:10.30.0.0",
        "sub0-slash:24",
        "sub0-zone:us-east-1a",

        "sub1-cidr:10.30.1.0",
        "sub1-slash:24",
        "sub1-zone:us-east-1b",

        "sub2-cidr:10.30.2.0",
        "sub2-slash:24",
        "sub2-zone:us-east-1c",

        "sub3-cidr:10.30.3.0",
        "sub3-slash:24",
        "sub3-zone:us-east-1d",

        "sub4-cidr:10.30.4.0",
        "sub4-slash:24",
        "sub4-zone:us-east-1e",

        "sub5-cidr:10.30.5.0",
        "sub5-slash:24",
        "sub5-zone:us-east-1f",
      }
      wf(spec,path+"/states/us-east-1/staging/networking/public-subnets-config")
      // us-east-1 staging private
      spec = []string {
        "subnet-count:6",

        "sub0-cidr:10.30.6.0",
        "sub0-slash:24",
        "sub0-zone:us-east-1a",

        "sub1-cidr:10.30.7.0",
        "sub1-slash:24",
        "sub1-zone:us-east-1b",

        "sub2-cidr:10.30.8.0",
        "sub2-slash:24",
        "sub2-zone:us-east-1c",

        "sub3-cidr:10.30.9.0",
        "sub3-slash:24",
        "sub3-zone:us-east-1d",

        "sub4-cidr:10.30.10.0",
        "sub4-slash:24",
        "sub4-zone:us-east-1e",

        "sub5-cidr:10.30.11.0",
        "sub5-slash:24",
        "sub5-zone:us-east-1f",
      }
      wf(spec,path+"/states/us-east-1/staging/networking/private-subnets-config")
      // us-east-1 production public
      spec = []string {
        "subnet-count:6",

        "sub0-cidr:10.40.0.0",
        "sub0-zone:us-east-1a",

        "sub1-cidr:10.40.1.0",
        "sub1-zone:us-east-1b",

        "sub2-cidr:10.40.2.0",
        "sub2-zone:us-east-1c",

        "sub3-cidr:10.40.3.0",
        "sub3-zone:us-east-1d",

        "sub4-cidr:10.40.4.0",
        "sub4-zone:us-east-1e",

        "sub5-cidr:10.40.5.0",
        "sub5-zone:us-east-1f",
      }
      wf(spec,path+"/states/us-east-1/production/networking/public-subnets-config")
      // us-east-1 production private
      spec = []string {
        "subnet-count:6",

        "sub0-cidr:10.40.6.0",
        "sub0-slash:24",
        "sub0-zone:us-east-1a",

        "sub1-cidr:10.40.7.0",
        "sub1-slash:24",
        "sub1-zone:us-east-1b",

        "sub2-cidr:10.40.8.0",
        "sub2-slash:24",
        "sub2-zone:us-east-1c",

        "sub3-cidr:10.40.9.0",
        "sub3-slash:24",
        "sub3-zone:us-east-1d",

        "sub4-cidr:10.40.10.0",
        "sub4-slash:24",
        "sub4-zone:us-east-1e",

        "sub5-cidr:10.40.11.0",
        "sub5-slash:24",
        "sub5-zone:us-east-1f",
      }
      wf(spec,path+"/states/us-east-1/production/networking/private-subnets-config")
      // us-east-2 management public
      spec = []string {
        "subnet-count:3",
        "sub0-cidr:10.50.0.0/24",
        "sub0-zone:us-east-2a",
        "sub1-cidr:10.50.1.0/24",
        "sub1-zone:us-east-2b",
        "sub2-cidr:10.50.2.0/24",
        "sub2-zone:us-east-2c",
      }
      wf(spec,path+"/states/us-east-2/management/networking/public-subnets-config")
      // us-east-2 management private
      spec = []string {
        "subnet-count:3",
        "sub0-cidr:10.50.3.0/24",
        "sub0-zone:us-east-2a",
        "sub1-cidr:10.50.4.0/24",
        "sub1-zone:us-east-2b",
        "sub2-cidr:10.50.5.0/24",
        "sub2-zone:us-east-2c",
      }
      wf(spec,path+"/states/us-east-2/management/networking/private-subnets-config")
      // us-east-2 development public
      spec = []string {
        "subnet-count:3",
        "sub0-cidr:10.60.0.0/24",
        "sub0-zone:us-east-2a",
        "sub1-cidr:10.60.1.0/24",
        "sub1-zone:us-east-2b",
        "sub2-cidr:10.60.2.0/24",
        "sub2-zone:us-east-2c",
      }
      wf(spec,path+"/states/us-east-2/development/networking/public-subnets-config")
      // us-east-2 development private
      spec = []string {
        "subnet-count:3",
        "sub0-cidr:10.60.3.0/24",
        "sub0-zone:us-east-2a",
        "sub1-cidr:10.60.4.0/24",
        "sub1-zone:us-east-2b",
        "sub2-cidr:10.60.5.0/24",
        "sub2-zone:us-east-2c",
      }
      wf(spec,path+"/states/us-east-2/development/networking/private-subnets-config")
      // us-east-2 staging public
      spec = []string {
        "subnet-count:3",
        "sub0-cidr:10.70.0.0/24",
        "sub0-zone:us-east-2a",
        "sub1-cidr:10.70.1.0/24",
        "sub1-zone:us-east-2b",
        "sub2-cidr:10.70.2.0/24",
        "sub2-zone:us-east-2c",
      }
      wf(spec,path+"/states/us-east-2/staging/networking/public-subnets-config")
      // us-east-2 staging private
      spec = []string {
        "subnet-count:3",
        "sub0-cidr:10.70.3.0/24",
        "sub0-zone:us-east-2a",
        "sub1-cidr:10.70.4.0/24",
        "sub1-zone:us-east-2b",
        "sub2-cidr:10.70.5.0/24",
        "sub2-zone:us-east-2c",
      }
      wf(spec,path+"/states/us-east-2/staging/networking/private-subnets-config")
      // us-east-2 production public
      spec = []string {
        "subnet-count:3",
        "sub0-cidr:10.80.0.0/24",
        "sub0-zone:us-east-2a",
        "sub1-cidr:10.80.1.0/24",
        "sub1-zone:us-east-2b",
        "sub2-cidr:10.80.2.0/24",
        "sub2-zone:us-east-2c",
      }
      wf(spec,path+"/states/us-east-2/production/networking/public-subnets-config")
      // us-east-2 production private
      spec = []string {
        "subnet-count:3",
        "sub0-cidr:10.80.3.0/24",
        "sub0-zone:us-east-2a",
        "sub1-cidr:10.80.4.0/24",
        "sub1-zone:us-east-2b",
        "sub2-cidr:10.80.5.0/24",
        "sub2-zone:us-east-2c",
      }
      wf(spec,path+"/states/us-east-2/production/networking/private-subnets-config")
      // us-west-1 management public
      spec = []string {
        "subnet-count:2",
        "sub0-cidr:10.90.0.0/24",
        "sub0-zone:us-west-1a",
        "sub1-cidr:10.90.1.0/24",
        "sub1-zone:us-west-1b",
      }
      wf(spec,path+"/states/us-west-1/management/networking/public-subnets-config")
      // us-west-1 management private
      spec = []string {
        "subnet-count:2",
        "sub0-cidr:10.90.3.0/24",
        "sub0-zone:us-west-1a",
        "sub1-cidr:10.90.4.0/24",
        "sub1-zone:us-west-1b",
      }
      wf(spec,path+"/states/us-west-1/management/networking/private-subnets-config")
      // us-west-1 development public
      spec = []string {
        "subnet-count:2",
        "sub0-cidr:10.100.0.0/24",
        "sub0-zone:us-west-1a",
        "sub1-cidr:10.100.1.0/24",
        "sub1-zone:us-west-1b",
      }
      wf(spec,path+"/states/us-west-1/development/networking/public-subnets-config")
      // us-west-1 development private
      spec = []string {
        "subnet-count:2",
        "sub0-cidr:10.100.3.0/24",
        "sub0-zone:us-west-1a",
        "sub1-cidr:10.100.4.0/24",
        "sub1-zone:us-west-1b",
      }
      wf(spec,path+"/states/us-west-1/development/networking/private-subnets-config")
      // us-west-1 staging public
      spec = []string {
        "subnet-count:2",
        "sub0-cidr:10.110.0.0/24",
        "sub0-zone:us-west-1a",
        "sub1-cidr:10.110.1.0/24",
        "sub1-zone:us-west-1b",
      }
      wf(spec,path+"/states/us-west-1/staging/networking/public-subnets-config")
      // us-west-1 staging private
      spec = []string {
        "subnet-count:2",
        "sub0-cidr:10.110.3.0/24",
        "sub0-zone:us-west-1a",
        "sub1-cidr:10.110.4.0/24",
        "sub1-zone:us-west-1b",
      }
      wf(spec,path+"/states/us-west-1/staging/networking/private-subnets-config")
      // us-west-1 production public
      spec = []string {
        "subnet-count:2",
        "sub0-cidr:10.120.0.0/24",
        "sub0-zone:us-west-1a",
        "sub1-cidr:10.120.1.0/24",
        "sub1-zone:us-west-1b",
      }
      wf(spec,path+"/states/us-west-1/production/networking/public-subnets-config")
      // us-west-1 production private
      spec = []string {
        "subnet-count:2",
        "sub0-cidr:10.120.3.0/24",
        "sub0-zone:us-west-1a",
        "sub1-cidr:10.120.4.0/24",
        "sub1-zone:us-west-1b",
      }
      wf(spec,path+"/states/us-west-1/production/networking/private-subnets-config")
      // us-west-2 mangement public
      spec = []string {
        "subnet-count:4",
        "sub0-cidr:10.130.0.0/24",
        "sub0-zone:us-west-2a",
        "sub1-cidr:10.130.1.0/24",
        "sub1-zone:us-west-2b",
        "sub2-cidr:10.130.2.0/24",
        "sub2-zone:us-west-2c",
        "sub3-cidr:10.130.3.0/24",
        "sub3-zone:us-west-2d",
      }
      wf(spec,path+"/states/us-west-2/management/networking/public-subnets-config")
      // us-west-2 mangement private
      spec = []string {
        "subnet-count:4",
        "sub0-cidr:10.130.4.0/24",
        "sub0-zone:us-west-2a",
        "sub1-cidr:10.130.5.0/24",
        "sub1-zone:us-west-2b",
        "sub2-cidr:10.130.6.0/24",
        "sub2-zone:us-west-2c",
        "sub3-cidr:10.130.7.0/24",
        "sub3-zone:us-west-2d",
      }
      wf(spec,path+"/states/us-west-2/management/networking/private-subnets-config")
      // us-west-2 development public
      spec = []string {
        "subnet-count:4",
        "sub0-cidr:10.140.0.0/24",
        "sub0-zone:us-west-2a",
        "sub1-cidr:10.140.1.0/24",
        "sub1-zone:us-west-2b",
        "sub2-cidr:10.140.2.0/24",
        "sub2-zone:us-west-2c",
        "sub3-cidr:10.140.3.0/24",
        "sub3-zone:us-west-2d",
      }
      wf(spec,path+"/states/us-west-2/development/networking/public-subnets-config")
      // us-west-2 development private
      spec = []string {
        "subnet-count:4",
        "sub0-cidr:10.140.4.0/24",
        "sub0-zone:us-west-2a",
        "sub1-cidr:10.140.5.0/24",
        "sub1-zone:us-west-2b",
        "sub2-cidr:10.140.6.0/24",
        "sub2-zone:us-west-2c",
        "sub3-cidr:10.140.7.0/24",
        "sub3-zone:us-west-2d",
      }
      wf(spec,path+"/states/us-west-2/development/networking/private-subnets-config")
      // us-west-2 staging public
      spec = []string {
        "subnet-count:4",
        "sub0-cidr:10.150.0.0/24",
        "sub0-zone:us-west-2a",
        "sub1-cidr:10.150.1.0/24",
        "sub1-zone:us-west-2b",
        "sub2-cidr:10.150.2.0/24",
        "sub2-zone:us-west-2c",
        "sub3-cidr:10.150.3.0/24",
        "sub3-zone:us-west-2d",
      }
      wf(spec,path+"/states/us-west-2/staging/networking/public-subnets-config")
      // us-west-2 staging private
      spec = []string {
        "subnet-count:4",
        "sub0-cidr:10.150.4.0/24",
        "sub0-zone:us-west-2a",
        "sub1-cidr:10.150.5.0/24",
        "sub1-zone:us-west-2b",
        "sub2-cidr:10.150.6.0/24",
        "sub2-zone:us-west-2c",
        "sub3-cidr:10.150.7.0/24",
        "sub3-zone:us-west-2d",
      }
      wf(spec,path+"/states/us-west-2/staging/networking/private-subnets-config")
      // us-west-2 production public
      spec = []string {
        "subnet-count:4",
        "sub0-cidr:10.160.0.0/24",
        "sub0-zone:us-west-2a",
        "sub1-cidr:10.160.1.0/24",
        "sub1-zone:us-west-2b",
        "sub2-cidr:10.160.2.0/24",
        "sub2-zone:us-west-2c",
        "sub3-cidr:10.160.3.0/24",
        "sub3-zone:us-west-2d",
      }
      wf(spec,path+"/states/us-west-2/production/networking/public-subnets-config")
      // us-west-2 production private
      spec = []string {
        "subnet-count:4",
        "sub0-cidr:10.160.4.0/24",
        "sub0-zone:us-west-2a",
        "sub1-cidr:10.160.5.0/24",
        "sub1-zone:us-west-2b",
        "sub2-cidr:10.160.6.0/24",
        "sub2-zone:us-west-2c",
        "sub3-cidr:10.160.7.0/24",
        "sub3-zone:us-west-2d",
      }
      wf(spec,path+"/states/us-west-2/production/networking/private-subnets-config")
      spec = []string {
        "current_state:initialized",
      }
      wf(spec,path+"/states/us-east-1/management/state")
      wf(spec,path+"/states/us-east-1/development/state")
      wf(spec,path+"/states/us-east-1/staging/state")
      wf(spec,path+"/states/us-east-1/production/state")
      wf(spec,path+"/states/us-east-2/management/state")
      wf(spec,path+"/states/us-east-2/development/state")
      wf(spec,path+"/states/us-east-2/staging/state")
      wf(spec,path+"/states/us-east-2/production/state")
      wf(spec,path+"/states/us-west-1/management/state")
      wf(spec,path+"/states/us-west-1/development/state")
      wf(spec,path+"/states/us-west-1/staging/state")
      wf(spec,path+"/states/us-west-1/production/state")
      wf(spec,path+"/states/us-west-2/management/state")
      wf(spec,path+"/states/us-west-2/development/state")
      wf(spec,path+"/states/us-west-2/staging/state")
      wf(spec,path+"/states/us-west-2/production/state")
      spec = []string {
        "has_been_deployed:false",
      }
      wf(spec,path+"/states/us-east-1/management/deployment_state")
      wf(spec,path+"/states/us-east-1/development/deployment_state")
      wf(spec,path+"/states/us-east-1/staging/deployment_state")
      wf(spec,path+"/states/us-east-1/production/deployment_state")
      wf(spec,path+"/states/us-east-2/management/deployment_state")
      wf(spec,path+"/states/us-east-2/development/deployment_state")
      wf(spec,path+"/states/us-east-2/staging/deployment_state")
      wf(spec,path+"/states/us-east-2/production/deployment_state")
      wf(spec,path+"/states/us-west-1/management/deployment_state")
      wf(spec,path+"/states/us-west-1/development/deployment_state")
      wf(spec,path+"/states/us-west-1/staging/deployment_state")
      wf(spec,path+"/states/us-west-1/production/deployment_state")
      wf(spec,path+"/states/us-west-2/management/deployment_state")
      wf(spec,path+"/states/us-west-2/development/deployment_state")
      wf(spec,path+"/states/us-west-2/staging/deployment_state")
      wf(spec,path+"/states/us-west-2/production/deployment_state")
    }
    fmt.Println("AWS Cloud Initialized")
  }
}
