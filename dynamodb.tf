resource "aws_dynamodb_table" "basic-dynamodb-table" {
  name           = "Messages"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "Name"
  range_key      = "Content"

  attribute {
    name = "Name"
    type = "S"
  }

  attribute {
    name = "Content"
    type = "S"
  }

  ttl {
    attribute_name = "TimeToExist"
    enabled        = false
  }

  tags = {
    Name        = "messages"
    Environment = "development"
  }
}