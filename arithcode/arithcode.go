package arithcode

/*
#cgo CFLAGS: -I. -O2
#include "arithcode.h"
#include <string.h>

// encode_buf wraps encode_u8_u8 to avoid passing a Go pointer-to-pointer to C.
// cdf must have 257 elements (for 256 byte symbols).
// Returns 0 on success, negative on failure (e.g. output buffer too small).
static int encode_buf(unsigned char *out, size_t out_cap, size_t *out_len,
                      unsigned char *in, size_t in_len, float *cdf) {
	void *p = out;
	size_t n = out_cap;
	int ret = encode_u8_u8(&p, &n, in, in_len, cdf, 256);
	*out_len = n;
	return ret;
}
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// encodeMu serializes calls to encode_u8_u8 because the C implementation uses
// a function-local static variable and is not thread-safe.
var encodeMu sync.Mutex

// Encode compresses data using the provided 257-element CDF (cdf[0]=0.0, cdf[256]=1.0).
// The CDF must cover all 256 possible byte values with non-zero probabilities.
// Returns nil on failure (e.g. output buffer overflow for highly incompressible data).
func Encode(data []byte, cdf *[257]float32) []byte {
	if len(data) == 0 {
		return []byte{}
	}

	encodeMu.Lock()
	defer encodeMu.Unlock()

	// Pre-allocate output: arithmetic coding worst case is ~input*1.1 + framing.
	// 2x + 512 bytes is safe for all Meshtastic packet sizes.
	outBuf := make([]byte, len(data)*2+512)
	var outLen C.size_t

	ret := C.encode_buf(
		(*C.uchar)(unsafe.Pointer(&outBuf[0])),
		C.size_t(len(outBuf)),
		&outLen,
		(*C.uchar)(unsafe.Pointer(&data[0])),
		C.size_t(len(data)),
		(*C.float)(unsafe.Pointer(&cdf[0])),
	)
	runtime.KeepAlive(outBuf)
	runtime.KeepAlive(data)
	runtime.KeepAlive(cdf)

	if ret != 0 {
		return nil
	}

	n := int(outLen)
	result := make([]byte, n)
	copy(result, outBuf[:n])
	return result
}
