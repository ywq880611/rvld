TESTS := $(wildcard tests/*.sh)

build:
	go build

clean:
	go clean
	rm -rf out
	@printf '\e[32mClean all outputs\e[0m\n'

test: build
	$(MAKE) $(TESTS)
	@printf '\e[32mPassed all tests\e[0m\n'

$(TESTS):
	@echo 'Testing' $@
	@./$@
	@printf '\e[32mOK\e[0m\n'

.PHONY: build clean test $(TESTS)