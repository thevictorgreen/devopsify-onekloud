//aws_create_vpc.go
package main
import (
  "fmt"
  "os/user"
  "log"
)

func aws_create_vpc(region string, environment string,domain string,bucket string,keypair string) {
  //fmt.Println("Creating AWS vpc in region: " + region + " for environment: " + environment)

  if check_aws_initialized() == true {
    if get_environment_state(environment,region) == "initialized" {
      //GET USER OBJECT
      usr, err := user.Current()
      if err != nil {
        log.Fatal( err )
      }

      //GET UNIQUE REGION
      region_id := getSettings(region+"_region_id")

      // GET VPC CIDRBLOCK
      cidrblock := getSettings(region+"_"+environment+"_"+"cidrblock")
      cidrslash := getSettings(region+"_"+environment+"_"+"cidrslash")

      //CREATE FOLDERS
      spec := []string {
    		"mkdir -p AWS/"+region+"/"+environment+"/credentials",
    	  "mkdir -p AWS/"+region+"/"+environment+"/iac-cloudformation",
    	  "mkdir -p AWS/"+region+"/"+environment+"/iac-ansible",
    	  "mkdir -p AWS/"+region+"/"+environment+"/iac-packer",
    	  "mkdir -p AWS/"+region+"/"+environment+"/iac-terraform",
    	  "mkdir -p AWS/"+region+"/"+environment+"/networking",
    	}
    	es(spec)

      //INSTALL ANSIBLE ROOT PROJECT FOLDER
      ansible_root_project_src := getSettings("ansible_root_project_src")
      ansible_root_project_ver := getSettings("ansible_root_project_ver")
      spec = []string {
        "curl -o ansible.tar.gz https://codeload.github.com/"+ansible_root_project_src+"/v"+ansible_root_project_ver,
        "tar -xvf ansible.tar.gz",
        "mv devopsify-ansible-bootstrap-"+ansible_root_project_ver+" iac-ansible",
        "mv iac-ansible AWS/"+region+"/"+environment+"/",
        "rm -rf ansible.tar.gz",
        "echo "+environment+" > AWS/"+region+"/"+environment+"/iac-ansible/environment.cfg",
      }
      es(spec)

      //INSTALL TERRAFORM VPC FOLDER
      terraform_root_project_src := getSettings(region+"_vpc_src")
      terraform_root_project_ver := getSettings(region+"_vpc_ver")
      spec = []string {
        "curl -o terraform.tar.gz https://codeload.github.com/"+terraform_root_project_src+"/v"+terraform_root_project_ver,
        "tar -xvf terraform.tar.gz",
        "mv devopsify-terraform-impl-aws-"+region+"-vpc-"+terraform_root_project_ver+" iac-terraform",
        "mv iac-terraform AWS/"+region+"/"+environment+"/",
        "rm -rf terraform.tar.gz",
        "echo "+environment+" > AWS/"+region+"/"+environment+"/iac-terraform/environment.cfg",
        "mkdir -p AWS/"+region+"/"+environment+"/iac-terraform/y",
    	  "mkdir -p AWS/"+region+"/"+environment+"/iac-terraform/z",
      }
      es(spec)

      //INSTALL TERRAFORM EC2 INSTANCES FOLDER
      terraform_ec2_src := getSettings("terraform_ec2_src")
      terraform_ec2_ver := getSettings("terraform_ec2_ver")
      spec = []string {
        "curl -o terraform.tar.gz https://codeload.github.com/"+terraform_ec2_src+"/v"+terraform_ec2_ver,
        "tar -xvf terraform.tar.gz",
        "mv devopsify-terraform-impl-aws-ec2-bootstrap-"+terraform_ec2_ver+" ec2_templates",
        "mv ec2_templates AWS/"+region+"/"+environment+"/iac-terraform/",
        "rm -rf terraform.tar.gz",
      }
      es(spec)


      //UPDATE ENVIRONMENT IN inf_networking.tf, inv_servers.tf, and machine_sample.tf
      spec = []string {
        "sed -i \"\" 's/AAAAA/"+environment+"/g' AWS/"+region+"/"+environment+"/iac-terraform/output.tf",
        "sed -i \"\" 's/AAAAA/"+environment+"/g' AWS/"+region+"/"+environment+"/iac-terraform/tf_backend.tf",
        "sed -i \"\" 's/AAAAA/"+environment+"/g' AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
        "find AWS/"+region+"/"+environment+"/iac-terraform/ec2_templates -type f -exec sed -i \"\" 's/AAAAA/"+environment+"/g' {} \\;",
        //find AWS/us-east-1/management/iac-terraform/ec2_templates -type f -exec sed -i "" "s/AAAAA/management/g" {} \;
      }
      es(spec)

      //DOWNLOAD ANSIBLE ROLES
      spec = []string {
        "cd AWS/"+region+"/"+environment+"/iac-ansible/",
        "ansible-galaxy install -r requirements.yml",
      }
      es(spec)

      first_octets := ""
      //FOR us-east-1
      if region == "us-east-1" && environment == "management" {
        first_octets = "10.10"
      } else
      if region == "us-east-1" && environment == "development" {
        first_octets = "20.10"
      } else
      if region == "us-east-1" && environment == "staging" {
        first_octets = "30.10"
      } else
      if region == "us-east-1" && environment == "production" {
        first_octets = "40.10"
      }
      //FOR us-east-2
      if region == "us-east-2" && environment == "management" {
        first_octets = "50.10"
      } else
      if region == "us-east-2" && environment == "development" {
        first_octets = "60.10"
      } else
      if region == "us-east-2" && environment == "staging" {
        first_octets = "70.10"
      } else
      if region == "us-east-2" && environment == "production" {
        first_octets = "80.10"
      }
      //FOR us-west-1
      if region == "us-west-1" && environment == "management" {
        first_octets = "90.10"
      } else
      if region == "us-west-1" && environment == "development" {
        first_octets = "100.10"
      } else
      if region == "us-west-1" && environment == "staging" {
        first_octets = "110.10"
      } else
      if region == "us-west-1" && environment == "production" {
        first_octets = "120.10"
      }
      //For us-west-2
      if region == "us-west-2" && environment == "management" {
        first_octets = "130.10"
      } else
      if region == "us-west-2" && environment == "development" {
        first_octets = "140.10"
      } else
      if region == "us-west-2" && environment == "staging" {
        first_octets = "150.10"
      } else
      if region == "us-west-2" && environment == "production" {
        first_octets = "160.10"
      }

      //CREATE VPC
      path := "AWS/"+region+"/"+environment+"/networking/vpc-config"
      spec = []string {
        "name:"+environment+region_id,
        "cidr-block:"+cidrblock,
        "cidr-slash:"+cidrslash,
        "ec2-keypair:"+keypair,
        "primary-domain:"+domain,
        "private-zone:"+environment+"."+domain+".",
        "backward-zone:"+first_octets+".in-addr.arpa.",
        "public-zone:"+domain+".",
        "s3-prefix:"+bucket,
        "s3-bucket:"+bucket+"-"+environment+"-infra",
        "vpc-s3-key:"+environment+"/"+environment+"-"+region+"-vpc-terraform.tfstate",
        "dns-s3-key:"+environment+"/"+environment+"-"+region+"-route53-terraform.tfstate",
      }
      wf(spec,path)

      //CREATE PUBLIC NACL
      path = "AWS/"+region+"/"+environment+"/networking/public-nacl-config"
      spec = []string {
        "num-rules:9",
        "rule0:--ingress --rule-number 100 --protocol tcp --port-range From=22,To=22 --cidr-block 0.0.0.0/0  --rule-action allow",
        "rule1:--ingress --rule-number 110 --protocol tcp --port-range From=80,To=80 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule2:--ingress --rule-number 120 --protocol tcp --port-range From=443,To=443 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule3:--ingress --rule-number 130 --protocol tcp --port-range From=1024,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule4:--ingress --rule-number 140 --protocol udp --port-range From=1024,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule5:--ingress --rule-number 150 --protocol icmp --icmp-type-code Code=-1,Type=-1 --port-range From=-1,To=-1 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule6:--egress --rule-number 160 --protocol tcp --port-range From=0,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule7:--egress --rule-number 170 --protocol udp --port-range From=0,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule8:--egress --rule-number 180 --protocol icmp --icmp-type-code Code=-1,Type=-1 --port-range From=-1,To=-1 --cidr-block 0.0.0.0/0 --rule-action allow",
      }
      wf(spec,path)

      //CREATE PUBLIC SUBNETS
      path = "AWS/"+region+"/"+environment+"/networking/public-subnets-config"
      src := usr.HomeDir + "/.onecloud/states/"+region+"/"+environment+"/networking/public-subnets-config"
      pub_spec,err := readLines(src)
      if err != nil {
        log.Fatal(err)
      } else {
        wf(pub_spec,path)
      }

      //CREATE PUBLIC SECURITY GROUPS
      path = "AWS/"+region+"/"+environment+"/networking/public-security-group-config"
      spec = []string {
        "num-rules:4",
        "rule0:--protocol tcp --port 22 --cidr 0.0.0.0/0",
        "rule1:--protocol tcp --port 80 --cidr 0.0.0.0/0",
        "rule2:--protocol tcp --port 443 --cidr 0.0.0.0/0",
        "rule3:--protocol icmp --port -1 --cidr 0.0.0.0/0",
      }
      wf(spec,path)

      //CREATE PRIVATE NACL
      path = "AWS/"+region+"/"+environment+"/networking/private-nacl-config"
      spec = []string {
        "num-rules:9",
        "rule0:--ingress --rule-number 100 --protocol tcp --port-range From=22,To=22 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule1:--ingress --rule-number 110 --protocol tcp --port-range From=80,To=80 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule2:--ingress --rule-number 120 --protocol tcp --port-range From=443,To=443 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule3:--ingress --rule-number 130 --protocol tcp --port-range From=1024,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule4:--ingress --rule-number 140 --protocol udp --port-range From=1024,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule5:--ingress --rule-number 150 --protocol icmp --icmp-type-code Code=-1,Type=-1 --port-range From=-1,To=-1 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule6:--egress --rule-number 160 --protocol tcp --port-range From=0,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule7:--egress --rule-number 170 --protocol udp --port-range From=0,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule8:--egress --rule-number 180 --protocol icmp --icmp-type-code Code=-1,Type=-1 --port-range From=-1,To=-1 --cidr-block 0.0.0.0/0 --rule-action allow",
      }
      wf(spec,path)

      //CREATE PRIVATE SUBNETS
      path = "AWS/"+region+"/"+environment+"/networking/private-subnets-config"
      src = usr.HomeDir + "/.onecloud/states/"+region+"/"+environment+"/networking/private-subnets-config"
      pri_spec,err := readLines(src)
      if err != nil {
        log.Fatal(err)
      } else {
        wf(pri_spec,path)
      }

      //CREATE PRIVATE SECURITY GROUPS
      path = "AWS/"+region+"/"+environment+"/networking/private-security-group-config"
      spec = []string {
        "num-rules:4",
        "rule0:--protocol tcp --port 22 --cidr 0.0.0.0/0",
        "rule1:--protocol tcp --port 80 --cidr 0.0.0.0/0",
        "rule2:--protocol tcp --port 443 --cidr 0.0.0.0/0",
        "rule3:--protocol icmp --port -1 --cidr 0.0.0.0/0",
      }
      wf(spec,path)

      //CREATE BASTION NACL
      path = "AWS/"+region+"/"+environment+"/networking/bastion-nacl-config"
      spec = []string {
        "num-rules:11",
        "rule0:--ingress --rule-number 100 --protocol tcp --port-range From=22,To=22 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule1:--ingress --rule-number 110 --protocol tcp --port-range From=1024,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule2:--ingress --rule-number 120 --protocol udp --port-range From=1024,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule3:--ingress --rule-number 130 --protocol icmp --icmp-type-code Code=-1,Type=-1 --port-range From=-1,To=-1 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule4:--ingress --rule-number 140 --protocol tcp --port-range From=80,To=80 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule5:--ingress --rule-number 150 --protocol tcp --port-range From=443,To=443 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule6:--ingress --rule-number 160 --protocol tcp --port-range From=943,To=943 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule7:--ingress --rule-number 170 --protocol udp --port-range From=1194,To=1194 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule8:--egress --rule-number 180 --protocol tcp --port-range From=0,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule9:--egress --rule-number 190 --protocol udp --port-range From=0,To=65535 --cidr-block 0.0.0.0/0 --rule-action allow",
        "rule10:--egress --rule-number 200 --protocol icmp --icmp-type-code Code=-1,Type=-1 --port-range From=-1,To=-1 --cidr-block 0.0.0.0/0 --rule-action allow",
      }
      wf(spec,path)

      //CREATE BASTION SECURITY GROUPS
      path = "AWS/"+region+"/"+environment+"/networking/bastion-security-group-config"
      spec = []string {
        "num-rules:7",
        "rule0:--protocol tcp --port 22 --cidr 0.0.0.0/0",
        "rule1:--protocol tcp --port 80 --cidr 0.0.0.0/0",
        "rule2:--protocol udp --port 53 --cidr 0.0.0.0/0",
        "rule3:--protocol tcp --port 443 --cidr 0.0.0.0/0",
        "rule4:--protocol tcp --port 943 --cidr 0.0.0.0/0",
        "rule5:--protocol udp --port 1194 --cidr 0.0.0.0/0",
        "rule6:--protocol icmp --port -1 --cidr 0.0.0.0/0",
      }
      wf(spec,path)

      //CREATE VPC CHECKLIST FILE
      path = "AWS/"+region+"/"+environment+"/networking/vpc-checklist"
      spec = []string {
        "public-nacls:no",
        "public-subnets:no",
        "public-security-groups:no",
        "private-nacls:no",
        "private-subnets:no",
        "private-security-groups:no",
      }
      wf(spec,path)
      //UPDATE ENVIRONMENT'S STATE
      update_environment_state("current_state:created",environment,region)

      //TASKS COMPLETE
      fmt.Println("Created AWS VPC config for environment:" + environment + " with cidrblock:" + cidrblock + "/" +cidrslash+" in region:" + region)
    } else {
      fmt.Println("AWS VPC for " + environment + " in region:" + region + " has already been created.")
    }
  } else {
    fmt.Println("aws must be initialized first. run: onecloud build --provider aws --item cloud --action initialize")
  }
}
