# ---------------------------------------------------------------------------------------------------------------------
# VPC outputs
# ---------------------------------------------------------------------------------------------------------------------
output "vpc_public_ip" {
  description = "VPC Public IP"
  value       = module.vpc.vpc_nat_ip
}

output "vpc_id" {
  description = "VPC internal ID"
  value       = module.vpc.vpc_id
}

output "domain" {
  description = "Subnet domain"
  value       = var.domain
}
# ---------------------------------------------------------------------------------------------------------------------
# Private subnets outputs
# ---------------------------------------------------------------------------------------------------------------------
output "subnet_private_id" {
  description = "UUID list of your private subnets"
  value       = module.subnets_private.subnet_id
}

# ---------------------------------------------------------------------------------------------------------------------
# Public subnets outputs
# ---------------------------------------------------------------------------------------------------------------------
output "subnet_public_id" {
  description = "UUID list of your public subnets"
  value       = module.subnets_public.subnet_id
}

output "subnet_public_acl_id" {
  description = "ACL ID list for your public subnets"
  value       = module.subnets_public.subnet_acl_id
}

output "subnet_public_acl_rule_id" {
  description = "ACL Rule ID list for your public subnets"
  value       = module.subnets_public.subnet_acl_rule_id
}
