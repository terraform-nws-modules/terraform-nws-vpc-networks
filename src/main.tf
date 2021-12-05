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

module "subnets_private" {
  source = "github.com/terraform-nws-modules/terraform-nws-subnet/src"

  name   = var.subnet_private_name
  cidr   = var.subnet_private_cidr
  domain = var.domain
  vpc_id = module.vpc.vpc_id
  public = "false"
}

# module "subnets-public" {
#   source = "github.com/terraform-nws-modules/terraform-nws-subnet/src"

#   name   = var.subnet_name
#   cidr   = var.subnet_cidr
#   domain = var.domain
#   vpc_id = module.vpc.vpc_id
#   public = "true"
# }
