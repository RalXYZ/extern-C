# `extern "C"`
This project demonstrates how to mix C with other programming languages. 
It aims to create a [Rosetta Stone](https://en.wikipedia.org/wiki/Rosetta_Stone) of the procedure shown in the following pseudocode:  

```go
// use programming language x

func main() {
    caller_language := "x-language";
    caller_stack := generate_random_number();
    println("[%04d] start of x-language main procedure", caller_stack);

    y_language_procedure(caller_language, caller_stack);

    println("[{%04d}] end of x-language main procedure", caller_stack);
}
```

```fortran
! use programming language y

FUNCTION y_language_procedure(caller_language, caller_stack)
    callee_stack = generate_random_number()
    PRINT "[%04d] start of a y-language procedure, called by [%04d], %s", callee_stack, caller_language, caller_stack

    PRINT "[%04d] start of a y-language procedure", callee_stack
END FUNCTION y_language_procedure
```

As shown above, the main function written in x language calls a function written in y language. The output of the whole procedure looks like this:  

```
[4199] start of x-language main procedure
[8549] start of a y-language procedure, called by x-language [4199]
[8549] end of a y-language procedure
[4199] end of x-language main procedure
```

The cross-language API, as mentioned above, needs to receive a integer and a string, in order to show their basic ability of mix programming.  

**All the callee procedures will be archived into static libraries using their own toolchain.** For example, to build a C static library, instead of using `cargo`, we use `gcc`.  

Currently implemented:  

| Caller | Callee |
| :----: | :----: |
| C | golang |
| C | rust |
| golang | C |
| rust | C |  

Will implement:  
- C++

## Usage
Fisrtly, change directory to one of the `x-call-y`. If you want to build all `x-call-y`s, just stay in the root directory.  

To build the project:  

```shell
make
```

To clean the build targets:

```shell
make clean
```