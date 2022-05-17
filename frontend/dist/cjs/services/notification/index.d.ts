import { Readable, Writable } from "svelte/store";
import type { BasicAPI } from "../../lib/api/impl";
import type { SockdMessage } from "../../lib/sockd/stypes";
import type { SockdService } from "../sockd";
export interface Options {
    basicAPI: BasicAPI;
    sockdMuxer: SockdService;
}
interface State {
    messages: object[];
    loading: boolean;
    cursor: number;
}
export declare class Notification {
    basicAPI: BasicAPI;
    state: Writable<State>;
    isPendingRead: Readable<boolean>;
    is_open: Writable<boolean>;
    constructor(opts: Options);
    handler: (message: SockdMessage) => void;
    init: () => Promise<void>;
    set_read_notifications: (id: number) => Promise<void>;
    delete_notification: (id: number) => Promise<void>;
    toggle_notification: () => void;
}
export {};
