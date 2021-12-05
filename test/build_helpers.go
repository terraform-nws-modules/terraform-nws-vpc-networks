package test

func buildPrivateConfig(cfg testCaseT) map[string]interface{} {
	conf := map[string]interface{}{
		"vpc_name":            vpcName,
		"vpc_cidr":            vpcCidr,
		"domain":              domain,
		"subnet_private_name": cfg.subnetPrivName,
		"subnet_private_cidr": cfg.subnetPrivCidr,
	}
	return conf
}

func buildPublicConfig(cfg testCaseT) map[string]interface{} {
	conf := map[string]interface{}{
		"vpc_name":           vpcName,
		"vpc_cidr":           vpcCidr,
		"domain":             domain,
		"subnet_public_name": cfg.subnetPubName,
		"subnet_public_cidr": cfg.subnetPubCidr,
	}
	return conf
}

func buildFullConfig(cfg testCaseT) map[string]interface{} {
	conf := map[string]interface{}{
		"vpc_name":            vpcName,
		"vpc_cidr":            vpcCidr,
		"domain":              domain,
		"subnet_private_name": cfg.subnetPrivName,
		"subnet_private_cidr": cfg.subnetPrivCidr,
		"subnet_public_name":  cfg.subnetPubName,
		"subnet_public_cidr":  cfg.subnetPubCidr,
	}
	return conf
}
