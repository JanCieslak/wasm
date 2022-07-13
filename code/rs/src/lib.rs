use wasm_bindgen::prelude::*;

#[wasm_bindgen]
extern {
    pub fn alert(msg: &str);
}

#[wasm_bindgen]
pub fn Calculate(n: i32) -> i32 {
    return 4 * n;
}

