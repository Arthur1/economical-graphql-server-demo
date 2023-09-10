data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "server" {
  name               = format("%s-lambda-role", local.name)
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy" "lambda_basic_execution" {
  arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "this" {
  role       = aws_iam_role.server.name
  policy_arn = data.aws_iam_policy.lambda_basic_execution.arn
}

resource "aws_lambda_function" "server" {
  function_name    = local.name
  runtime          = "provided.al2"
  handler          = "main"
  role             = aws_iam_role.server.arn
  architectures    = ["arm64"]
  memory_size      = 256
  timeout          = 7
  package_type     = "Zip"
  filename         = "../server-lambda.zip"
  source_code_hash = filebase64sha256("../server-lambda.zip")
}

resource "aws_lambda_permission" "api" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.server.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.api.execution_arn}/*/*"
}
