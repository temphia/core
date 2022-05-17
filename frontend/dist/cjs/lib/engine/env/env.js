"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Env = void 0;
const folder_1 = require("../../api/folder");
const sockd_1 = require("../../sockd/sockd");
const room_1 = require("../../sockd/room");
const fetch_1 = require("./fetch");
class Env {
    constructor(opts) {
        this.init = async () => {
            await this._sockd.init();
        };
        this.PreformAction = async (name, data) => {
            const encoded = JSON.stringify(data);
            try {
                const resp = await this._fetch(name, encoded);
                const ctype = resp.headers.get("Content-Type");
                if (resp.status !== 200) {
                    const txt = await resp.text();
                    return {
                        status_ok: false,
                        content_type: ctype,
                        body: txt,
                    };
                }
                const respData = await resp.json();
                return {
                    body: respData,
                    content_type: ctype,
                    status_ok: true,
                };
            }
            catch (error) {
                return {
                    status_ok: false,
                    body: error,
                };
            }
        };
        this.startup_payload = () => {
            return this._startup_payload;
        };
        this.PreformParentAction = async (name, data) => {
            const key = "fixme => generate";
            const p = new Promise((resolve, reject) => {
            });
            this._pending_pipe_msg.set(key, null);
            this._pipe.send("aaa", name, data);
            // fixme => implement
        };
        this.FolderAPI = (ticket) => {
            return new folder_1.FolderAPI(this._opts.base_url, ticket);
        };
        this.SockdAPI = (room) => {
            let rs = this._sockd_rooms.get(room);
            if (!rs) {
                rs = new room_1.SockdRoom(this._sockd, room);
                this._sockd_rooms.set(room, rs);
            }
            return rs;
        };
        window["debug_env"] = this; // only for debug remove this 
        this._opts = opts;
        this._sockd_rooms = new Map();
        this._pending_pipe_msg = new Map();
        this._pipe = opts.pipe;
        this._startup_payload = opts.startup_payload;
        this._fetch = fetch_1.actionFetch(`${opts.base_url}engine/${opts.plug}/${opts.agent}/exec_con`, opts.token);
        const sockdUrl = `${opts.base_url}engine/${opts.plug}/${opts.agent}/exec_ws`;
        this._sockd = new sockd_1.Sockd(sockdUrl);
        this._sockd.OnSockdMessage((msg) => {
            if (!msg.room) {
                console.log("no room message", msg);
                return;
            }
            if (msg.room === "plugs_dev") {
                console.log("PLUG DEBUG =>", msg.payload);
                return;
            }
            const room = this._sockd_rooms.get(msg.room);
            if (!room) {
                console.log("room without handler =>");
                return;
            }
            room.ProcessMessage(msg);
        });
    }
}
exports.Env = Env;
