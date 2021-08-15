main: main.go endpoint
	cd endpoint && make
	go build -o $@

.PHONY: clean
clean:
	rm main
	cd endpoint && make clean
