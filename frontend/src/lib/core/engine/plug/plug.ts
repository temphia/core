import type { FolderAPI } from "../../../api/folder";
import type { SockdRoom } from "../../sockd/room";
import type { ActionResponse } from "../env";
import type { Registry } from "../registry";

export const registerPlugFactory = (entryName: string, factory: (opts: any) => void) => registerFactory("plug.factory", entryName, factory)
export const registerExecLoaderFactory = (name: string, factory: (opts: any) => void) => registerFactory("loader.factory", name, factory)
export const registerFactory = (ftype: string, name: string, factory: (opts: any) => void) => {
    const pf = window["__register_factory__"];
    if (!pf) {
        console.warn("factory registry not found");
        return;
    }
    pf(ftype, name, factory);
}



export interface Environment {
    PreformAction: (name: string, data: any) => Promise<ActionResponse>
    PreformParentAction: (name: string, data: any) => Promise<any>    
    FolderAPI: (ticket: string) => FolderAPI
    SockdAPI(room: string): SockdRoom
}

export interface FactoryOptions {
    plug: string
    agent: string
    entry: string
    env: Environment
    target: HTMLElement
    payload?: any
    registry: Registry
}
