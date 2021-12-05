# ---------------------------------------------------------------------------------------------------------------------
# VPC settings
# ---------------------------------------------------------------------------------------------------------------------
variable "vpc_cidr" {
  description = "Your VPC full CIDR"
  type        = string
}

variable "vpc_name" {
  description = "Your VPC name"
  type        = string
  default     = "vpc-mycorp"
}

variable "domain" {
  description = "Your VPC network domain"
  type        = string
  default     = "test.local"
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
  description = "Your public subnets name"
  type        = list(string)
  default     = null
}

variable "subnet_public_cidr" {
  description = "Your public subnets CIDR"
  type        = list(string)
  default     = null
}
