import type { Pipe } from "../../../../services/engine/pipe"
import { FolderAPI } from "../../../api/folder"
import { Sockd } from "../../sockd"
import type { SockdMessage } from "../../sockd/message"
import { SockdRoom } from "../../sockd/room"
import { actionFetch } from "./fetch"

export interface EnvOptions {
    token: string
    plug: string
    agent: string
    base_url: string
    parent_secret?: string
    pipe: Pipe
    startup_payload?: any
}

export interface ActionResponse {
    status_ok: boolean
    content_type?: string
    body: any
}

export class Environment {
    _opts: EnvOptions // only for debug remove this 
    _sockd: Sockd
    _sockd_rooms: Map<string, SockdRoom>
    _fetch: (name: string, data: string) => Promise<Response>

    _startup_payload?: any
    _pipe: Pipe
    _pending_pipe_msg: Map<string, Promise<any>>

    constructor(opts: EnvOptions) {
        window["debug_env"] = this // only for debug remove this 

        this._opts = opts
        this._sockd_rooms = new Map()
        this._pending_pipe_msg = new Map()

        this._pipe = opts.pipe
        this._startup_payload = opts.startup_payload
        this._fetch = actionFetch( `${opts.base_url}engine/${opts.plug}/${opts.agent}/exec_con`, opts.token)
        const sockdUrl =`${opts.base_url}engine/${opts.plug}/${opts.agent}/exec_ws`
        this._sockd = new Sockd(sockdUrl)
        this._sockd.OnSockdMessage((msg: SockdMessage) => {
            if (!msg.room) {
                console.log("no room message", msg)
                return
            }
            if (msg.room === "plugs_dev") {
                console.log("PLUG DEBUG =>", msg.payload)
                return
            }
            const room = this._sockd_rooms.get(msg.room)
            if (!room) {
                console.log("room without handler =>")
                return
            }
            room.ProcessMessage(msg)
        })
    }

    init = async () => {
        await this._sockd.init()
    }

    PreformAction = async (name: string, data: any): Promise<ActionResponse> => {
        const encoded = JSON.stringify(data)
        try {
            const resp = await this._fetch(name, encoded);
            const ctype = resp.headers.get("Content-Type")

            if (resp.status !== 200) {
                const txt = await resp.text();
                return {
                    status_ok: false,
                    content_type: ctype,
                    body: txt,
                }
            }

            const respData = await resp.json()
            return {
                body: respData,
                content_type: ctype,
                status_ok: true,
            }
        } catch (error) {
            return {
                status_ok: false,
                body: error,
            }
        }
    }

    startup_payload = () => {
        return this._startup_payload
    }

    PreformParentAction = async (name: string, data: any) : Promise<any> => {

        const key = "fixme => generate"

        const p = new Promise((resolve, reject) => {


        })

        

        this._pending_pipe_msg.set(key, null)

        this._pipe.send("aaa", name, data)



        // fixme => implement
    }


    FolderAPI = (ticket: string): FolderAPI => {
        return new FolderAPI(this._opts.base_url, ticket)
    }

    SockdAPI = (room: string): SockdRoom => {
        let rs = this._sockd_rooms.get(room)
        if (!rs) {
            rs = new SockdRoom(this._sockd, room)
            this._sockd_rooms.set(room, rs)
        }
        return rs
    }
} 

