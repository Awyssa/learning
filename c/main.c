#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static const char DEFAULT_FILE[] = "index.html";

/* Extract the URL path from an HTTP request line.
Returns a heap-allocated string, or NULL on failure. */
char *get_path(const char *req) {
    // Skip the method (e.g. "GET ")
    const char *start = strchr(req, ' ');
    if (!start) return NULL;
    start++;
    // Find the end of the path
    const char *end = strchr(start, ' ');
    if (!end) return NULL;
    // Strip trailing slash
    while (end > start + 1 && end[-1] == '/') end--;
    size_t len = end - start;
    char *path = malloc(len + 1);
    if (!path) return NULL;
    memcpy(path, start, len);
    path[len] = '\0';
    return path;
}

/* Append "/index.html" to path.
Returns a new heap-allocated string, frees the original, or NULL on failure. */
char *append_index(char *path) {
    size_t path_len = strlen(path);
    size_t suffix_len = 1 + strlen(DEFAULT_FILE); // '/' + "index.html"
    char *result = malloc(path_len + suffix_len + 1);
    if (!result) {
        free(path);
        return NULL;
    }
    memcpy(result, path, path_len);
    result[path_len] = '/';
    memcpy(result + path_len + 1, DEFAULT_FILE, suffix_len); // includes '\0'
    free(path);
    return result;
}

int main(void) {
    const char *req = "GET /homepage/about/profile HTTP/1.1";
    char *path = get_path(req);
    if (!path) { 
        fprintf(stderr, "Failed to parse path\n"); return 1; 
    }
    path = append_index(path);
    if (!path) { 
        fprintf(stderr, "Out of memory\n"); return 1; 
    }

    int fd = open(path, 0_RDONLY);

    if(fd == -1) {
        // handle error
    }

    char buffer[100];

    size_t len = read(fd, buffer, 100);

    free(path);
    return 0;
}