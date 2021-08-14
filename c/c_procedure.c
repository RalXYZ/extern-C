#include <time.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "c_procedure.h"

#define LEN 50

void c_procedure(char* src_lang) {
    srand((unsigned int)time(NULL) + 1);
    int stack = rand() % 10000;
    printf("[%04d] start of a C procedure, called by %s\n", stack, src_lang);

    printf("[%04d] end of a C procedure\n", stack);
}
