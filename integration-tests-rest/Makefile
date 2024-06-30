test: prepare_test_environment
	docker-compose exec -T http go test ./...
	docker-compose down

prepare_test_environment: docker_build_test
	docker-compose down
	docker-compose up -d

docker_build_test:
	docker build . -t service_test --target=test

e2e_test:
	go test `go list ./... | grep e2e_test`

.PHONY: test e2e_test
