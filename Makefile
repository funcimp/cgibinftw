run:
	docker-compose build && \
	docker-compose up \
	--exit-code-from cgibinftw

test:
	TEST_MODE=true make run