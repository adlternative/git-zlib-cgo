### git-zlib-go

golang "compress/zlib" does not provide fine-grained control over zlib behavior, e.g. get the size of the input data.

So create git-zlib-go which can control zlib more detailed.

Partially Use [git/zlib.c](https://github.com/git/git/blob/master/zlib.c) as C code, Invoked via cgo.

I only did inflate part. If you need deflate, welcome to contribute.