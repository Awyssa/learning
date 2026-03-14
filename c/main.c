#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const char *DEFAULT_FILE = "/index.html";

char *get_path_name(char *req) {
    char *start, *end;

    // Iterate over the req string until we get a space, and save the start as there (skip 'GET ' and save the start as the space)
    for (start = req; start[0] != ' '; start++) {
        if (!start[0]) return NULL;
    }

    // increment by 1 to skip the empty space
    start++;

    // get the end char of the path using the same logic
    for (end = start; end[0] != ' '; end++) {
        if (!end[0]) return NULL;
    }

    // calc the diff in memory addresses between the start and end of the path
    int diff = end - start;

    // allocate the length of the URL path memory to the heap
    char *path = malloc(diff + 1);

    // save the url path to the memory we just allocated
    for (int i = 0; i < diff; i++) {
        path[i] = start[i];
    }
    
    // add the null byte to identify where the string stops
    path[diff] = '\0';

    // return the start of the path string memory address
    return path;
}

char *append_html(char *path) {
    // Append 'index.html' so we know what file to get
    memcpy(
        path + strlen(path), 
        DEFAULT_FILE,
        strlen(DEFAULT_FILE) + 1
    );

    return 0;
}

int main() {
    char *req = "GET /homepage/about/profile HTTP/1.1";

    char *path = get_path_name(req);

    printf("This is the path: %s\n", path);

    append_html(path);

    printf("This is the path after appending: %s\n", path);

    free(path);

    return 0;
}