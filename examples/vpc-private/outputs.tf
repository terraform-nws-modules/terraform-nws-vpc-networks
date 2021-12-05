# ---------------------------------------------------------------------------------------------------------------------
# VPC outputs
# ---------------------------------------------------------------------------------------------------------------------
output "vpc_public_ip" {
  description = "VPC Public IP"
  value       = module.vpc_networks.vpc_public_ip
}

output "vpc_id" {
  description = "VPC internal ID"
  value       = module.vpc_networks.vpc_id
}

output "domain" {
  description = "Subnet domain"
  value       = module.vpc_networks.domain
}
# ---------------------------------------------------------------------------------------------------------------------
# Private subnets settings
# ---------------------------------------------------------------------------------------------------------------------
output "subnet_private_id" {
  description = "UUID of the new subnet"
  value       = module.vpc_networks.subnet_private_id
}
