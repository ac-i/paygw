.PHONY: go-all
go-all: go-tidy go-fmt go-lint go-test go-bench

.PHONY: podman-scratch
podman-scratch: go-all podman-build podman-run curl-test podman-stop podman-prune

.PHONY: podman-separate-scratch
podman-separate-scratch: go-all podman-build podman-separate-run curl-test podman-separate-stop podman-prune

.PHONY: podman-all
podman-all: podman-run curl-test podman-stop

.PHONY: podman-separate-all
podman-separate-all: podman-separate-run curl-test podman-separate-stop

.PHONY: docker-scratch
docker-scratch: go-all docker-build docker-run curl-test docker-stop

.PHONY:docker-all
docker-all: docker-run curl-test docker-stop


.PHONY: go-tidy
go-tidy:
	@echo "*** go-tidy ***"
	@go mod tidy -go=1.17

.PHONY: go-fmt
go-fmt:
	@echo "*** go-fmt ***"
	@go fmt ./...

.PHONY: go-lint
go-lint:
	@echo "*** go-lint ***"
	# @$(HOME)/go/bin/golangci-lint run ./...

.PHONY: go-test
go-test:
	@echo "*** go-test ***"
	@go test paygw/test

.PHONY: go-bench
go-bench:
	@echo "*** go-bench ***"
	@go test -benchmem -bench . paygw/test

.PHONY: podman-build
podman-build:
	@echo "*** podman-build ***"
	@podman system df
	@podman build -t paygw .
	@podman image prune -f
	@podman volume prune -f
	@podman system df

.PHONY: podman-run
podman-run:
	@echo "*** podman-run ***"
	@until podman run --rm -d --name paygw -p 8080:8080 -p 8090:8090 paygw; do sleep 1; done
	@podman image prune -f
	@podman container ps

.PHONY: podman-separate-run
podman-separate-run:
	@echo "*** podman-separate-run ***"
	@until podman run --rm -d --name paygw-grpc -e PAYGW_SERVERS=grpc --network host paygw; do sleep 1; done
	@until podman run --rm -d --name paygw-restgw -e GRPC_ADDRESS=localhost:8090 -e PAYGW_SERVERS=restgw --network host paygw; do sleep 1; done
	@podman image prune -f
	@podman container ps

.PHONY: podman-stop
podman-stop:
	@echo "*** podman-stop ***"
	@podman container ps
	@podman stop paygw

.PHONY: podman-separate-stop
podman-separate-stop:
	@echo "*** podman-separate-stop ***"
	@podman container ps
	@podman stop paygw-restgw
	@podman stop paygw-grpc

.PHONY: podman-prune
podman-prune:
	@echo "*** podman-prune ***"
	# @podman system df
	@podman image prune -f
	# @podman system prune -f
	@podman volume prune -f
	@podman system df

.PHONY: curl-test
curl-test:
	@echo "*** curl-test ***"
	@echo " TODO "


.PHONY: docker-build
docker-build:
	@echo "*** docker-build ***"
	@docker build -t paygw .
	@docker image prune -f
	@docker volume prune -f
	@docker system df

.PHONY: docker-run
docker-run:
	@echo "*** docker-run ***"
	@until docker run --rm -d --name paygw -p 8080:8080 -p 8090:8090 paygw; do sleep 1; done
	@docker image prune -f
	@docker container ps

.PHONY: docker-stop
docker-stop:
	@echo "*** docker-stop ***"
	@docker container ps
	@docker stop paygw
