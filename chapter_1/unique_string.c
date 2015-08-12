//
//  unique_string.c
//  
//
//  Created by Pawel Cisek on 12/08/15.
//
//

#include "unique_string.h"


int main(int argc, char *argv[]){
    char* str = argv[1];
    bool uniq = unique_string(str);
    fputs( uniq ? "true \n" : "false \n", stdout);
}

bool unique_string(char* str){
    int array_size = (int)strlen(str);
    for (int i = 0; i < array_size; i++) {
        char test_char = str[i];
        for (int j = 0; j < array_size; j++) {
            if (test_char == str[j] && i != j){
                return false;
            }
        }
    }
    return true;
}