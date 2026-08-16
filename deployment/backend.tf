resource "aws_s3_bucket" "tf_state" {
  bucket = "rssplus-tfstate.piny940.com"

  tags = {
    Name        = "RSSPlus tfstate"
  }
}
terraform {
  backend "s3" {
    bucket       = "rssplus-tfstate.piny940.com"
    key          = "terraform.tfstate"
    region = "ap-northeast-1"
    use_lockfile = true
  }
}
