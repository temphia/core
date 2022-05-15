use super::core::{
    read_buf_as_string, wasm_err_or_binary, wasm_err_or_nothing, wasm_err_or_str,
};
use serde::{Deserialize, Serialize};
use serde_json;
use simple_error::SimpleError;

#[link(wasm_import_module = "temphia")]
extern "C" {

    // plugKV
    fn _plugkv_set(
        txid: i32,
        key_ptr: *const u8,
        key_len: usize,
        val_ptr: *const u8,
        val_len: usize,
        opt_ptr: *const u8,
        opt_len: usize,
        rid: *const i32,
    ) -> i32;
    fn _plugkv_update(
        txid: i32,
        key_ptr: *const u8,
        key_len: usize,
        val_ptr: *const u8,
        val_len: usize,
        opt_ptr: *const u8,
        opt_len: usize,
        rid: *const i32,
    ) -> i32;
    fn _plugkv_get(txid: i32, key_ptr: *const u8, key_len: usize, rid: *const i32) -> i32;
    fn _plugkv_del(txid: i32, key_ptr: *const u8, key_len: usize, rid: *const i32) -> i32;
    fn _plugkv_del_batch(txid: i32, keys_ptr: *const u8, keys_len: usize, rid: *const i32) -> i32;
    fn _plugkv_query(txid: i32, qptr: *const u8, qlen: usize, rid: *const i32) -> i32;

    fn _plugkv_new_txn(rid: *const i32) -> i32;
    fn _plugkv_rollback(txid: i32, rid: *const i32) -> i32;
    fn _plugkv_commit(txid: i32, rid: *const i32) -> i32;

}

pub struct PlugKV {
    txid: i32,
}

impl PlugKV {
    pub fn new() -> Self {
        Self { txid: 0 }
    }

    pub fn quick_set(&self, key: &str, val: &str) -> Result<(), SimpleError> {
        let rid = 0;
        let resp = unsafe {
            _plugkv_set(
                self.txid,
                key.as_ptr(),
                key.len(),
                val.as_ptr(),
                val.len(),
                std::ptr::null(),
                0,
                &rid,
            )
        };

        wasm_err_or_nothing(resp, rid)
    }

    pub fn quick_update(&self, key: &str, val: &str) -> Result<(), SimpleError> {
        let rid = 0;
        let resp = unsafe {
            _plugkv_update(
                self.txid,
                key.as_ptr(),
                key.len(),
                val.as_ptr(),
                val.len(),
                std::ptr::null(),
                0,
                &rid,
            )
        };
        wasm_err_or_nothing(resp, rid)
    }

    pub fn quick_get(&self, key: &str) -> Result<String, SimpleError> {
        let rid = 0;
        let resp = unsafe { _plugkv_get(self.txid, key.as_ptr(), key.len(), &rid) };
        
        match wasm_err_or_str(resp, rid) {
            Ok(x) => {
                let value = gjson::get(&x, "value").str().to_string();
                return Ok(value);
            }
            Err(e) => return Err(e),
        }
    }

    pub fn delete(&self, key: &str) -> Result<(), SimpleError> {
        let rid = 0;
        let resp = unsafe { _plugkv_del(self.txid, key.as_ptr(), key.len(), &rid) };
        wasm_err_or_nothing(resp, rid)
    }

    pub fn batch_delete(&self, keys: Vec<&str>) -> Result<(), SimpleError> {
        let key = keys.join(",");
        let rid = 0;
        let resp = unsafe { _plugkv_del_batch(self.txid, key.as_ptr(), key.len(), &rid) };
        wasm_err_or_nothing(resp, rid)
    }

    pub fn query(&self, qstr: &str) -> Result<Vec<PlugValue>, SimpleError> {
        let rid = 0;
        let resp = unsafe { _plugkv_query(self.txid, qstr.as_ptr(), qstr.len(), &rid) };

        let data = wasm_err_or_binary(resp, rid)?;

        let value: Result<Vec<PlugValue>, serde_json::Error> = serde_json::from_slice(&data);

        return match value {
            Ok(x) => Ok(x),
            Err(e) => Err(SimpleError::from(e)),
        };
    }

    pub fn new_txt(&self) -> Result<Self, SimpleError> {
        let rid = 0;
        if self.txid != 0 {
            return Err(SimpleError::new("nested txn not alowed"));
        }

        let resp = unsafe { _plugkv_new_txn(&rid) };

        if resp > 0 {
            return Ok(Self { txid: resp });
        }
        Err(SimpleError::new(read_buf_as_string(-resp as usize, rid)))
    }

    pub fn rollback(&self) -> Result<(), SimpleError> {
        let rid = 0;
        if self.txid == 0 {
            return Err(SimpleError::new("not inside txn"));
        }

        let resp = unsafe { _plugkv_rollback(self.txid, &rid) };
        wasm_err_or_nothing(resp, rid)
    }

    pub fn commit(&self) -> Result<(), SimpleError> {
        let rid = 0;
        if self.txid == 0 {
            return Err(SimpleError::new("not inside txn"));
        }

        let resp = unsafe { _plugkv_commit(self.txid, &rid) };
        wasm_err_or_nothing(resp, rid)
    }
}

#[derive(Debug, Serialize, Deserialize)]
pub struct PlugValue {
    pub key: String,
    pub value: String,
    pub tag1: Option<String>,
    pub tag2: Option<String>,
    pub tag3: Option<String>,
    pub ttl: Option<String>,
}

#[derive(Debug)]
pub struct SetOptions<'a> {
    pub tag1: Option<&'a str>,
    pub tag2: Option<&'a str>,
    pub tag3: Option<&'a str>,
    pub ttl: i64,
}

#[derive(Debug)]
pub struct UpdateOptions<'a> {
    pub force_ver_update: bool,
    pub with_version: bool,
    pub version: i32,
    pub tag1: Option<&'a str>,
    pub tag2: Option<&'a str>,
    pub tag3: Option<&'a str>,
    pub ttl: i64,
}

#[derive(Debug)]
pub struct QueryOptions<'a> {
    pub key_prefix: Option<&'a str>,
    pub load_meta: bool,
    pub tag1: Option<Vec<&'a str>>,
    pub tag2: Option<Vec<&'a str>>,
    pub tag3: Option<Vec<&'a str>>,
    pub page: i32,
    pub page_count: i32,
}
