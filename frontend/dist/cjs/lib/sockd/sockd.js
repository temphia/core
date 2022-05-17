"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Sockd = void 0;
const ws_1 = require("../core/ws");
class Sockd {
    constructor(url) {
        this.init = async () => {
            this._ws = this._builder.build();
        };
        this.handleIncoming = (_, ev) => {
            // fixme => handle system messages
            const data = JSON.parse(ev.data);
            this._handler(data);
        };
        this.OnSockdMessage = (h) => {
            this._handler = h;
        };
        this.SendSockd = (message) => {
            this._ws.send(JSON.stringify(message));
        };
        console.log("CONNECTING WS @ ", url);
        this._builder = new ws_1.WebsocketBuilder(url);
        this._builder.onMessage(this.handleIncoming);
        this._builder.withBackoff(new ws_1.LinearBackoff(1, 3));
        this._builder.withBuffer(new ws_1.LRUBuffer(20));
    }
}
exports.Sockd = Sockd;
