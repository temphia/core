import type { PipeHandler, Pipe } from "../../lib/engine";
export declare class IFramePipe implements Pipe {
    _secret: string;
    _handlers: Set<PipeHandler>;
    constructor(secret: string);
    set_handler: (fn: PipeHandler) => void;
    remove_handler: (fn: PipeHandler) => void;
    send: (xid: string, action: string, data: any) => void;
}
