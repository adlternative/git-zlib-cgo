package git_zlib_cgo

/*
#cgo CFLAGS: -Werror=implicit
#cgo pkg-config: zlib

#include "git-zlib.h"

char* git_zstream_msg(char *strm) {
  return (((git_zstream*)strm)->z).msg;
}

int git_zstream_inflate_init(char*strm) {
	return git_inflate_init((git_zstream*)strm);
}

int git_zstream_inflate_end(char*strm) {
	return git_inflate_end((git_zstream*)strm);
}

int git_zstream_inflate(char*strm, int flush) {
	return git_inflate((git_zstream*)strm, flush);
}

unsigned int git_zstream_total_in(char *strm) {
  return ((git_zstream*)strm)->total_in;
}

unsigned int git_zstream_total_out(char *strm) {
  return ((git_zstream*)strm)->total_out;
}

unsigned int git_zstream_avail_out(char *strm) {
  return ((git_zstream*)strm)->avail_out;
}

unsigned int git_zstream_avail_in(char *strm) {
  return ((git_zstream*)strm)->avail_in;
}

void git_zstream_set_in_buf(char *strm, void *buf, unsigned int len) {
  ((git_zstream*)strm)->next_in = (unsigned char *)buf;
  ((git_zstream*)strm)->avail_in = len;
}

void git_zstream_set_out_buf(char *strm, void *buf, unsigned int len) {
  ((git_zstream*)strm)->next_out = (unsigned char*)buf;
  ((git_zstream*)strm)->avail_out = len;
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type GitZStream [unsafe.Sizeof(C.git_zstream{})]C.char

func (strm *GitZStream) SetInBuf(buf []byte, size int) {
	if buf == nil {
		C.git_zstream_set_in_buf(&strm[0], nil, C.uint(size))
	} else {
		C.git_zstream_set_in_buf(&strm[0], unsafe.Pointer(&buf[0]), C.uint(size))
	}
}

func (strm *GitZStream) SetOutBuf(buf []byte, size int) {
	if buf == nil {
		C.git_zstream_set_out_buf(&strm[0], nil, C.uint(size))
	} else {
		C.git_zstream_set_out_buf(&strm[0], unsafe.Pointer(&buf[0]), C.uint(size))
	}
}

func (strm *GitZStream) AvailIn() int {
	return int(C.git_zstream_avail_in(&strm[0]))
}

func (strm *GitZStream) TotalIn() int {
	return int(C.git_zstream_total_in(&strm[0]))
}

func (strm *GitZStream) TotalOut() int {
	return int(C.git_zstream_total_out(&strm[0]))
}

func (strm *GitZStream) AvailOut() int {
	return int(C.git_zstream_avail_out(&strm[0]))
}

func (strm *GitZStream) msg() string {
	return C.GoString(C.git_zstream_msg(&strm[0]))
}

func (strm *GitZStream) InflateInit() error {
	result := C.git_zstream_inflate_init(&strm[0])
	if result != Z_OK {
		return fmt.Errorf("git-zlib: inflate init failed with (%v): %v", result, strm.msg())
	}

	return nil
}

func (strm *GitZStream) InflateEnd() error {
	result := C.git_zstream_inflate_end(&strm[0])
	if result != Z_OK {
		return fmt.Errorf("git-zlib: inflate end failed with (%v): %v", result, strm.msg())
	}

	return nil
}

func (strm *GitZStream) Inflate(flag int) (int, error) {
	ret := C.git_zstream_inflate(&strm[0], C.int(flag))
	switch ret {
	case Z_NEED_DICT:
		ret = Z_DATA_ERROR
		fallthrough
	case Z_DATA_ERROR, Z_MEM_ERROR:
		return int(ret), fmt.Errorf("git-zlib: failed to inflate (%v): %v", ret, strm.msg())
	}
	return int(ret), nil
}
