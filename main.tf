provider "aws" {
  profile = "terraform"
  region  = "ap-northeast-1"
}

resource "aws_instance" "hellow-world" {
  ami           = "ami-0bba69335379e17f8"
  instance_type = "t2.micro"

  tags = {
    Name = "HellowWorld"
  }

  user_data = <<EOF
!"/bin/bash
amazon-linux-extras install nginx1.12 -y
systemctl start nginx
EOF
}
