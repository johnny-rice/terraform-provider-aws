resource "aws_datasync_location_azure_blob" "test" {
{{- template "region" }}
  agent_arns          = [aws_datasync_agent.test.arn]
  authentication_type = "SAS"
  container_url       = "https://myaccount.blob.core.windows.net/mycontainer"
  subdirectory        = "/myvdir1/myvdir2"

  sas_configuration {
    token = "sp=r&st=2023-12-20T14:54:52Z&se=2023-12-20T22:54:52Z&spr=https&sv=2021-06-08&sr=c&sig=aBBKDWQvyuVcTPH9EBp%%2FXTI9E%%2F%%2Fmq171%%2BZU178wcwqU%%3D"
  }
{{- template "tags" . }}
}

# testAccLocationAzureBlobConfig_base

resource "aws_datasync_agent" "test" {
{{- template "region" }}
  ip_address = aws_instance.test.public_ip
  name       = var.rName
}

# testAccAgentAgentConfig_base


# Reference: https://docs.aws.amazon.com/datasync/latest/userguide/deploy-agents.html
data "aws_ssm_parameter" "aws_service_datasync_ami" {
{{- template "region" }}
  name = "/aws/service/datasync/ami"
}

resource "aws_internet_gateway" "test" {
{{- template "region" }}
  vpc_id = aws_vpc.test.id
}

resource "aws_route_table" "test" {
{{- template "region" }}
  vpc_id = aws_vpc.test.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.test.id
  }
}

resource "aws_route_table_association" "test" {
{{- template "region" }}
  subnet_id      = aws_subnet.test[0].id
  route_table_id = aws_route_table.test.id
}

resource "aws_security_group" "test" {
{{- template "region" }}
  name   = var.rName
  vpc_id = aws_vpc.test.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "test" {
{{- template "region" }}
  depends_on = [aws_internet_gateway.test]

  ami                    = data.aws_ssm_parameter.aws_service_datasync_ami.value
  instance_type          = data.aws_ec2_instance_type_offering.available.instance_type
  vpc_security_group_ids = [aws_security_group.test.id]
  subnet_id              = aws_subnet.test[0].id

  # The Instance must have a public IP address because the aws_datasync_agent retrieves
  # the activation key by making an HTTP request to the instance
  associate_public_ip_address = true
}

{{ template "acctest.ConfigVPCWithSubnets" 1 }}

# acctest.AvailableEC2InstanceTypeForAvailabilityZone("aws_subnet.test[0].availability_zone", "m5.2xlarge", "m5.4xlarge")

data "aws_ec2_instance_type_offering" "available" {
{{- template "region" }}
  filter {
    name   = "instance-type"
    values = local.preferred_instance_types
  }

  filter {
    name   = "location"
    values = [aws_subnet.test[0].availability_zone]
  }

  location_type            = "availability-zone"
  preferred_instance_types = local.preferred_instance_types
}

locals {
  preferred_instance_types = ["m5.2xlarge", "m5.4xlarge"]
}
