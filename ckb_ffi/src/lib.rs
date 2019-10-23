// reference: https://github.com/yangby-cryptape/ckb-ffi
pub mod conv;
pub mod error;

#[repr(C)]
pub struct Buffer {
    len: u64,
    data: *mut u8,
}

#[no_mangle]
pub unsafe extern "C" fn ckb_encode(
    mut output: &mut Buffer,
    type_name: *const libc::c_char,
    json: *const libc::c_char,
) -> i32 {
    let mut retcode = 0;
    let type_name = cstring_to_str(type_name);
    let json_str = cstring_to_str(json);

    if let Ok(mol_bytes) = conv::try_convert(type_name, json_str.as_bytes(), true) {
        vector_into_buffer(&mut output, mol_bytes);
    } else {
        retcode = 1;
    }
    retcode
}

#[no_mangle]
pub unsafe extern "C" fn ckb_decode(
    mut output: &mut Buffer,
    type_name: *const libc::c_char,
    mol: *const Buffer,
) -> i32 {
    let mut retcode = 0;
    let type_name = cstring_to_str(type_name);
    let slice_mol = buffer_to_slice(mol);

    if let Ok(json_bytes) = conv::try_convert(type_name, slice_mol, false) {
        vector_into_buffer(&mut output, json_bytes);
    } else {
        retcode = 1;
    }
    retcode
}

#[no_mangle]
pub extern "C" fn buffer_free(buf: Buffer) {
    let slice = unsafe { std::slice::from_raw_parts_mut(buf.data, buf.len as usize) };
    let slice = slice.as_mut_ptr();
    unsafe {
        Box::from_raw(slice);
    }
}

fn vector_into_buffer(output: &mut Buffer, vec: Vec<u8>) {
    let mut buf = vec.into_boxed_slice();
    output.data = buf.as_mut_ptr();
    output.len = buf.len() as u64;
    std::mem::forget(buf);
}

unsafe fn buffer_to_slice(buf: *const Buffer) -> &'static [u8] {
    std::slice::from_raw_parts((*buf).data, (*buf).len as usize)
}

unsafe fn cstring_to_str(input: *const libc::c_char) -> &'static str {
    &std::ffi::CStr::from_ptr(input)
        .to_str()
        .expect("convert a C string to rust &str should be ok")
}

#[cfg(test)]
mod tests {
    use super::*;

    impl Buffer {
        pub fn new() -> Self {
            Buffer {
                len: 0,
                data: Vec::new().as_mut_ptr(),
            }
        }
    }

    #[test]
    fn should_encode_decode_transaction() {
        let tx_json = r#"{
            "version": "0x0",
            "cell_deps": [],
            "header_deps": [],
            "inputs": [],
            "outputs": [],
            "witnesses": [],
            "outputs_data": []
        }"#;

        let mut mol_buf = Buffer::new();
        let mut json_buf = Buffer::new();

        let type_name = std::ffi::CString::new("Transaction").expect("type name cstring fail");
        let json = std::ffi::CString::new(tx_json).expect("tx json cstring fail");

        unsafe {
            let retcode = ckb_encode(&mut mol_buf, type_name.as_ptr(), json.as_ptr());
            assert_eq!(retcode, 0);

            let retcode = ckb_decode(&mut json_buf, type_name.as_ptr(), &mol_buf);
            assert_eq!(retcode, 0);

            // Check transaction json
            let expect_tx: ckb_jsonrpc_types::Transaction =
                serde_json::from_str(tx_json).expect("tx_json from str fail");

            let json = buffer_to_slice(&json_buf);

            let tx: ckb_jsonrpc_types::Transaction =
                serde_json::from_slice(json).expect("json buf from slice fail");

            assert_eq!(tx, expect_tx);
        }
    }
}
