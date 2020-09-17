//export_aws_vpc_settings.go
package main

func export_aws_vpc_settings(region string,environment string) {
  spec := []string {
    "# GET PARAMETERS",

    "cidr_block=$( grep 'cidr-block' AWS/"+region+"/"+environment+"/networking/vpc-config | awk -F: '{print $2}' )",
    "cidr_slash=$( grep 'cidr-slash' AWS/"+region+"/"+environment+"/networking/vpc-config | awk -F: '{print $2}' )",
    "ec2_keypair=$( grep 'ec2-keypair' AWS/"+region+"/"+environment+"/networking/vpc-config | awk -F: '{print $2}' )",
    "primary_domain=$( grep 'primary-domain' AWS/"+region+"/"+environment+"/networking/vpc-config | awk -F: '{print $2}' )",
    "private_zone=$( grep 'private-zone' AWS/"+region+"/"+environment+"/networking/vpc-config | awk -F: '{print $2}' )",
    "backward_zone=$( grep 'backward-zone' AWS/"+region+"/"+environment+"/networking/vpc-config | awk -F: '{print $2}' )",
    "bucket=$( grep 's3-prefix' AWS/"+region+"/"+environment+"/networking/vpc-config | awk -F: '{print $2}' )",

    "sub0_publc=$( grep \"sub0-cidr\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",
    "sub0_pubsl=$( grep \"sub0-slash\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",

    "sub1_publc=$( grep \"sub1-cidr\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",
    "sub1_pubsl=$( grep \"sub1-slash\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",

    "sub2_publc=$( grep \"sub2-cidr\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",
    "sub2_pubsl=$( grep \"sub2-slash\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",

    "sub3_publc=$( grep \"sub3-cidr\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",
    "sub3_pubsl=$( grep \"sub3-slash\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",

    "sub4_publc=$( grep \"sub4-cidr\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",
    "sub4_pubsl=$( grep \"sub4-slash\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",

    "sub5_publc=$( grep \"sub5-cidr\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",
    "sub5_pubsl=$( grep \"sub5-slash\" AWS/"+region+"/"+environment+"/networking/public-subnets-config | awk -F: '{print $2}')",

    "sub0_privc=$( grep \"sub0-cidr\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",
    "sub0_prisl=$( grep \"sub0-slash\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",

    "sub1_privc=$( grep \"sub1-cidr\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",
    "sub1_prisl=$( grep \"sub1-slash\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",

    "sub2_privc=$( grep \"sub2-cidr\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",
    "sub2_prisl=$( grep \"sub2-slash\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",

    "sub3_privc=$( grep \"sub3-cidr\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",
    "sub3_prisl=$( grep \"sub3-slash\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",

    "sub4_privc=$( grep \"sub4-cidr\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",
    "sub4_prisl=$( grep \"sub4-slash\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",

    "sub5_privc=$( grep \"sub5-cidr\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",
    "sub5_prisl=$( grep \"sub5-slash\" AWS/"+region+"/"+environment+"/networking/private-subnets-config | awk -F: '{print $2}')",

    "sed -i \"\" \"s/OOOOO/$bucket/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_backend.tf",
    "sed -i \"\" \"s/OOOOO/$bucket/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/PPPPP/$primary_domain/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/QQQQQ/$backward_zone/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/RRRRR/$ec2_keypair/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/BBBBB/$cidr_block/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/ABABAB/$cidr_slash/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/CCCCC/$sub0_publc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/CDCDCD/$sub0_pubsl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/DDDDD/$sub1_publc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/DEDEDE/$sub1_pubsl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/EEEEE/$sub2_publc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/EFEFEF/$sub2_pubsl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/FFFFF/$sub3_publc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/FGFGFG/$sub3_pubsl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/GGGGG/$sub4_publc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/GHGHGH/$sub4_pubsl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/HHHHH/$sub5_publc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/HIHIHI/$sub5_pubsl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/IIIII/$sub0_privc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/IJIJIJ/$sub0_prisl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/JJJJJ/$sub1_privc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/JKJKJK/$sub1_prisl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/KKKKK/$sub2_privc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/KLKLKL/$sub2_prisl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/LLLLL/$sub3_privc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/LMLMLM/$sub3_prisl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/MMMMM/$sub4_privc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/MNMNMN/$sub4_prisl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",

    "sed -i \"\" \"s/NNNNN/$sub5_privc/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
    "sed -i \"\" \"s/NONONO/$sub5_prisl/g\" AWS/"+region+"/"+environment+"/iac-terraform/tf_vars.tf",
  }
  es(spec)
}
