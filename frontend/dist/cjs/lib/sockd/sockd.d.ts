import type { SockdMessage, SockdHandler, ISockd } from "./stypes";
import { Websocket, WebsocketBuilder } from "../core/ws";
export declare class Sockd implements ISockd {
    _ws: Websocket;
    _handler: (message: SockdMessage) => void;
    _builder: WebsocketBuilder;
    constructor(url: string);
    init: () => Promise<void>;
    private handleIncoming;
    OnSockdMessage: (h: SockdHandler) => void;
    SendSockd: (message: SockdMessage) => void;
}
