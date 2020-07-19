from cffi import FFI

ffi = FFI()

lib = ffi.dlopn('../../build/libprime.so')