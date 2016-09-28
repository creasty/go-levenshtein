#include <stdlib.h>

int
levenshtein(char *a, int a_size, char *b, int b_size)
{
    if (a == b) {
        return 0;
    }
    if (a_size == 0) {
        return b_size;
    }
    if (b_size == 0) {
        return a_size;
    }

    unsigned int *m = calloc(a_size, sizeof(unsigned int));
    if (m == NULL) {
        return -1;
    }

    unsigned int ia = 0;
    unsigned int ib = 0;
    unsigned int da;
    unsigned int db;
    unsigned int d;

    while (ia < a_size) {
        m[ia] = ia + 1;
        ia++;
    }

    while (ib < b_size) {
        char c = b[ib];
        d = da = ib++;
        ia = -1;

        while (++ia < a_size) {
            db = (c == a[ia]) ? da : da + 1;
            da = m[ia];

            m[ia] = d = (da > d)
                ? (db > d ? d + 1 : db)
                : (db > da ? da + 1 : db);
        }
    }

    free(m);

    return d;
}
