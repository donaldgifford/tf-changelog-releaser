#--------------------------------------------------------------
# Shared Variables
#--------------------------------------------------------------

variable "account" {
  default     = ""
  description = "Name of account for default tagging."
  type        = string
}

variable "region" {
  description = "AWS Region Used"
  type        = string
}

variable "prefix" {
  description = "Prefix variable for naming convention"
  type        = string
}

variable "delimiter" {
  default     = "-"
  description = "Delimiter for tags"
  type        = string
}

variable "cidr_block" {
  default     = "192.168.0.0/22"
  description = "CIDR Block - /22 Required"
  type        = string
}

variable "azs" {
  description = "A list of availability zones in the region"
  type        = list(string)
  default     = []
}

variable "public_subnets" {
  default     = ["192.168.0.0/24", "192.168.0.1/24", "192.168.0.2/24"]
  description = "public subnet inputs"
  type        = list(string)
}

#--------------------------------------------------------------
# Data Provider definitions
#--------------------------------------------------------------
data "aws_availability_zones" "available" {
  state = "available"
}

data "aws_caller_identity" "current" {}

#--------------------------------------------------------------
# Provider definitions
#--------------------------------------------------------------

provider "aws" {
  region = var.region
}

