provider "aws" {
    profile = "default"
    region = "ap-southeast-5"
}

resource "aws_instance" "app_server"{
    ami = ""
    instance_type = var.ec2_intance_type

    tags = {
        Name = var.instance_name
    }
}