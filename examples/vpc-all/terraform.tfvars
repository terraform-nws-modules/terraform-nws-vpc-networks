vpc_name            = "vpc-test"
vpc_cidr            = "10.0.1.0/24"
domain              = "test.local"
subnet_private_name = ["net0-private", "net1-private"]
subnet_private_cidr = ["10.0.1.0/30", "10.0.1.10/30"]
subnet_public_name  = ["net0-public", "net1-public"]
subnet_public_cidr  = ["10.0.1.30/30", "10.0.1.30/30"]
