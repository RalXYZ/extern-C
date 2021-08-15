use rand::Rng;

extern "C" {
    fn c_procedure(
        caller_language: *const libc::c_char,
        caller_stack: libc::c_int
    ) -> ();
}

fn main() {
    let caller_stack = rand::thread_rng().gen_range(1..10_000);
    println!("[{:0>4}] start of rust main procedure", caller_stack);

    let language = std::ffi::CString::new("rust").unwrap();
    unsafe { 
        c_procedure(language.as_ptr(), caller_stack) 
    };
    
    println!("[{:0>4}] end of rust main procedure", caller_stack);
}