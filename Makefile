
format:
	${DOCKRUN} bash ./scripts/format.sh

check: format
	${DOCKRUN} bash ./scripts/check.sh

test: check
	${DOCKRUN} bash ./scripts/test.sh
