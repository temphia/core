use super::core::{
    get_resp, read_buf_as_string, wasm_err_or_binary, wasm_err_or_nothing, wasm_err_or_str,
};
use simple_error::SimpleError;

#[link(wasm_import_module = "temphia")]
extern "C" {

    // cabinet
    fn _cabinet_add_file(
        folder_ptr: *const u8,
        folder_len: usize,
        file_ptr: *const u8,
        file_len: usize,
        contents_ptr: *const u8,
        contents_len: usize,
        rid: *const i32,
    ) -> i32;

    fn _cabinet_list_folder(folder_ptr: *const u8, folder_len: usize, rid: *const i32) -> i32;
    fn _cabinet_get_file(
        folder_ptr: *const u8,
        folder_len: usize,
        file_ptr: *const u8,
        file_len: usize,
        rid: *const i32,
    ) -> i32;

    fn _cabinet_del_file(
        folder_ptr: *const u8,
        folder_len: usize,
        file_ptr: *const u8,
        file_len: usize,
        rid: *const i32,
    ) -> i32;
}

#[derive(Default)]
pub struct Cabinet {
    folder: String,
}

impl Cabinet {
    pub fn add_file(&self, file: &str, data: Vec<u8>) -> Result<(), SimpleError> {
        let rid = 0;
        let resp = unsafe {
            _cabinet_add_file(
                self.folder.as_ptr(),
                self.folder.len(),
                file.as_ptr(),
                file.len(),
                data.as_ptr(),
                data.len(),
                &rid,
            )
        };

        wasm_err_or_nothing(resp, rid)
    }

    pub fn list_folder(&self) -> Result<Vec<String>, SimpleError> {
        let rid = 0;
        let resp = unsafe { _cabinet_list_folder(self.folder.as_ptr(), self.folder.len(), &rid) };

        let fresp = wasm_err_or_str(resp, rid);

        return match fresp {
            Ok(rstr) => {
                let mut v = Vec::new();
                rstr.split(',').for_each(|x| {
                    v.push(x.to_string());
                });
                Ok(v)
            }

            Err(e) => Err(e),
        };
    }

    pub fn get_file(&self, file: &str) -> Result<Vec<u8>, SimpleError> {
        let rid = 0;
        let resp = unsafe {
            _cabinet_get_file(
                self.folder.as_ptr(),
                self.folder.len(),
                file.as_ptr(),
                file.len(),
                &rid,
            )
        };
        wasm_err_or_binary(resp, rid)
    }

    pub fn del_file(&self, file: &str) -> Result<(), SimpleError> {
        let rid = 0;
        let resp = unsafe {
            _cabinet_del_file(
                self.folder.as_ptr(),
                self.folder.len(),
                file.as_ptr(),
                file.len(),
                &rid,
            )
        };

        wasm_err_or_nothing(resp, rid)
    }
}
