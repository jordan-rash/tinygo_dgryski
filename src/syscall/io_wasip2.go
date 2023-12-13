//go:build wasip2

package syscall

type __wasi_list_u8 struct {
	data
	err
}

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]input-stream.read
func __wasi_io_streams_input_stream_read(self int32, length uint64) __wasi_result_list_error

//go:wasmimport wasi:io/streams@0.2.0-rc-2023-11-10 [method]output-stream.write
func __wasi_io_streams_output_stream_write(self int32, contect __wasi_list_u8) __wasi_result_list_error


func Write(fd int, p []byte) (n int, err error) {
_ = __wasi_io_streams_input_stream_read
}

func Read(fd int, p []byte) (n int, err error) {

}