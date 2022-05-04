DIR=./...
COVER_OUT=coverage.out

.PHONY: test vet fmt

default: test_all vet fmt

test_all: unit_test test_cover

unit_test:
	go test -v $(DIR)

test_cover:
	go test $(DIR) -coverprofile $(COVER_OUT)
	go tool cover -func $(COVER_OUT)
	rm $(COVER_OUT)

vet:
	@echo "go vet ."
	@go vet $(DIR) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w .