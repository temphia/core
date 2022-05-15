use simple_error::SimpleError;

#[derive(Debug, Default)]
pub struct EventRequest {
    pub id: String,
    pub r#type: String,
    pub name: String,
    pub ctx_vars: Option<Vec<u8>>,
    pub payload: Option<Vec<u8>>,
}

#[link(wasm_import_module = "temphia")]
extern "C" {
    // executor

    fn _get_event_step1(
        load_paylaod: i32,
        load_ctxvar: i32,
        r#id_len: *const i32,
        r#type_len: *const i32,
        r#name_len: *const i32,
        r#ctx_var_len: *const i32,
        r#data_len: *const i32,
    );

    fn _get_event_step2(
        idptr: *const u8,
        typeptr: *const u8,
        nameptr: *const u8,
        ctxvarptr: *const u8,
        pptr: *const u8,
    );

    fn _set_event_reply(meta_ptr: *const u8, meta_len: usize, data_ptr: *const u8);
    fn _get_resp(ptr: *const u8, rid: i32);

    fn _log(ptr: *const u8, len: usize);
    fn _lazy_log(ptr: *const u8, len: usize);
    fn _sleep(time: i32);
    fn _get_self_file(ptr: *const u8, len: usize, rid: *const i32) -> i32;

}

pub unsafe fn get_resp(ptr: *const u8, rid: i32) {
    _get_resp(ptr, rid)
}

pub struct Core {}

impl Core {
    pub fn log(msg: &str) {
        unsafe { _log(msg.as_ptr(), msg.len()) }
    }

    pub fn lazy_log(msg: &str) {
        unimplemented!()
    }

    pub fn sleep(time: i32) {
        unsafe {
            _sleep(time);
        }
    }

    pub fn get_self_file(file: &str) -> Result<Vec<u8>, SimpleError> {
        let rid = 0;
        let resp;
        unsafe {
            resp = _get_self_file(file.as_ptr(), file.len(), &rid);
        }

        wasm_err_or_binary(resp, rid)
    }

    pub fn get_event_ctx(ctxvar: bool, payload: bool) -> EventRequest {
        let load_payload = if ctxvar { 1 } else { 0 };
        let load_ctxvar = if payload { 1 } else { 0 };
        let id_len = 0;
        let type_len = 0;
        let name_len = 0;
        let ctx_var_len = 0;
        let data_len = 0;

        unsafe {
            _get_event_step1(
                load_payload,
                load_ctxvar,
                &id_len,
                &type_len,
                &name_len,
                &ctx_var_len,
                &data_len,
            );
        }

        let mut req = EventRequest {
            id: {
                let size = id_len as usize;
                let mut s = String::with_capacity(size);
                unsafe {
                    s.as_mut_vec().set_len(size);
                }
                s
            },
            name: {
                let size = name_len as usize;
                let mut s = String::with_capacity(size);
                unsafe {
                    s.as_mut_vec().set_len(size);
                }
                s
            },
            r#type: {
                let size = type_len as usize;
                let mut s = String::with_capacity(size);
                unsafe {
                    s.as_mut_vec().set_len(size);
                }
                s
            },
            ctx_vars: None,
            payload: None,
        };

        let mut ctxvarptr: *const u8 = std::ptr::null();
        if ctx_var_len != 0 {
            req.ctx_vars = {
                let size = ctx_var_len as usize;
                let mut v: Vec<u8> = Vec::with_capacity(size);
                unsafe {
                    v.set_len(size);
                }
                ctxvarptr = v.as_ptr();
                Some(v)
            };
        }

        let mut dataptr: *const u8 = std::ptr::null();
        if data_len != 0 {
            req.payload = {
                let size = data_len as usize;
                let mut v: Vec<u8> = Vec::with_capacity(size);
                unsafe {
                    v.set_len(size);
                }
                dataptr = v.as_ptr();

                Some(v)
            };
        }

        unsafe {
            _get_event_step2(
                req.id.as_mut_ptr(),
                req.r#type.as_mut_ptr(),
                req.name.as_mut_ptr(),
                ctxvarptr,
                dataptr,
            );
        }
        req
    }

    pub fn set_reply(data: *const u8) {}

    pub fn set_reply_with_meta(meta_ptr: *const u8, data_ptr: *const u8) {}
}

pub fn read_buf_as_vec_bytes(size: usize, rid: i32) -> Vec<u8> {
    let mut total: Vec<u8> = vec![0u8; size];

    Core::log(format!("VEC IS {} long", total.len()).as_str());

    unsafe { _get_resp(total.as_mut_ptr() as *const u8, rid) };
    total
}

pub fn read_buf_as_string(size: usize, rid: i32) -> String {
    let mut total = String::with_capacity(size);
    unsafe {
        total.as_mut_vec().set_len(size);
        _get_resp(total.as_mut_ptr() as *const u8, rid)
    };
    return total;
}

pub fn wasm_err_or_nothing(resp: i32, rid: i32) -> Result<(), SimpleError> {
    return if resp < 0 {
        Err(SimpleError::new(read_buf_as_string(-resp as usize, rid)))
    } else {
        Ok(())
    };
}

pub fn wasm_err_or_binary(resp: i32, rid: i32) -> Result<Vec<u8>, SimpleError> {
    return if resp < 0 {
        Err(SimpleError::new(read_buf_as_string(-resp as usize, rid)))
    } else {
        Ok(read_buf_as_vec_bytes(resp as usize, rid))
    };
}

pub fn wasm_err_or_str(resp: i32, rid: i32) -> Result<String, SimpleError> {
    return if resp < 0 {
        Err(SimpleError::new(read_buf_as_string(-resp as usize, rid)))
    } else {
        Ok(read_buf_as_string(resp as usize, rid))
    };
}
