
// utils
export declare function _log(message: String ): void
export declare function _http_call(request: HTTPRequest): HTTPResponse

// syncer
export declare function _broadcast(room: string, payload: any): void
export declare function _send_sessions(session: string, payload: any ): void
export declare function _add_to_room(session:string, room: string): void
export declare function _kick_from_room(session:string, room: string): void
export declare function _list_room_session(room: string): void
export declare function _ban_session(room: string): void

// signaler
export declare function _emit_signal(signal: string, payload: any): any
export declare function _list_outgoing_signal(): SiganlType[]
export declare function _list_incomming_signal():  SiganlType[]

// state kv
export declare function _state_kv_tx_start(): number
export declare function _state_kv_get(tx: number, key: string): StateValue
export declare function _state_kv_set(tx: number, key: string, val: StateValue): any
export declare function _state_kv_list(tx: number): StateValue[] | null
export declare function _state_kv_listByAud(tx: number): StateValue[] | null
export declare function _state_kv_delete(tx: number): any
export declare function _state_kv_tx_finish(tx: number): any


export interface HTTPRequest {
    url: string;
    method: string;
    headers?: Map<string, string>;
    body?: any;
}

export interface HTTPResponse {
    status_code: number
    headers: Map<string, string>;
    body: string | Uint8Array
}

export interface SiganlType {
    slug: string
    plug_id: string
    agent_id: string
}

export interface StateValue {
    inner: string
    aud: string
}
