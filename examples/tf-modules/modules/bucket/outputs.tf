#--------------------------------------------------------------
# Output definitions
#--------------------------------------------------------------

# Example Bucket Outputs
output "example_bucket_id" {
  value       = aws_s3_bucket.example.id
  description = "Example bucket id"
}
output "example_bucket_arn" {
  value       = aws_s3_bucket.example.arn
  description = "Example bucket Arn"
}
output "example_bucket_domain_name" {
  value       = aws_s3_bucket.example.bucket_domain_name
  description = "Example bucket s3 domain name"
}
output "example_bucket_region" {
  value       = aws_s3_bucket.example.region
  description = "Example bucket region"
}
# Example bucket logs outputs
output "example_logs_bucket_id" {
  value       = aws_s3_bucket.example_logs.id
  description = "Example logs bucket id"
}

output "example_logs_bucket_arn" {
  value       = aws_s3_bucket.example_logs.arn
  description = "Example logs bucket arn"
}
output "example_logs_bucket_domain_name" {
  value       = aws_s3_bucket.example_logs.bucket_domain_name
  description = "Example Logs bucket s3 domain name"
}
output "example_logs_bucket_region" {
  value       = aws_s3_bucket.example_logs.region
  description = "Example Logs bucket region"
}
