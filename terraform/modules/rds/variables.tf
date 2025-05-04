variable "project" {
  description = "プロジェクト名"
  type        = string
}

variable "subnet_ids" {
  description = "RDS用のサブネットリスト"
  type        = list(string)
}

variable "security_group_ids" {
  description = "RDSにアタッチするセキュリティグループIDリスト"
  type        = list(string)
}

variable "db_password_secret_arn" {
  type        = string
  description = "Secrets ManagerのDBパスワードシークレットのARN"
}

variable "secrets_manager_dependency" {
  description = "Secrets Managerリソースに対する依存関係"
  type        = any
}
