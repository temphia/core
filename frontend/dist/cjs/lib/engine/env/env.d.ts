import { FolderAPI } from "../../api/folder";
import type { ISockd } from "../../sockd/stypes";
import { SockdRoom } from "../../sockd/room";
import type { ActionResponse, Environment, Pipe } from "../ecore/index";
export interface EnvOptions {
    token: string;
    plug: string;
    agent: string;
    base_url: string;
    parent_secret?: string;
    pipe: Pipe;
    startup_payload?: any;
}
export declare class Env implements Environment {
    _opts: EnvOptions;
    _sockd: ISockd;
    _sockd_rooms: Map<string, SockdRoom>;
    _fetch: (name: string, data: string) => Promise<Response>;
    _startup_payload?: any;
    _pipe: Pipe;
    _pending_pipe_msg: Map<string, Promise<any>>;
    constructor(opts: EnvOptions);
    init: () => Promise<void>;
    PreformAction: (name: string, data: any) => Promise<ActionResponse>;
    startup_payload: () => any;
    PreformParentAction: (name: string, data: any) => Promise<any>;
    FolderAPI: (ticket: string) => FolderAPI;
    SockdAPI: (room: string) => SockdRoom;
}
