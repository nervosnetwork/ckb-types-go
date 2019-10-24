all: build package

build:
	cd ckb_ffi && cargo build --release
	cp ckb_ffi/target/release/libckb_ffi.a ./libckb_ffi.a
	c-for-go -ccincl types.yml

package:
	mv ./generated/types.go ./types/buf_ffi.go
	mv ./generated/cgo_helpers.* ./types/
	mv ./generated/generated.go ./types/codec_ffi.go
	mv ./generated/const.go     ./types/const.go
	sed -i 's/package generated/package types/' ./types/buf_ffi.go
	sed -i 's/package generated/package types/' ./types/cgo_helpers.go
	sed -i 's/package generated/package types/' ./types/codec_ffi.go
	sed -i 's/package generated/package types/' ./types/const.go
	rm -r ./generated

clean:
	rm ./libckb_ffi.a
	rm ./types/buf_ffi.go ./types/cgo_helpers.* ./types/codec_ffi.go

clean-all:
	clean
	rm -r ckb_ffi/target
