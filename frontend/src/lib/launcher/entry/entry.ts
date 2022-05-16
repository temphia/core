import { IFramePipe } from "../../../services/engine/pipe";
import type { LoaderOptions } from "../../engine/ecore";
import { initFactory, startExecFactory } from "../../engine/registry";
import { Env } from "../../engine/env";

console.log("loader init using...");
initFactory();

window.addEventListener(
  "load",
  async () => {
    const opts = window["__loader_options__"] as LoaderOptions;
    if (!opts) {
      console.log("Loader Options not found");
      return;
    }

    console.log("iframe portal opts @=>", opts);

    const pipe = new IFramePipe(opts.parent_secret);

    const env = new Env({
      agent: opts.agent,
      plug: opts.plug,
      token: opts.token,
      base_url: opts.base_url,
      parent_secret: opts.parent_secret,
      pipe,
    });

    await env.init();

    pipe.send("", "env_loaded", {});

    startExecFactory({
      plug: opts.plug,
      agent: opts.agent,
      entry: opts.entry,
      env: env,
      target: document.getElementById("plugroot"),
      exec_loader: opts.exec_loader,
      payload: null,
    });
  },
  false
);
