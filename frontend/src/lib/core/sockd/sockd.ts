

import type { SockdMessage, SockdHandler } from "./message"
import { LinearBackoff, LRUBuffer, Websocket, WebsocketBuilder } from "../ws"

export class Sockd {
    _ws: Websocket
    _handler: (message: SockdMessage) => void
    _builder: WebsocketBuilder

    constructor(url: string) {
        console.log("CONNECTING WS @ ", url)
        this._builder = new WebsocketBuilder(url)
        this._builder.onMessage(this.handleIncoming)
        this._builder.withBackoff(new LinearBackoff(1, 3))
        this._builder.withBuffer(new LRUBuffer(20))
    }

    init = async () => {
        this._ws = this._builder.build()
    }

    private handleIncoming = (_: Websocket, ev: MessageEvent) => {
        // fixme => handle system messages

        const data = JSON.parse(ev.data)


        this._handler(data)
    }

    OnSockdMessage = (handler: (message: SockdMessage) => void): void => {
        this._handler = handler
    }

    SendSockd = (message: SockdMessage): void => {
        this._ws.send(JSON.stringify(message))
    }
}

