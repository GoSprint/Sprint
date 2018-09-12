QUIZ=sample

.PHONY: all test $(QUIZ)

all: $(QUIZ)

$(QUIZ):
	@ln -sf ../Makefile $@/Makefile
	$(MAKE) -C $@ test

test:
	@go test

.PHONY: clean

clean:
