#include <time.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char language[] = "C";

extern int rust_procedure(const char*, int);

int main() {
    srand((unsigned int)time(NULL));
    int stack = rand() % 10000;
    
    printf("[%04d] start of C main procedure\n", stack);

    rust_procedure(language, stack);

    printf("[%04d] end of C main procedure\n", stack);

    return 0;
}
