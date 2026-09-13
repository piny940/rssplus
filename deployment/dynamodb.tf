resource "aws_dynamodb_table" "feed_items" {
  name         = "rssplus_feed_items"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "feed_id"

  attribute {
    name = "feed_id"
    type = "S"
  }
}
