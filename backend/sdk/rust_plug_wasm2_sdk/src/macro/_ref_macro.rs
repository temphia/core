// expected sysntax
#[mylib::main[opt1="xyz", opt2="mno"]]
fn event_handler1(event: EventRequest) -> EventResponse {
    // do sth coll with event

    return EventResponse {};
}

// transformed entry
extern "C" fn entry_event_handler1(arg1: i32, arg2: i32) {
    let req = read_req(arg1, arg2);
    let resp = event_handler1(req);
    write_resp(resp);
}

// mylib stuff

struct EventRequest {}
struct EventResponse {}

pub fn read_req(arg1: i32, arg2: i32) -> EventRequest {
    EventRequest {}
}

pub fn write_resp(resp: EventResponse) {}

/*

 A =>: A basic example using macro_rules! (and no options yet):
#[doc(hidden)] /** Not part of the public API */ pub
use ::paste::paste as __paste;

#[macro_export]
macro_rules! my_main {(
    fn $fname:ident ( $($args:tt)* )
    $(-> $Ret:ty)?
    $body:block
) => (
  $crate::__paste! {
    #[no_mangle] pub extern "C"
    fn [< entry_ $fname >] (
         arg1: ::core::primitive::i32,
         arg2: ::core::primitive::i32,
    )
    {
        let req = $crate::read_req(arg1, arg2);
        let resp = ({
            fn $fname ($($args)*)
            $(-> $Ret)?
            $body

            $fname
        })(req);
        $crate::write_resp(resp);
    }
  }
)}

[7:34 PM] A =>: You'd call it as
my_main! {
    fn event_handler1(event: EventRequest) -> EventResponse {
        // do sth coll with event

        return EventResponse {};
    }
}

or, if using https://docs.rs/macro_rules_attribute, as
#[apply(my_main!)]
fn event_handler1(event: EventRequest) -> EventResponse {
    // do sth coll with event

    return EventResponse {};
}
[7:35 PM] A =>: If you want to directly feature a #[my_main] kind of attribute, and also handle extra modifiers such as opt1 = "xyz" and so on, you'll need to look into proc-macros.

*/
