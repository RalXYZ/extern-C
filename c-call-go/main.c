#include <time.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "endpoint/go_procedure.h"

char language[] = "C";

int main() {
    srand((unsigned int)time(NULL));
    int stack = rand() % 10000;
    
    printf("[%04d] start of C main procedure\n", stack);

    GoString lang_to_go = {
        language, 
        strlen(language)
    };
    go_procedure(lang_to_go, (GoInt)stack);

    printf("[%04d] end of C main procedure\n", stack);

    return 0;
}
