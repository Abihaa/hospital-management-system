
format:
	./scripts/format.sh

check: format
	./scripts/check.sh

test: check
	./scripts/test.sh
