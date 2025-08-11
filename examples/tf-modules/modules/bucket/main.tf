#--------------------------------------------------------------
# Bucket and bucket logs bucket
#--------------------------------------------------------------

resource "aws_s3_bucket" "example" {
  bucket = "${var.prefix}${var.delimiter}${var.bucket_name}"

  tags = {
    Name = "${var.prefix}${var.delimiter}${var.bucket_name}"
  }
}

resource "aws_s3_bucket" "example_logs" {
  bucket = "${var.prefix}${var.delimiter}${var.bucket_name}${var.delimiter}logs"

  tags = {
    Name = "${var.prefix}${var.delimiter}${var.bucket_name}${var.delimiter}logs"
  }
}

#--------------------------------------------------------------
# sse configuration
#--------------------------------------------------------------

resource "aws_s3_bucket_server_side_encryption_configuration" "example" {
  bucket = aws_s3_bucket.example.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}


resource "aws_s3_bucket_server_side_encryption_configuration" "example_logs" {
  bucket = aws_s3_bucket.example_logs.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

#--------------------------------------------------------------
# ownership controls and ACL
#--------------------------------------------------------------


resource "aws_s3_bucket_ownership_controls" "example" {
  bucket = aws_s3_bucket.example.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_ownership_controls" "example_logs" {
  bucket = aws_s3_bucket.example_logs.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}


resource "aws_s3_bucket_acl" "example_logs" {
  depends_on = [aws_s3_bucket_ownership_controls.example_logs]
  bucket     = aws_s3_bucket.example_logs.id
  acl        = "log-delivery-write"
}

resource "aws_s3_bucket_acl" "example" {
  depends_on = [aws_s3_bucket_ownership_controls.example]
  bucket     = aws_s3_bucket.example.id
  access_control_policy {
    grant {
      grantee {
        id   = data.aws_canonical_user_id.current.id
        type = "CanonicalUser"
      }
      permission = "FULL_CONTROL"
    }

    grant {
      grantee {
        type = "Group"
        uri  = "http://acs.amazonaws.com/groups/s3/LogDelivery"
      }
      permission = "READ_ACP"
    }

    owner {
      id = data.aws_canonical_user_id.current.id
    }
  }
}

#--------------------------------------------------------------
# logging configuration
#--------------------------------------------------------------

resource "aws_s3_bucket_logging" "example" {
  bucket        = aws_s3_bucket.example.id
  target_bucket = aws_s3_bucket.example_logs.id
  target_prefix = "log/"
}

#--------------------------------------------------------------
#  bucket metrics
#--------------------------------------------------------------

resource "aws_s3_bucket_metric" "example" {
  bucket = aws_s3_bucket.example.id
  name   = "EntireBucket"
}
