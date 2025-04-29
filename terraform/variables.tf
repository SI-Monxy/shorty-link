variable "project" {
  type        = string
  description = "プロジェクト名の prefix"
}

variable "region" {
  type        = string
  description = "AWS リージョン"
}

variable "cidr_block" {
  type        = string
  description = "VPC の CIDR ブロック"
}

variable "availability_zones" {
  type        = list(string)
  description = "使用する AZ のリスト"
}

variable "public_subnet_cidrs" {
  type        = list(string)
  description = "Public サブネットの CIDR 一覧"
}
