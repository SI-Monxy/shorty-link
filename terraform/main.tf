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
