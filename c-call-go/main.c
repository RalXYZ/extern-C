#include <time.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "endpoint/go_procedure.h"

char caller_language[] = "C";

int main() {
    srand((unsigned int)time(NULL));
    int caller_stack = rand() % 10000;
    
    printf("[%04d] start of C main procedure\n", caller_stack);

    GoString caller_language_to_go = {
        caller_language, 
        strlen(caller_language)
    };
    go_procedure(caller_language_to_go, (GoInt)caller_stack);

    printf("[%04d] end of C main procedure\n", caller_stack);

    return 0;
}
