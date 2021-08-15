main: main.c endpoint
	cd endpoint && make
	$(CC) -o $@ $< endpoint/go_procedure.a -lpthread

.PHONY: clean
clean:
	cd endpoint && make clean
	rm main
