var _buffer = [];
var core = {
    log: function (message) { return _log(message); },
    log_lazy: function (message) { return _buffer.push(message); },
    lazy_log_send: function () {
        _log_lazy(_buffer);
        _buffer.length = 0;
    },
    sleep: function (t) { return _sleep(t); },
    self_file: function (file) { return _get_self_file_as_str(file); }
};

// this is noop method to prevent entry function to treeshaking/deadcode eleminiation 
// or whater by rollup/typescript
var registerHandler = function (handler) {
    JSON.stringify(handler);
};

function main(event) {
    core.log("This is test" + JSON.stringify(event));
}
registerHandler(main);
