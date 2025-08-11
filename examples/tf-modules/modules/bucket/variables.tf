#--------------------------------------------------------------
# Shared Variables
#--------------------------------------------------------------
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

variable "bucket_name" {
  default     = "example"
  description = "Bucket name"
  type        = string
}
#--------------------------------------------------------------
# Provider definitions
#--------------------------------------------------------------
provider "aws" {
  region = var.region
}

#--------------------------------------------------------------
# Data Providers
#--------------------------------------------------------------

data "aws_canonical_user_id" "current" {}
