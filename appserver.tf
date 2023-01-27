# -----------------------------------
# key pair
# ---------------------------------
resource "aws_key_pair" "keypair" {
  key_name   = "${var.project}-${var.environment}-keypair"
  public_key = file("./src/tastylog-dev-keypair.pub")

  tags = {
    Name    = "${var.project}-${var.environment}-keypair"
    Env     = var.environment
    Project = var.project
  }
}

# -----------------------------------
# EC2 instance
# ---------------------------------
resource "aws_instance" "app_server" {
  ami                         = data.aws_ami.app.id
  instance_type               = "t2.micro"
  subnet_id                   = aws_subnet.public_subnet_1a.id
  associate_public_ip_address = true
  vpc_security_group_ids = [
    aws_security_group.app_sg.id,
    aws_security_group.opmng_sg.id,
  ]
  key_name = aws_key_pair.keypair.key_name
  tags = {
    Name    = "${var.project}-${var.environment}-app-server"
    Env     = var.environment
    Project = var.project
    Type    = "app"
  }
}

