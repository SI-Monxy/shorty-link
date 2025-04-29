variable "project" {
  type        = string
  description = "プロジェクト名（prefix）"
}

variable "vpc_cidr" {
  type        = string
  description = "VPCのCIDRブロック"
}

variable "public_subnet_cidrs" {
  type        = list(string)
  description = "パブリックサブネットのCIDR一覧"
}

variable "azs" {
  type        = list(string)
  description = "使用するAZ（可用性ゾーン）"
}
