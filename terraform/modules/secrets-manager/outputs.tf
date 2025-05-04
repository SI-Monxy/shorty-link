output "secret_arn" {
  description = "Secrets Manager Secret ARN"
  value       = aws_secretsmanager_secret.db_password.arn
}

output "secret_string" {
  description = "Secrets Managerに保存されたランダムパスワード（極力outputsでは使わない推奨）"
  sensitive   = true
  value       = random_password.db_password.result
}
