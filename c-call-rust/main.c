#include <time.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char caller_language[] = "C";

extern int rust_procedure(const char*, int);

int main() {
    srand((unsigned int)time(NULL));
    int caller_stack = rand() % 10000;
    
    printf("[%04d] start of C main procedure\n", caller_stack);

    rust_procedure(caller_language, caller_stack);

    printf("[%04d] end of C main procedure\n", caller_stack);

    return 0;
}
