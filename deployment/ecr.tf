resource "aws_ecr_repository" "rssplus" {
  name                 = "rssplus"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}
