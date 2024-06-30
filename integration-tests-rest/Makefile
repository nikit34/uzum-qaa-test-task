prepare_test_environment: docker_build_test
	docker-compose down -v
	docker-compose up -d

docker_build_test:
	docker build . -t service_test --target=test

e2e_test:
	go test `go list ./... | grep e2e_test`

unit_test:
	docker-compose exec -T http go test `go list ./... | grep -v e2e_test`

.PHONY: test prepare_test_environment e2e_test unit_test
