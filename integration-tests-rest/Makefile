test: docker_build_test
	docker-compose down
	docker-compose up -d
	docker-compose exec -T http go test ./...
	docker-compose down

docker_build_test:
	docker build . -t service_test --target=test

unit_test:
	go test `go list ./... | grep -v e2e_test`


e2e_test:
	go test `go list ./... | grep e2e_test`
