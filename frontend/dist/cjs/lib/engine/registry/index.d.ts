import type { Environment, Factory, FactoryOptions } from "../ecore/ecore";
export declare class Registry {
    _factories: Map<string, (opts: FactoryOptions) => void>;
    _watchers: Map<string, (() => void)[]>;
    _type_watchers: Map<string, ((factory: Factory) => void)[]>;
    constructor();
    RegisterFactory: (type: string, name: string, factory: Factory) => void;
    WatchLoad: (type: string, name: string, timeout: number) => Promise<void>;
    OnTypeLoad: (typename: string, callback: (factory: Factory) => void) => void;
    Get: (type: string, name: string) => Factory;
    GetAll: (type: string) => Factory[];
    InstanceAll: (type: string, opts: FactoryOptions) => void;
    Instance: (type: string, name: string, opts: FactoryOptions) => void;
}
export declare const initFactory: () => void;
export interface ExecFactoryOptions {
    exec_loader?: string;
    plug: string;
    agent: string;
    entry: string;
    env: Environment;
    target: HTMLElement;
    payload?: any;
}
export declare const startExecFactory: (opts: ExecFactoryOptions) => Promise<void>;
