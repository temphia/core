"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SockdService = void 0;
const room_1 = require("../../lib/sockd/room");
const SOCKD_NOTIFICATION_ROOM = "sys.users";
const SOCKD_DTABLE_ROOM = "sys.dtable";
class SockdService {
    constructor(sockd) {
        this.handle = (msg) => {
            switch (msg.room) {
                case SOCKD_NOTIFICATION_ROOM:
                    this._noti_room.ProcessMessage(msg);
                    break;
                case SOCKD_DTABLE_ROOM:
                    this._dtable_room.ProcessMessage(msg);
                    break;
                default:
                    console.log("Room not found", msg);
                    break;
            }
        };
        this.get_notification_room = () => {
            return this._noti_room;
        };
        this.get_dyn_room = () => {
            return this._dtable_room;
        };
        this.change_group = async (source, group, ticket) => {
            return this._basicAPi.dtable_change({
                group,
                source,
                ticket,
            });
        };
        this._sockd = sockd;
        this._noti_room = new room_1.SockdRoom(sockd, SOCKD_NOTIFICATION_ROOM);
        this._dtable_room = new room_1.SockdRoom(sockd, SOCKD_DTABLE_ROOM);
        sockd.OnSockdMessage(this.handle);
    }
}
exports.SockdService = SockdService;
