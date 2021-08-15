main: main.c endpoint/procedure.rs endpoint/Cargo.toml
	cd endpoint && cargo build
	$(CC) -o $@ $< endpoint/target/debug/librust_procedure.a -lpthread -ldl

.PHONY: clean
clean:
	cd endpoint && cargo clean
	rm main
