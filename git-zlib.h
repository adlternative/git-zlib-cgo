#include <zlib.h>

typedef struct git_zstream {
	z_stream z;
	unsigned long avail_in;
	unsigned long avail_out;
	unsigned long total_in;
	unsigned long total_out;
	unsigned char *next_in;
	unsigned char *next_out;
} git_zstream;

int git_inflate_init(git_zstream *);
int git_inflate_end(git_zstream *);
int git_inflate(git_zstream *, int flush);
