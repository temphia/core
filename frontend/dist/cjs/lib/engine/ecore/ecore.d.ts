import type { FolderAPI } from "../../api/folder";
import type { SockdRoom } from "../../sockd/room";
export declare const MODE_IFRAME = "IFRAME";
export declare const MODE_RAW_DOM = "RAW_DOM";
export declare const MODE_SUB_ORIGIN = "SUB_ORIGIN";
export interface LoaderOptions {
    token: string;
    entry: string;
    exec_loader: string;
    plug: string;
    agent: string;
    base_url: string;
    parent_secret?: string;
    startup_payload?: any;
}
export interface ActionResponse {
    status_ok: boolean;
    content_type?: string;
    body: any;
}
export interface Environment {
    PreformAction: (name: string, data: any) => Promise<ActionResponse>;
    PreformParentAction: (name: string, data: any) => Promise<any>;
    FolderAPI: (ticket: string) => FolderAPI;
    SockdAPI(room: string): SockdRoom;
}
export interface FactoryOptions {
    plug: string;
    agent: string;
    entry: string;
    env: Environment;
    target: HTMLElement;
    payload?: any;
    registry: any;
}
export declare type Factory = (opts: FactoryOptions) => void;
export interface PipeMessage {
    action?: string;
    xid: string;
    data: any;
    parent_secret?: string;
}
export declare type PipeHandler = (xid: string, action: string, data: any) => {};
export interface Pipe {
    send(xid: string, action: string, data: any): void;
    set_handler(fn: PipeHandler): void;
    remove_handler(fn: PipeHandler): void;
}
