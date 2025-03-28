GOPATH=${HOME}/go

%:
	@true

PHONY: run
run:
	./scripts/run.sh $(filter-out $@,$(MAKECMDGOALS))
.PHONY: up
up:
	./scripts/container-up-local.sh
down:
	./scripts/container-down-local.sh