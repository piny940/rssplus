data "archive_file" "notify_once" {
  type        = "zip"
  source_file = "${path.module}/build/bootstrap"
  output_path = "${path.module}/build/notify_once.zip"
}

resource "aws_lambda_function" "notify_once" {
  function_name = "rssplus_notify_once"
  role          = aws_iam_role.lambda_notify_once.arn

  runtime          = "provided.al2023"
  handler          = "bootstrap"
  filename         = data.archive_file.notify_once.output_path
  source_code_hash = data.archive_file.notify_once.output_base64sha256

  memory_size = 512
  timeout     = 60

  architectures = ["x86_64"]
}
resource "aws_scheduler_schedule" "notify_once" {
  name       = "rssplus_notify_once"
  group_name = "default"

  flexible_time_window {
    mode = "FLEXIBLE"
    maximum_window_in_minutes = 5
  }

  schedule_expression = "rate(1 hours)"

  target {
    arn      = aws_lambda_function.notify_once.arn
    role_arn = aws_iam_role.event_bridge_notify_once.arn
  }
}
