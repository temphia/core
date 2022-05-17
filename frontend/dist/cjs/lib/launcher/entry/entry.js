"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const pipe_1 = require("../../../services/engine/pipe");
const registry_1 = require("../../engine/registry");
const env_1 = require("../../engine/env");
console.log("loader init using...");
registry_1.initFactory();
window.addEventListener("load", async () => {
    const opts = window["__loader_options__"];
    if (!opts) {
        console.log("Loader Options not found");
        return;
    }
    console.log("iframe portal opts @=>", opts);
    const pipe = new pipe_1.IFramePipe(opts.parent_secret);
    const env = new env_1.Env({
        agent: opts.agent,
        plug: opts.plug,
        token: opts.token,
        base_url: opts.base_url,
        parent_secret: opts.parent_secret,
        pipe,
    });
    await env.init();
    pipe.send("", "env_loaded", {});
    registry_1.startExecFactory({
        plug: opts.plug,
        agent: opts.agent,
        entry: opts.entry,
        env: env,
        target: document.getElementById("plugroot"),
        exec_loader: opts.exec_loader,
        payload: null,
    });
}, false);
