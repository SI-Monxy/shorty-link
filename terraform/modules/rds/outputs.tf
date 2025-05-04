output "endpoint" {
  description = "RDSのエンドポイント"
  value       = aws_db_instance.this.endpoint
}

output "username" {
  description = "DBユーザー名"
  value       = aws_db_instance.this.username
}
