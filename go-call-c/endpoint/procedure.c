#include <time.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "c_procedure.h"

#define LEN 50

void c_procedure(char* src_lang, int caller_stack) {
    srand((unsigned int)time(NULL) + 1);
    int stack = rand() % 10000;
    printf("[%04d] start of a C procedure, called by %s [%04d]\n", stack, src_lang, caller_stack);

    printf("[%04d] end of a C procedure\n", stack);
}
