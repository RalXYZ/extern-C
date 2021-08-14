main: main.c c/c_procedure.* go/go.*
	cd c && make
	cd go && make
	$(CC) -o main main.c go/go.a c/libc_procedure.a -lpthread

.PHONY: clean
clean:
	cd c && make clean
	cd go && make clean
	rm main
