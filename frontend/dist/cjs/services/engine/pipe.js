"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.IFramePipe = void 0;
class IFramePipe {
    constructor(secret) {
        this.set_handler = (fn) => {
            this._handlers.add(fn);
        };
        this.remove_handler = (fn) => {
            this._handlers.delete(fn);
        };
        this.send = (xid, action, data) => {
            const message = JSON.stringify({
                xid,
                data,
                action,
                parent_secret: this._secret,
            });
            window.parent.postMessage(message, '*');
        };
        this._secret = secret;
        this._handlers = new Set();
        window.addEventListener('message', (ev) => {
            const decoded = JSON.parse(ev.data);
            this._handlers.forEach((fn) => fn(decoded.xid, decoded.action, decoded.data));
        });
    }
}
exports.IFramePipe = IFramePipe;
