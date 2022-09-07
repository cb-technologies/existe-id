
# Configure and downloading plugins for aws provider

provider "aws" {
  region     = "${var.aws_region}"
}

# Creating VPC
resource "aws_default_vpc" "exist_vpc" {
//  cidr_block       = "${var.vpc_cidr}"
//  instance_tenancy = "default"

  tags = {
    Name = "Exist VPC"
  }
}

# Creating Security Group
resource "aws_security_group" "existvpc" {
//  vpc_id      = "${data.aws_vpc.default.id}"
  vpc_id        = aws_default_vpc.exist_vpc.id
  name        = "exist_vpc"
  description = "Allow all inbound for Postgres"
ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "random_string" "exist-db-password" {
  length  = 32
  upper   = true
  number  = true
  special = false
}


# Postgres RDS instance
resource "aws_db_instance" "exist-postgres" {
  identifier             = "exist-identifier"
  name                   = "existdb"
  instance_class         = "db.t3.micro"
  allocated_storage      = 5
  engine                 = "postgres"
  engine_version         = "13.4"
  skip_final_snapshot    = true
  publicly_accessible    = true
  vpc_security_group_ids = [aws_security_group.existvpc.id]
  username               = "existuser"
  password               = "random_string.exist-db-password.result}"
}