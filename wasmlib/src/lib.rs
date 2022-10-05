use std::mem;
use std::os::raw::{c_void};

#[no_mangle]
pub extern fn allocate(size: usize) -> *mut c_void {
    let mut buffer = Vec::with_capacity(size);
    let pointer = buffer.as_mut_ptr();
    mem::forget(buffer);

    pointer as *mut c_void
}

#[no_mangle]
pub extern fn deallocate(pointer: *mut c_void, capacity: usize) {
    unsafe {
        let _ = Vec::from_raw_parts(pointer, 0, capacity);
    }
}

#[no_mangle]
pub extern "C" fn run(data_ptr: *const u8, length: usize) -> *const u8 {
    use core::slice;
    let data: Vec<u8> = unsafe { slice::from_raw_parts(data_ptr, length).to_vec() };

    let s = String::from_utf8(data).expect("not string");

    println!("in: {s:?}");

    b"out: wasm!\0".as_ptr()
}
