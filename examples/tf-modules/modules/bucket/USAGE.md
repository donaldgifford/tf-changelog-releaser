## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_s3_bucket.example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket) | resource |
| [aws_s3_bucket.example_logs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket) | resource |
| [aws_s3_bucket_acl.example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_acl) | resource |
| [aws_s3_bucket_acl.example_logs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_acl) | resource |
| [aws_s3_bucket_logging.example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_logging) | resource |
| [aws_s3_bucket_metric.example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_metric) | resource |
| [aws_s3_bucket_ownership_controls.example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_ownership_controls) | resource |
| [aws_s3_bucket_ownership_controls.example_logs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_ownership_controls) | resource |
| [aws_s3_bucket_server_side_encryption_configuration.example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_server_side_encryption_configuration) | resource |
| [aws_s3_bucket_server_side_encryption_configuration.example_logs](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_server_side_encryption_configuration) | resource |
| [aws_canonical_user_id.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/canonical_user_id) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_bucket_name"></a> [bucket\_name](#input\_bucket\_name) | Bucket name | `string` | `"example"` | no |
| <a name="input_delimiter"></a> [delimiter](#input\_delimiter) | Delimiter for tags | `string` | `"-"` | no |
| <a name="input_prefix"></a> [prefix](#input\_prefix) | Prefix variable for naming convention | `string` | n/a | yes |
| <a name="input_region"></a> [region](#input\_region) | AWS Region Used | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_example_bucket_arn"></a> [example\_bucket\_arn](#output\_example\_bucket\_arn) | Example bucket Arn |
| <a name="output_example_bucket_domain_name"></a> [example\_bucket\_domain\_name](#output\_example\_bucket\_domain\_name) | Example bucket s3 domain name |
| <a name="output_example_bucket_id"></a> [example\_bucket\_id](#output\_example\_bucket\_id) | Example bucket id |
| <a name="output_example_bucket_region"></a> [example\_bucket\_region](#output\_example\_bucket\_region) | Example bucket region |
| <a name="output_example_logs_bucket_arn"></a> [example\_logs\_bucket\_arn](#output\_example\_logs\_bucket\_arn) | Example logs bucket arn |
| <a name="output_example_logs_bucket_domain_name"></a> [example\_logs\_bucket\_domain\_name](#output\_example\_logs\_bucket\_domain\_name) | Example Logs bucket s3 domain name |
| <a name="output_example_logs_bucket_id"></a> [example\_logs\_bucket\_id](#output\_example\_logs\_bucket\_id) | Example logs bucket id |
| <a name="output_example_logs_bucket_region"></a> [example\_logs\_bucket\_region](#output\_example\_logs\_bucket\_region) | Example Logs bucket region |