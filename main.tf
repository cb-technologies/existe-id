
# Configure and downloading plugins for aws provider

provider "aws" {
  region     = "${var.aws_region}"
}

# Retreiving VPC

data "aws_vpc" "default" {
  default = true
}
# Creating Security Group
resource "aws_security_group" "exist_security_group" {
  vpc_id      = "${data.aws_vpc.default.id}"
  name        = "exist"
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

# Create Postgres Cluster

resource "aws_rds_cluster" exist_cluster {
  cluster_identifier = "aurora-exist-cluster"
  engine = "aurora-postgresql"
  engine_version = "13.4"
  database_name = "existdb"
  master_username = "postgres"
  master_password = "exist_password"
  iam_database_authentication_enabled = true
  vpc_security_group_ids = [aws_security_group.exist_security_group.id]
}

# Postgres RDS instance
resource "aws_db_instance" "exist-postgres" {
  count                  = 1
  identifier             = "exist-identifier"
  cluster_identifer      =  aws_rds_cluster.exist_cluster.id
  name                   = "existdb"
  instance_class         = "db.t3.micro"
  allocated_storage      = 5
  engine                 = "postgres"
  engine_version         = aws_rds_cluster.exist_cluster.engine_version
  skip_final_snapshot    = true
  publicly_accessible    = true
  vpc_security_group_ids = [aws_security_group.exist_security_group.id]
  username               = "existuser"
}

//data "aws_caller_identity" "current_exist"{}
//
//variable "dbuser"{
//  type    = string
//  default = "dbuser"
//}