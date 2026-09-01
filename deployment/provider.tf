terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.60.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "2.8"
    }
  }
}
provider "aws" {
  region = "ap-northeast-1"
}
