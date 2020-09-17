package main


func check_aws_vpc_config(region string,environment string) int {

  //CHECK VPC CONFIG
  spec := []string {
    "name=$( grep 'name' AWS/"+region+"/"+environment+"/networking/vpc-config | awk -F: '{print $2}' )",
    "cidrblock=$( grep 'cidr-block' AWS/"+region+"/"+environment+"/networking/vpc-config | awk -F: '{print $2}' )",
    "ec2keypair=$( grep 'ec2-keypair' AWS/"+region+"/"+environment+"/networking/vpc-config | awk -F: '{print $2}' )",
    "",
    "if [[ \"$name\" ]] && [[ \"$cidrblock\" ]] && [[ \"$ec2keypair\" ]]",
    "then",
    "  echo 0",
    "else",
    "  echo 1",
    "fi",
  }
  retBuf := []rune(es(spec))
  retVal := string(retBuf[0])
  if retVal == "0" {
    return 0
  } else {
    return 1
  }
}
