resource "aws_lambda_function" "notify_once" {
  function_name = "rssplus_notify_once"
  role          = aws_iam_role.lambda_notify_once.arn
  package_type  = "Image"
  image_uri     = "${aws_ecrpublic_repository.rssplus.repository_uri}:latest"

  memory_size = 512
  timeout     = 60

  architectures = ["x86_64"]
}