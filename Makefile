.PHONY: fmt
fmt:
	find . -print | grep --regex '.*\.go' | xargs goimports -w -local
	find . -print | grep --regex '.*\.go' | xargs gofmt -s -l

.PHONY: sub
sub:
	docker compose up gcs
.PHONY: pub
pub:
	dockercompose exec app go run cmd/sub/main.go publish -m hai
