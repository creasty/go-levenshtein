#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>

unsigned int
levenshtein(int32_t *a, size_t a_size, int32_t *b, size_t b_size)
{
    unsigned int *m = calloc(a_size, sizeof(unsigned int));
    if (m == NULL) {
        return -1;
    }

    unsigned int ia = 0;
    unsigned int ib = 0;
    unsigned int d;

    while (ia < a_size) {
        m[ia] = ia + 1;
        ia++;
    }

    while (ib < b_size) {
        int32_t c = b[ib];
        unsigned int da = ib++;

        d = da;
        ia = -1;

        while (++ia < a_size) {
            unsigned int db = (c == a[ia]) ? da : da + 1;
            da = m[ia];

            m[ia] = d = (da > d)
                ? (db > d ? d + 1 : db)
                : (db > da ? da + 1 : db);
        }
    }

    free(m);

    return d;
}
