use super::core::wasm_err_or_nothing;
use simple_error::SimpleError;

#[link(wasm_import_module = "temphia")]
extern "C" {

    // sockd
    fn _sockd_send_direct(
        room_ptr: *const u8,
        room_len: usize,
        conns_ptr: *const u8,
        conns_len: usize,
        payload_ptr: *const u8,
        payload_len: usize,
        rid: *const i32,
    ) -> i32;
    fn _sockd_send_broadcast(
        room_ptr: *const u8,
        room_len: usize,
        payload_ptr: *const u8,
        payload_len: usize,
        rid: *const i32,
    ) -> i32;
    fn _sockd_send_tagged(
        room_ptr: *const u8,
        room_len: usize,
        tags_ptr: *const u8,
        tags_len: usize,
        ignore_ptr: *const u8,
        ignore_len: usize,
        payload_ptr: *const u8,
        payload_len: usize,
        rid: *const i32,
    ) -> i32;
    fn _sockd_add_to_room(
        room_ptr: *const u8,
        room_len: usize,
        conn_ptr: *const u8,
        conn_len: usize,
        tags_ptr: *const u8,
        tags_len: usize,
        rid: *const i32,
    ) -> i32;
    fn _sockd_kick_from_room(
        room_ptr: *const u8,
        room_len: usize,
        conn_ptr: *const u8,
        conn_len: usize,
        rid: *const i32,
    ) -> i32;
    fn _sockd_list_room_conns(room_ptr: *const u8, room_len: usize, rid: *const i32) -> i32;
    fn _sockd_ban_conn(conn_ptr: *const u8, conn_len: usize, rid: *const i32) -> i32;
}

pub struct SockdRoom {
    room: String,
}

impl SockdRoom {
    pub fn new(room: &str) -> Self {
        Self {
            room: room.to_string(),
        }
    }

    pub fn is_defined() -> bool {
        return true;
    }

    pub fn send_direct(&self, conns: Vec<&str>, data: &[u8]) -> Result<(), SimpleError> {
        let _conns = conns.join(",");
        let rid = 0;
        let resp = unsafe {
            _sockd_send_direct(
                self.room.as_ptr(),
                self.room.len(),
                _conns.as_ptr(),
                _conns.len(),
                data.as_ptr(),
                data.len(),
                &rid,
            )
        };

        wasm_err_or_nothing(resp, rid)
    }

    pub fn send_broadcast(&self, data: &[u8]) -> Result<(), SimpleError> {
        let rid = 0;

        let resp = unsafe {
            _sockd_send_broadcast(
                self.room.as_ptr(),
                self.room.len(),
                data.as_ptr(),
                data.len(),
                &rid,
            )
        };

        wasm_err_or_nothing(resp, rid)
    }

    pub fn send_tagged(&self, tags: Vec<&str>, data: &[u8]) -> Result<(), SimpleError> {
        let _tags = tags.join(",");
        let rid = 0;

        let resp = unsafe {
            _sockd_send_tagged(
                self.room.as_ptr(),
                self.room.len(),
                _tags.as_ptr(),
                _tags.len(),
                std::ptr::null(),
                0,
                data.as_ptr(),
                data.len(),
                &rid,
            )
        };

        wasm_err_or_nothing(resp, rid)
    }

    pub fn add_to_room(&self, conn: &str, tags: Vec<&str>) -> Result<(), SimpleError> {
        let rid = 0;
        let _tags = tags.join(",");

        let resp = unsafe {
            _sockd_add_to_room(
                self.room.as_ptr(),
                self.room.len(),
                conn.as_ptr(),
                conn.len(),
                _tags.as_ptr(),
                _tags.len(),
                &rid,
            )
        };

        wasm_err_or_nothing(resp, rid)
    }

    pub fn kick_from_room(&self, conn: &str, tags: Vec<&str>) -> Result<(), SimpleError> {
        let _tags = tags.join(",");
        let rid = 0;

        let resp = unsafe {
            _sockd_kick_from_room(
                self.room.as_ptr(),
                self.room.len(),
                conn.as_ptr(),
                conn.len(),
                &rid,
            )
        };

        wasm_err_or_nothing(resp, rid)
    }

    pub fn ban_conn(&self, conn: &str) -> Result<(), SimpleError> {
        let rid = 0;
        let resp = unsafe { _sockd_ban_conn(conn.as_ptr(), conn.len(), &rid) };
        wasm_err_or_nothing(resp, rid)
    }
}
