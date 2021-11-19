.PHONY: fmt
fmt:
	find . -print | grep --regex '.*\.go' | xargs goimports -w -local
	find . -print | grep --regex '.*\.go' | xargs gofmt -s -l

.PHONY: gcs
gcs:
	docker compose up gcs
