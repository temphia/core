"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PlugExec = exports.EngineService = void 0;
const lib_1 = require("../../lib");
const iframe_build_1 = require("../../lib/launcher/entry/iframe_build");
const EXEC_TYPE_STD = "stdplug";
const EXEC_TYPE_DATA = "dataplug";
class EngineService {
    constructor(eapi) {
        this.get_exec = (secret) => {
            return this.instances.get(secret);
        };
        this._on_message = (ev) => {
            try {
                const decoded = JSON.parse(ev.data);
                const exec = this.instances.get(decoded["parent_secret"]);
                exec.on_message(decoded.xid, decoded.action, decoded.data);
            }
            catch (error) {
                console.log("engine interframe communication error", error);
            }
        };
        this.instance_stdplug = async (plug, agent) => {
            return this.instance(plug, agent, EXEC_TYPE_STD, {});
        };
        this.instance_qapp = async (qapp) => {
            return this.instance("", "", EXEC_TYPE_STD, qapp);
        };
        this.instance_dataplug = async (hook) => {
            return this.instance(hook["plug_id"], hook["agent_id"], EXEC_TYPE_DATA, hook);
        };
        this.instance = async (plug, agent, exec_type, extra) => {
            const { data, status } = await this.engine_api.launcher_json(plug, agent, {
                exec_type,
                exec_data: extra,
            });
            if (status !== 200) {
                console.warn("err loading", data);
                return;
            }
            const secret = lib_1.generateId();
            const exec = new PlugExec({
                agent: agent,
                engine_data: data,
                exec_type: exec_type,
                plug: plug,
                secret: secret,
                parent: this,
            });
            this.instances.set(secret, exec);
            return exec;
        };
        this.clear_exec = (secretId) => {
            this.instances.delete(secretId);
        };
        this.engine_api = eapi;
        this.instances = new Map();
        window.addEventListener('message', this._on_message);
    }
}
exports.EngineService = EngineService;
// this opposite of wormhole (@parent side)
class PlugExec {
    constructor(opts) {
        this.set_handler = (h) => {
            this.message_handler = h;
        };
        this.run = async (target, launch_data) => {
            this.itarget = document.createElement("iframe");
            target.appendChild(this.itarget);
            const src = iframe_build_1.iframeTemplateBuild({
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
            });
            this.itarget.setAttribute("srcdoc", src);
            this.itarget.style.height = "100%";
            this.itarget.style.width = "100%";
        };
        // it is only called by engine service
        this.on_message = (xid, action, data) => {
            console.log("EVENT =>", xid, action, data);
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
        };
        this.send_message = (data) => {
            const _data = JSON.stringify(data);
            this.itarget.contentWindow.postMessage(data, '*');
        };
        this.is_active = () => {
            return !!this.itarget;
        };
        this.close = () => {
            if (this.itarget) {
                this.itarget.remove();
            }
            this.parent.clear_exec(this.secret);
            this.message_handler = null;
            this.parent = null;
        };
        this.plug = opts.plug;
        this.agent = opts.agent;
        this.secret = opts.secret;
        this.exec_type = opts.exec_type;
        this.engine_data = opts.engine_data;
        this.parent = opts.parent;
    }
}
exports.PlugExec = PlugExec;
