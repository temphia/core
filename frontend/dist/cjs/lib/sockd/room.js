"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SockdRoom = void 0;
const stypes_1 = require("./stypes");
class SockdRoom {
    constructor(socket, room) {
        this.SendDirect = (data) => {
            this._socket.SendSockd({
                payload: data,
                type: stypes_1.MESSAGE_PEER_DIRECT,
                xid: "",
                from_id: "",
                room: this._room,
            });
        };
        this.SendBroadcast = (data) => {
            this._socket.SendSockd({
                payload: data,
                type: stypes_1.MESSAGE_PEER_BROADCAST,
                xid: "",
                from_id: "",
                room: this._room,
            });
        };
        this.SendTagged = (data, ticket, targets) => {
            this._socket.SendSockd({
                payload: data,
                type: stypes_1.MESSAGE_PEER_PUBLISH,
                xid: "",
                from_id: "",
                room: this._room,
                targets: targets,
                ticket: ticket,
            });
        };
        this.onMessage = (handler) => {
            this._onMessage = handler;
        };
        this.onPeer = (handler) => {
            this._onPeer = handler;
        };
        this.onServer = (handler) => {
            this._onServer = handler;
        };
        this.ProcessMessage = (message) => {
            if (this._onMessage) {
                this._onMessage(message);
            }
            switch (message.type) {
                case stypes_1.MESSAGE_SERVER_DIRECT:
                    if (this._onServer) {
                        this._onServer(message);
                    }
                case stypes_1.MESSAGE_SERVER_BROADCAST:
                    if (this._onServer) {
                        this._onServer(message);
                    }
                case stypes_1.MESSAGE_SERVER_PUBLISH:
                    if (this._onServer) {
                        this._onServer(message);
                    }
                case stypes_1.MESSAGE_PEER_DIRECT:
                    if (this._onPeer) {
                        this._onPeer(message);
                    }
                case stypes_1.MESSAGE_PEER_BROADCAST:
                    if (this._onPeer) {
                        this._onPeer(message);
                    }
                case stypes_1.MESSAGE_PEER_PUBLISH:
                    if (this._onPeer) {
                        this._onPeer(message);
                    }
                default:
                    break;
            }
        };
        this.IsConnected = async () => {
            return false;
        };
        this.LeaveRoom = () => {
            // fixme => impl
        };
        this._socket = socket;
        this._room = room;
    }
}
exports.SockdRoom = SockdRoom;
