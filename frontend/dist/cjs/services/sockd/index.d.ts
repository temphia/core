import type { Sockd } from "../../lib/sockd";
import type { SockdMessage } from "../../lib/sockd/stypes";
import { SockdRoom } from "../../lib/sockd/room";
import type { BasicAPI } from "../../lib/api/impl";
export declare class SockdService {
    _sockd: Sockd;
    _noti_room: SockdRoom;
    _dtable_room: SockdRoom;
    _basicAPi: BasicAPI;
    constructor(sockd: Sockd);
    handle: (msg: SockdMessage) => void;
    get_notification_room: () => SockdRoom;
    get_dyn_room: () => SockdRoom;
    change_group: (source: string, group: string, ticket: string) => Promise<import("axios").AxiosResponse<any>>;
}
