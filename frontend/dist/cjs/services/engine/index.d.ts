import { EngineAPI } from "../../lib";
export declare class EngineService {
    engine_api: EngineAPI;
    instances: Map<string, PlugExec>;
    constructor(eapi: EngineAPI);
    get_exec: (secret: string) => PlugExec;
    _on_message: (ev: any) => void;
    instance_stdplug: (plug: string, agent: string) => Promise<PlugExec>;
    instance_qapp: (qapp: object) => Promise<PlugExec>;
    instance_dataplug: (hook: object) => Promise<PlugExec>;
    instance: (plug: string, agent: string, exec_type: string, extra: object) => Promise<PlugExec>;
    clear_exec: (secretId: string) => void;
}
export interface ExecOptions {
    plug: string;
    agent: string;
    secret: string;
    exec_type: string;
    engine_data: object;
    exec_source?: object;
    parent: EngineService;
}
export declare class PlugExec {
    plug: string;
    agent: string;
    target: HTMLElement;
    secret: string;
    engine: EngineService;
    exec_type: string;
    itarget: HTMLIFrameElement;
    engine_data: object;
    message_handler?: (xid: string, action: string, data: any) => void;
    parent: EngineService;
    constructor(opts: ExecOptions);
    set_handler: (h: (xid: string, action: string, data: any) => void) => void;
    run: (target: HTMLElement, launch_data: object) => Promise<void>;
    on_message: (xid: string, action: string, data: any) => void;
    send_message: (data: any) => void;
    is_active: () => boolean;
    close: () => void;
}
