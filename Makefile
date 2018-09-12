QUIZ=sample 2018W37

.PHONY: all test $(QUIZ)

all: $(QUIZ)

$(QUIZ):
	@ln -sf ../Makefile $@/Makefile
	$(MAKE) -C $@ test

test:
	@go test

.PHONY: clean

clean:
