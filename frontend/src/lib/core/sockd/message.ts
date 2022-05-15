const MESSAGE_SERVER_DIRECT = "server_direct"
const MESSAGE_SERVER_BROADCAST = "server_broadcast"
const MESSAGE_SERVER_PUBLISH = "server_publish"
const MESSAGE_PEER_DIRECT = "peer_direct"
const MESSAGE_PEER_BROADCAST = "peer_broadcast"
const MESSAGE_PEER_PUBLISH = "peer_publish"

export type SockdHandler = (message: SockdMessage) => void

export interface SockdMessage {
    type: string
    xid: string
    room?: string
    from_id?: string
    server_ident?: string
    ticket?: string
    targets?: string[]
    payload: any
}

export {
    MESSAGE_SERVER_DIRECT,
    MESSAGE_SERVER_BROADCAST,
    MESSAGE_SERVER_PUBLISH,
    MESSAGE_PEER_DIRECT,
    MESSAGE_PEER_BROADCAST,
    MESSAGE_PEER_PUBLISH
}
