#![crate_type = "staticlib"]

use rand::Rng;

#[no_mangle]
pub unsafe extern "C" fn rust_procedure(
    caller_language: *const libc::c_char,
    caller_stack: libc::c_int
) {
    let callee_stack = rand::thread_rng().gen_range(1..10_000);
    let msg_str: &str = 
        std::ffi::CStr
        ::from_ptr(caller_language)
        .to_str()
        .unwrap_or("unknown");

    println!(
        "[{:0>4}] start of a rust procedure, called by {} [{:0>4}]", 
        callee_stack, 
        msg_str, 
        caller_stack,
    );

    println!(
        "[{:0>4}] end of a rust procedure",
        callee_stack, 
    );
}
