
# Configure and downloading plugins for aws provider

provider "aws" {
  region     = "${var.aws_region}"
}

# Creating VPC
resource "aws_vpc" "exist_vpc" {
  cidr_block       = "${var.vpc_cidr}"
  instance_tenancy = "default"

  tags = {
    Name = "Exist VPC"
  }
}


# Creating Internet Gateway
resource "aws_internet_gateway" "existgateway" {
  vpc_id = "${aws_vpc.exist_vpc.id}"
}

# Grant the internet access to VPC by updating its main route table
resource "aws_route" "internet_access" {
  route_table_id         = "${aws_vpc.exist_vpc.main_route_table_id}"
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = "${aws_internet_gateway.existgateway.id}"
}

# Creating 1st subnet
resource "aws_subnet" "demosubnet" {
  vpc_id                  = "${aws_vpc.exist_vpc.id}"
  cidr_block             = "${var.subnet_cidr}"
  map_public_ip_on_launch = true
  availability_zone = "us-east-1a"

  tags = {
    Name = "Exist subnet"
  }
}


# Creating Security Group
resource "aws_security_group" "existvpc" {
  vpc_id      = "${data.aws_vpc.default.id}"
  name        = "existvpc"
  description = "Allow all inbound for Postgres"
ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
data "aws_vpc" "default" {
  default = true
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