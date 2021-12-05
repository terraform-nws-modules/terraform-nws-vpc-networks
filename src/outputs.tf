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
# Private subnets settings
# ---------------------------------------------------------------------------------------------------------------------
output "subnet_private_id" {
  description = "UUID of the new subnet"
  value       = module.subnets_private.subnet_id
}

# output "subnet_acl_id" {
#   description = "ACL id for this subnet"
#   value       = module.subnet_private.subnet_acl_id
# }

# output "subnet_acl_rule_id" {
#   description = "ACL rule id for this subnet"
#   value       = module.subnet_private.subnet_acl_rule_id
# }
