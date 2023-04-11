provider "aws" {
  region = "us-west-2"
}

resource "aws_s3_bucket" "example" {
  bucket = "example-bucket-12345"
  tags = {
    example_tag = "required"
  }
}
