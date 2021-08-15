.PHONY: build clean

build:
	@for i in `ls`; do \
		if [ -d "$$i" ]; then \
			cd $$i && make; \
			cd ..; \
		fi; \
	done;

clean:
	@for i in `ls`; do \
		if [ -d "$$i" ]; then \
			cd $$i && make clean; \
			cd ..; \
		fi; \
	done;
