variable "project" {
  description = "プロジェクト名の prefix"
  type        = string
}

variable "cluster_name" {
  description = "ECSクラスタ名"
  type        = string
}

variable "container_image" {
  description = "ECRのコンテナイメージ"
  type        = string
}

variable "container_port" {
  description = "コンテナがリッスンするポート番号"
  type        = number
}

variable "subnets" {
  description = "起動に使うサブネットのリスト"
  type        = list(string)
}

variable "security_group_ids" {
  description = "ECSタスクに適用するセキュリティグループIDのリスト"
  type        = list(string)
}

variable "target_group_arn" {
  description = "ALBのターゲットグループARN"
  type        = string
}

variable "db_host" {
  description = "接続するDBホスト"
  type        = string
}

variable "db_port" {
  description = "接続するDBポート"
  type        = number
}

variable "db_user" {
  description = "接続するDBユーザー"
  type        = string
}

variable "db_password_secret_arn" {
  description = "DBパスワードのSecrets Manager ARN"
  type        = string
}

variable "region" {
  description = "AWSリージョン"
  type        = string
}
