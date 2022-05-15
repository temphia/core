export interface PipeMessage {
    action?: string
    xid: string
    data: any
    parent_secret?: string
}

export type PipeHandler = (xid: string, action: string, data: any) => {}

export interface Pipe {
    send(xid: string, action: string, data: any): void
    set_handler(fn: PipeHandler): void
    remove_handler(fn: PipeHandler): void
}

export class IFramePipe {
    _secret: string
    _handlers: Set<PipeHandler>

    constructor(secret: string) {
        this._secret = secret
        this._handlers = new Set()

        window.addEventListener('message', (ev) => {
            const decoded: PipeMessage = JSON.parse(ev.data);
            this._handlers.forEach((fn) => fn(decoded.xid, decoded.action, decoded.data))
        })
    }

    set_handler = (fn: PipeHandler) => {
        this._handlers.add(fn)
    }

    remove_handler = (fn: PipeHandler) => {
        this._handlers.delete(fn)
    }

    send = (xid: string, action: string, data: any) => {
        const message = JSON.stringify({
            xid,
            data,
            action,
            parent_secret: this._secret,
        });
        window.parent.postMessage(message, '*');
    }
}