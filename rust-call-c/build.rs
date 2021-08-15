fn main() {
    println!("cargo:rustc-link-search=all=endpoint");     // rustc -L endpoint ...
    println!("cargo:rustc-link-lib=static=c_procedure");  // rustc -l c_procedure
}
