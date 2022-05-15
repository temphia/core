import { EngineAPI, generateId } from "../../lib";
import { iframeTemplateBuild } from "../../lib/launcher/entry/iframe_build";
import type { PipeMessage } from "./pipe";

const EXEC_TYPE_STD = "stdplug"
const EXEC_TYPE_DATA = "dataplug"

export class EngineService {
    engine_api: EngineAPI
    instances: Map<string, PlugExec>

    constructor(eapi: EngineAPI) {
        this.engine_api = eapi
        this.instances = new Map()
        window.addEventListener('message', this._on_message)
    }

    get_exec = (secret: string) => {
        return this.instances.get(secret)
    }

    _on_message = (ev) => {
        try {
            const decoded: PipeMessage = JSON.parse(ev.data);
            const exec = this.instances.get(decoded["parent_secret"])
            exec.on_message(decoded.xid, decoded.action, decoded.data)
        } catch (error) {
            console.log("engine interframe communication error", error)
        }
    }

    instance_stdplug = async (plug: string, agent: string) => {
        return this.instance(plug, agent, EXEC_TYPE_STD, {})
    }

    instance_qapp = async (qapp: object) => {
        return this.instance("", "", EXEC_TYPE_STD, qapp)
    }

    instance_dataplug = async (hook: object) => {
        return this.instance(hook["plug_id"], hook["agent_id"], EXEC_TYPE_DATA, hook)
    }

    instance = async (plug: string, agent: string, exec_type: string, extra: object) => {
        const { data, status } = await this.engine_api.launcher_json(plug, agent, {
            exec_type,
            exec_data: extra,
        })

        if (status !== 200) {
            console.warn("err loading", data)
            return
        }

        const secret = generateId()
        const exec = new PlugExec({
            agent: agent,
            engine_data: data,
            exec_type: exec_type,
            plug: plug,
            secret: secret,
            parent: this,
        })
        this.instances.set(secret, exec)
        return exec
    }

    clear_exec = (secretId: string) => {
        this.instances.delete(secretId)
    }
}

export interface ExecOptions {
    plug: string
    agent: string
    secret: string
    exec_type: string
    engine_data: object
    exec_source?: object
    parent: EngineService
}

// this opposite of wormhole (@parent side)

export class PlugExec {
    plug: string
    agent: string
    target: HTMLElement
    secret: string
    engine: EngineService
    exec_type: string
    itarget: HTMLIFrameElement
    engine_data: object
    message_handler?: (xid: string, action: string, data: any) => void
    parent: EngineService

    constructor(opts: ExecOptions) {
        this.plug = opts.plug
        this.agent = opts.agent

        this.secret = opts.secret
        this.exec_type = opts.exec_type
        this.engine_data = opts.engine_data
        this.parent = opts.parent
    }

    set_handler = (h: (xid: string, action: string, data: any) => void) => {
        this.message_handler = h
    }

    run = async (target: HTMLElement, launch_data: object) => {
        this.itarget = document.createElement("iframe");
        target.appendChild(this.itarget)

        const src = iframeTemplateBuild({
            agent: this.agent,
            plug: this.plug,
            base_url: this.engine_data["base_url"],
            entry_name: this.engine_data["entry"],
            exec_loader: this.engine_data["exec_loader"],
            js_plug_script: this.engine_data["js_plug_script"],
            style_file: this.engine_data["style"],
            token: this.engine_data["token"] || "",
            ext_scripts: this.engine_data["ext_scripts"],
            parent_secret: this.secret,
        })

        this.itarget.setAttribute("srcdoc", src);
        this.itarget.style.height = "100%"
        this.itarget.style.width = "100%"
    }

    // it is only called by engine service
    on_message = (xid: string, action: string, data: any) => {
        console.log("EVENT =>", xid, action, data)

        // if (!this.message_handler) {
        //     return
        // }

        // console.log("EVENT =>", ev)
        // return

        // const decoded = JSON.parse(ev.data);
        // if (decoded.parent_secret !== this.secret) {
        //     console.log("wrong parent token")
        //     return
        // }
        // console.log("ON_MESSAGE@PARENT", decoded)
        // this.message_handler(decoded["data"])
    }

    send_message = (data: any) => {
        const _data = JSON.stringify(data)
        this.itarget.contentWindow.postMessage(data, '*')
    }

    is_active = () => {
        return !!this.itarget
    }

    close = () => {
        if (this.itarget) {
            this.itarget.remove()
        }

        this.parent.clear_exec(this.secret)
        this.message_handler = null
        this.parent = null
    }
}