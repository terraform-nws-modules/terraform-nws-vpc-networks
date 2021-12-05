COMMON?=common.go build_helpers.go valid_helpers.go
TIMEOUT=3m

show:
	@echo 'make lint - run linter'
	@echo 'make priv - run test for private subnets only'
	@echo 'make pub - run test for public subnets only'

lint:
	@golint *.go

priv:
	@go test -v -timeout $(TIMEOUT) $(COMMON) nws_vpc_private_network_test.go

pub:
	@go test -v -timeout $(TIMEOUT) $(COMMON) nws_vpc_public_network_test.go
