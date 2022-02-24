resource "aws_s3_bucket" "web" {
  bucket = "depfootprint-static-assets"
}

resource "aws_s3_bucket_website_configuration" "web" {
  bucket = aws_s3_bucket.web.bucket

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "index.html"
  }
}

output "web_host" {
  value = aws_s3_bucket.web.website_endpoint
}