use super::core::wasm_err_or_binary;

use simple_error::SimpleError;
use std::{collections::HashMap, io::Write};

static METHOD_GET: &'static str = "GET";
static METHOD_POST: &'static str = "POST";
static METHOD_PATCH: &'static str = "PATCH";
static METHOD_PUT: &'static str = "PUT";
static METHOD_DELETE: &'static str = "DELETE";

extern "C" {
    fn _http_call(
        method_ptr: *const u8,
        method_len: usize,
        url_ptr: *const u8,
        url_len: usize,
        header_ptr: *const u8,
        header_len: usize,

        body_ptr: *const u8,
        body_len: usize,

        resp_total_len: *const i32,
        resp_head_len: *const i32,
        rid: *const i32,
    ) -> i32; // resp(i32) => response(200, 401)
}

pub struct Builder {
    url: String,
    headers: Vec<u8>,
}

impl Builder {
    pub fn new(url: &str) -> Self {
        Self {
            url: url.to_string(),
            headers: Vec::new(),
        }
    }

    pub fn reset_headers(&mut self) {
        self.headers.clear();
    }

    pub fn set_headers(&mut self, key: &str, value: &str) {
        assert!(key
            .chars()
            .any(|x| x.is_control() || "(),/:;<=>?@[\\]{}".contains(x)));

        assert!(value.chars().any(|x| x.is_control()));
        self.headers
            .write_all(format!("{}:{}\n", key, value).as_bytes())
            .unwrap();
    }

    pub fn get(&self) -> Response {
        self.do_method(METHOD_GET, None)
    }

    pub fn post(&self, body: Option<Vec<u8>>) -> Response {
        self.do_method(METHOD_POST, body)
    }

    pub fn put(&self, body: Option<Vec<u8>>) -> Response {
        self.do_method(METHOD_PUT, body)
    }

    pub fn patch(&self, body: Option<Vec<u8>>) -> Response {
        self.do_method(METHOD_PATCH, body)
    }

    pub fn delete(&self, body: Option<Vec<u8>>) -> Response {
        self.do_method(METHOD_DELETE, body)
    }

    pub fn do_method(&self, method: &str, body: Option<Vec<u8>>) -> Response {
        let resp_total_len = 0;
        let resp_head_len = 0;
        let rid = 0;

        let status = match body {
            Some(x) => unsafe {
                _http_call(
                    method.as_ptr(),
                    method.len(),
                    self.url.as_ptr(),
                    self.url.len(),
                    self.headers.as_ptr(),
                    self.headers.len(),
                    x.as_ptr(),
                    x.len(),
                    &resp_total_len,
                    &resp_head_len,
                    &rid,
                )
            },
            None => unsafe {
                _http_call(
                    method.as_ptr(),
                    method.len(),
                    self.url.as_ptr(),
                    self.url.len(),
                    self.headers.as_ptr(),
                    self.headers.len(),
                    std::ptr::null(),
                    0,
                    &resp_total_len,
                    &resp_head_len,
                    &rid,
                )
            },
        };


        let mut data  = wasm_err_or_binary(resp_total_len, rid).unwrap();
        let (left, right) = data.split_at_mut(resp_head_len as usize);

        Response {
            body: right.to_vec(),
            headers: left.to_vec(),
            status: status,
        }
    }
}

pub struct Response {
    pub status: i32,
    pub headers: Vec<u8>,
    pub body: Vec<u8>,
}

impl Response {
    pub fn parse_header(&self) -> Result<HashMap<String, String>, SimpleError> {
        let mut hm = HashMap::new();
        for chunk in self.headers.split(|b| *b == b'\n') {
            let mut parts = chunk.splitn(2, |b| *b == b':');

            let key = parts.next().ok_or("could not parse header")?;
            let value = parts.next().ok_or("could not parse header")?;

            let _key = match std::str::from_utf8(key) {
                Err(err) => {
                    return Err(SimpleError::from(err));
                }
                Ok(x) => x.to_string(),
            };

            let _value = match std::str::from_utf8(value) {
                Err(err) => {
                    return Err(SimpleError::from(err));
                }
                Ok(x) => x.to_string(),
            };
            hm.insert(_key, _value);
        }
        Ok(hm)
    }

    pub fn get_header(&self, key: &str) -> Result<String, SimpleError> {
        for chunk in self.headers.split(|b| *b == b'\n') {
            let mut parts = chunk.splitn(2, |b| *b == b':');

            let ekey = parts.next().ok_or("could not parse header key")?;
            let value = parts.next().ok_or("Could not parse value")?;

            let _ekey = match std::str::from_utf8(ekey) {
                Err(err) => {
                    return Err(SimpleError::from(err));
                }
                Ok(x) => x.to_string(),
            };
            if _ekey != key {
                continue;
            }
            let _value = match std::str::from_utf8(value) {
                Err(err) => {
                    return Err(SimpleError::from(err));
                }
                Ok(x) => x.to_string(),
            };

            return Ok(_value);
        }

        Err(SimpleError::new("Not found"))
    }
}
