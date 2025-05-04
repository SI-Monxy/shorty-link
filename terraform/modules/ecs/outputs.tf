output "ecs_cluster_id" {
  description = "ECSクラスタID"
  value       = aws_ecs_cluster.this.id
}

output "ecs_service_name" {
  description = "ECSサービス名"
  value       = aws_ecs_service.this.name
}

output "ecs_task_definition_arn" {
  description = "ECSタスク定義のARN"
  value       = aws_ecs_task_definition.this.arn
}