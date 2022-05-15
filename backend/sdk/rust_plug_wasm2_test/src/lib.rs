use rust_plug_wasm2_sdk::intrinsic;
use intrinsic::{core, plugkv};

#[no_mangle]
pub extern "C" fn hello_wasm() {
    core::Core::log("START =>");
    let eq = core::Core::get_event_ctx(true, true);

    let str = format!("{:?}", eq);

    core::Core::log(&str);
    core::Core::log("END =>");


    {
        // plug test
        let pkv = plugkv::PlugKV::new();




        pkv.quick_get("err_one").expect_err("set err");
        pkv.quick_get("err_2").expect("value1");





    }
}