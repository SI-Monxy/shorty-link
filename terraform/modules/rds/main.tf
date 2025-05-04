data "aws_secretsmanager_secret_version" "db_password" {
  secret_id = var.db_password_secret_arn

  depends_on = [var.secrets_manager_dependency]
}

resource "aws_db_subnet_group" "this" {
  name       = "${var.project}-db-subnet-group"
  subnet_ids = var.subnet_ids

  tags = {
    Name = "${var.project}-db-subnet-group"
  }
}

resource "aws_db_instance" "this" {
  identifier              = "${var.project}-db"
  engine                  = "mysql"
  engine_version          = "8.0"
  instance_class          = "db.t3.micro"  # 無料枠＆コスパ重視
  allocated_storage       = 20
  storage_type            = "gp2"
  username                = "admin"
  password                = data.aws_secretsmanager_secret_version.db_password.secret_string
  db_name                 = "shorty"
  db_subnet_group_name    = aws_db_subnet_group.this.name
  vpc_security_group_ids  = var.security_group_ids
  publicly_accessible     = false
  skip_final_snapshot     = true

  tags = {
    Name = "${var.project}-db"
  }
}
