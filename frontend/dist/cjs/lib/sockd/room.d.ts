import type { SockdHandler, SockdMessage, ISockd, ISockdRoom } from "./stypes";
export declare class SockdRoom implements ISockdRoom {
    _socket: ISockd;
    _room: string;
    _onMessage?: SockdHandler;
    _onPeer?: SockdHandler;
    _onServer?: SockdHandler;
    constructor(socket: ISockd, room: string);
    SendDirect: (data: any) => void;
    SendBroadcast: (data: any) => void;
    SendTagged: (data: any, ticket: string, targets?: string[]) => void;
    onMessage: (handler: SockdHandler) => void;
    onPeer: (handler: SockdHandler) => void;
    onServer: (handler: SockdHandler) => void;
    ProcessMessage: (message: SockdMessage) => void;
    IsConnected: () => Promise<boolean>;
    LeaveRoom: () => void;
}
