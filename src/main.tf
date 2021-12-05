terraform {
  required_version = ">= 1.0.10"

  required_providers {
    nws = {
      source  = "nws/nws"
      version = "0.1.2"
    }
  }
}

module "vpc" {
  source = "github.com/terraform-nws-modules/terraform-nws-vpc/src"

  name   = var.vpc_name
  cidr   = var.vpc_cidr
  domain = var.domain
}

// Private subnets are accessible only thru VPN
module "subnets_private" {
  source = "github.com/terraform-nws-modules/terraform-nws-subnet/src"

  name   = var.subnet_private_name
  cidr   = var.subnet_private_cidr
  domain = var.domain
  vpc_id = module.vpc.vpc_id
  public = "false"
}

// Public subnets may have ACLs to allow in/out traffic from specified ports
module "subnets_public" {
  source = "github.com/terraform-nws-modules/terraform-nws-subnet/src"

  name   = var.subnet_public_name
  cidr   = var.subnet_public_cidr
  domain = var.domain
  vpc_id = module.vpc.vpc_id
  public = "true"

  // ACL setup
  acl_name              = var.acl_name
  acl_allowed_cidr_list = var.acl_allowed_cidr_list
  acl_allowed_port_list = var.acl_allowed_port_list

}
