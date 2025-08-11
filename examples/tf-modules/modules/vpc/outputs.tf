#--------------------------------------------------------------
# VPC Outputs
#--------------------------------------------------------------
output "vpc_id" {
  value       = aws_vpc.vpc.id
  description = "VPC ID"
}

output "azs" {
  value       = var.azs
  description = "List of Availability Zones used - Useful for Local Zone Support"
}

output "cidr_block" {
  value       = aws_vpc.vpc.cidr_block
  description = "Cidr block used for VPC"
}

#--------------------------------------------------------------
# Subnet Outputs
#--------------------------------------------------------------
output "private_subnet_ids" {
  value       = concat(aws_subnet.private[*].id)
  description = "List of Private subnet ids"
}


output "public_subnet_ids" {
  value       = concat(aws_subnet.public[*].id)
  description = "List of Public subnet ids"
}

#--------------------------------------------------------------
# Route Outputs
#--------------------------------------------------------------
output "public_route_tables" {
  value       = concat(aws_route_table.public[*].id)
  description = "List of Public route tables"
}

output "private_route_tables" {
  value       = concat(aws_route_table.private[*].id)
  description = "List of Private route tables"
}


