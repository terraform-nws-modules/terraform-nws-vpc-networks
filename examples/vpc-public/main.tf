terraform {
  required_version = ">= 1.0.10"

  required_providers {
    nws = {
      source  = "nws/nws"
      version = "0.1.2"
    }
  }
}

module "vpc_networks" {
  source = "../../src"

  vpc_name           = var.vpc_name
  vpc_cidr           = var.vpc_cidr
  domain             = var.domain
  subnet_public_name = var.subnet_public_name
  subnet_public_cidr = var.subnet_public_cidr
}
