#--------------------------------------------------------------
# VPC Resources
#--------------------------------------------------------------

# Create a VPC for the region associated with the AZ
resource "aws_vpc" "vpc" {

  enable_dns_hostnames = "true"
  enable_dns_support   = "true"

  cidr_block = var.cidr_block

  tags = {
    Name = var.prefix
  }
}

#--------------------------------------------------------------
# Subnet Resources
#--------------------------------------------------------------
# Create a subnet for the AZ within the regional VPC
resource "aws_subnet" "public" {
  count = length(var.azs)

  vpc_id     = aws_vpc.vpc.id
  cidr_block = cidrsubnet(cidrsubnet(aws_vpc.vpc.cidr_block, 1, 0), 4 / 2, count.index)

  availability_zone = element(var.azs, count.index)

  tags = {
    "Name" = format(
      "%s%spublic%s%s",
      var.prefix,
      var.delimiter,
      var.delimiter,
      var.azs[count.index]
    )
  }
}

resource "aws_subnet" "private" {
  count = length(var.azs)

  vpc_id            = aws_vpc.vpc.id
  cidr_block        = cidrsubnet(cidrsubnet(aws_vpc.vpc.cidr_block, 1, 1), 4 / 2, count.index)
  availability_zone = element(var.azs, count.index)

  tags = {
    "Name" = format(
      "%s%sprivate%s%s",
      var.prefix,
      var.delimiter,
      var.delimiter,
      var.azs[count.index]
    )
  }
}

#--------------------------------------------------------------
# IGW Resources
#--------------------------------------------------------------
resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.vpc.id

  tags = {
    Name = var.account
  }
}



#--------------------------------------------------------------
# Nat Gateway Resources
#--------------------------------------------------------------
resource "aws_eip" "nat" {
  tags = {
    "Name" = format(
      "%s%seip%s%s",
      var.prefix,
      var.delimiter,
      var.delimiter,
      "ngw"
    )
  }
}

resource "aws_nat_gateway" "gw" {
  allocation_id = aws_eip.nat.id
  subnet_id     = concat(aws_subnet.public.*.id)[0]

  depends_on = [aws_internet_gateway.gw]

  tags = {
    "Name" = format(
      "%s%sngw",
      var.prefix,
      var.delimiter
    )
  }
}

#--------------------------------------------------------------
# Public Route Table Resources
#--------------------------------------------------------------
resource "aws_route_table" "public" {
  count = length(var.azs)

  vpc_id = aws_vpc.vpc.id

  tags = {
    "Name" = format(
      "%s%spublic%s%s%s%s",
      var.prefix,
      var.delimiter,
      var.delimiter,
      var.azs[count.index],
      var.delimiter,
      "rt"
    )
  }
}

resource "aws_route" "public_internet_gateway" {
  count                  = length(var.azs)
  route_table_id         = element(aws_route_table.public.*.id, count.index)
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.gw.id

  timeouts {
    create = "5m"
  }
}

#--------------------------------------------------------------
# Private Route Table Resources
#--------------------------------------------------------------
resource "aws_route_table" "private" {
  count = length(var.azs)

  vpc_id = aws_vpc.vpc.id

  tags = {
    "Name" = format(
      "%s%sprivate%s%s%s%s",
      var.prefix,
      var.delimiter,
      var.delimiter,
      var.azs[count.index],
      var.delimiter,
      "rt"
    )
  }
}

resource "aws_route" "private_nat_gateway" {
  count = length(var.azs)

  route_table_id         = element(aws_route_table.private.*.id, count.index)
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.gw.id

  timeouts {
    create = "5m"
  }
}

#--------------------------------------------------------------
# VPN Route Table Resources
#--------------------------------------------------------------

resource "aws_route_table_association" "public" {
  count          = length(var.azs)
  subnet_id      = element(aws_subnet.public.*.id, count.index)
  route_table_id = element(aws_route_table.public.*.id, count.index)
}

resource "aws_route_table_association" "private" {
  count          = length(var.azs)
  subnet_id      = element(aws_subnet.private.*.id, count.index)
  route_table_id = element(aws_route_table.private.*.id, count.index)
}

#--------------------------------------------------------------
# s3 vpc endpoint gateway
#--------------------------------------------------------------

resource "aws_vpc_endpoint" "s3" {
  vpc_id          = aws_vpc.vpc.id
  service_name    = "com.amazonaws.${var.region}.s3"
  route_table_ids = concat(aws_route_table.public.*.id, aws_route_table.private.*.id)

  tags = {
    Name = "com.amazonaws.${var.region}.s3"
  }
}

