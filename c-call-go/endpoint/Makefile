go_procedure.a go_procedure.h: procedure.go
	go build -buildmode=c-archive -o go_procedure.a

.PHONY: clean
clean:
	rm *.h *.a
	