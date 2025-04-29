variable "project" {
  description = "プロジェクト名のプレフィックス"
  type        = string
}

variable "vpc_id" {
  description = "ALBを配置するVPC ID"
  type        = string
}

variable "public_subnet_ids" {
  description = "ALBを配置するパブリックサブネットのID一覧"
  type        = list(string)
}

variable "alb_sg_id" {
  description = "ALB用のセキュリティグループID"
  type        = string
}
