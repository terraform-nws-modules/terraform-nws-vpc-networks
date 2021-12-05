# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------
variable "vpc_cidr" {
  description = "Your VPC full CIDR"
  type        = string
}
# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------
variable "vpc_name" {
  description = "Your VPC name"
  type        = string
  default     = "vpc-mycompany"
}

variable "domain" {
  description = "Your VPC network domain"
  type        = string
  default     = "my.local"
}

# ---------------------------------------------------------------------------------------------------------------------
# Private subnets settings
# ---------------------------------------------------------------------------------------------------------------------
variable "subnet_private_name" {
  description = "Your private subnet name"
  type        = list(string)
  default     = null
}

variable "subnet_private_cidr" {
  description = "Your private subnets CIDR"
  type        = list(string)
  default     = null
}

# ---------------------------------------------------------------------------------------------------------------------
# Public subnets settings
# ---------------------------------------------------------------------------------------------------------------------
variable "subnet_public_name" {
  description = "Your public subnet name"
  type        = list(string)
  default     = null
}

variable "subnet_public_cidr" {
  description = "Your public subnets CIDR"
  type        = list(string)
  default     = null
}

variable "acl_name" {
  description = "Your ACL name"
  type        = string
  default     = "my-acl-public"
}

variable "acl_allowed_cidr_list" {
  description = "CIDR list for allowed ingress traffic. Allow all ingress traffic by default"
  type        = list(string)
  default     = ["0.0.0.0/0"]
}

variable "acl_allowed_port_list" {
  description = "Port list for allowed ingress traffic"
  type        = list(string)
  default     = ["80", "30000-30001"]
}
