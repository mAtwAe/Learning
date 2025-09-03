variable "instance_name"{
    description = "value of name tag for the ec2 instance"
    type = string
    default = "MyNewInstance"
}

variable "ec2_intance_type"{
    description = "AWS EC2 instance type"
    type = string
    default = "t3.micro"
}