fn main() {
    println!("cargo:rustc-link-search=all=endpoint");    // rustc -L endpoint ...
    println!("cargo:rustc-link-lib=dylib=c_procedure");  // rustc -l libc_procedure.a
}
