terraform {
  required_version = ">= 1.4.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "local" {
    path = "terraform.tfstate"
  }
}

provider "aws" {
  region = var.region
}

resource "aws_security_group" "db" {
  name        = "${var.project}-db-sg"
  description = "Allow MySQL from ECS"
  vpc_id      = module.network.vpc_id

  ingress {
    from_port       = 3306
    to_port         = 3306
    protocol        = "tcp"
    security_groups = [module.security_group.ecs_sg_id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.project}-db-sg"
  }
}

module "network" {
  source = "./modules/network"

  vpc_cidr            = var.cidr_block
  azs                 = var.availability_zones
  public_subnet_cidrs = var.public_subnet_cidrs
  project             = var.project
}

module "security_group" {
  source  = "./modules/security-group"

  project = var.project
  vpc_id  = module.network.vpc_id
}

module "alb" {
  source = "./modules/alb"

  project            = var.project
  vpc_id             = module.network.vpc_id
  public_subnet_ids  = module.network.public_subnet_ids
  alb_sg_id          = module.security_group.alb_sg_id
}

module "ecr" {
  source  = "./modules/ecr"
  project = var.project
}

module "ecs" {
  source             = "./modules/ecs"
  project            = var.project
  cluster_name       = "${var.project}-cluster"
  container_image    = "918560296443.dkr.ecr.ap-northeast-1.amazonaws.com/shorty-link-ecr:latest"
  container_port     = 8080
  subnets            = module.network.public_subnet_ids
  security_group_ids = [module.security_group.ecs_sg_id]
  target_group_arn   = module.alb.target_group_arn
  region             = var.region

  db_host               = "shorty-link-db.cb6qc8uocdcu.ap-northeast-1.rds.amazonaws.com"
  db_port               = 3306
  db_user               = "admin"
  db_password_secret_arn = module.secrets_manager.secret_arn
}

module "secrets_manager" {
  source  = "./modules/secrets-manager"
  project = var.project
}

module "rds" {
  source                 = "./modules/rds"
  project                = var.project
  subnet_ids             = module.network.public_subnet_ids
  security_group_ids     = [aws_security_group.db.id]
  db_password_secret_arn = module.secrets_manager.secret_arn
  secrets_manager_dependency  = module.secrets_manager
}
