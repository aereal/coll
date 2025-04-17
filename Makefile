GO               := go
TEST_PACKAGE     := ./...
COVERAGE_PROFILE := ./coverage.out

.PHONY: coverage
coverage:
	$(GO) test \
		-covermode count \
		-coverprofile $(COVERAGE_PROFILE) \
		-v \
		$(TEST_PACKAGE)
