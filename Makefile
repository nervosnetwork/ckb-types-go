all: build package

build:
	cd ckb_ffi && cargo build --release
	cp ckb_ffi/target/release/libckb_ffi.a ./libckb_ffi.a
	c-for-go -ccincl jsonrpc_types.yml

package:
	mv ./generated/types.go ./jsonrpc/types/buf_ffi.go
	mv ./generated/cgo_helpers.* ./jsonrpc/types/
	mv ./generated/generated.go ./jsonrpc/types/codec_ffi.go
	mv ./generated/const.go     ./jsonrpc/types/const.go
	sed -i 's/package generated/package types/' ./jsonrpc/types/buf_ffi.go
	sed -i 's/package generated/package types/' ./jsonrpc/types/cgo_helpers.go
	sed -i 's/package generated/package types/' ./jsonrpc/types/codec_ffi.go
	sed -i 's/package generated/package types/' ./jsonrpc/types/const.go
	rm -r ./generated

clean:
	rm ./libckb_ffi.a
	rm ./jsonrpc/types/buf_ffi.go ./jsonrpc/types/cgo_helpers.* ./jsonrpc/types/codec_ffi.go

clean-all:
	clean
	rm -r ckb_ffi/target
