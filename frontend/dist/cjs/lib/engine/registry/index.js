"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.startExecFactory = exports.initFactory = exports.Registry = void 0;
class Registry {
    constructor() {
        this.RegisterFactory = (type, name, factory) => {
            console.log(`START REGISTER FACTORY => type(${type}) name(${name})`);
            const key = [type, name].toString();
            this._factories.set(key, factory);
            const watchers = this._watchers.get(key);
            if (watchers) {
                console.log("Found watchers ", watchers);
                watchers.forEach((watcher) => watcher());
            }
            const typeWatchers = this._type_watchers.get(type);
            if (typeWatchers) {
                typeWatchers.forEach((f) => f(factory));
            }
            console.log(`END REGISTER FACTORY => type(${type}) name(${name})`);
        };
        this.WatchLoad = async (type, name, timeout) => {
            console.log("before Watching");
            const key = [type, name].toString();
            if (this._factories.has(key)) {
                console.log("found factories already");
                return Promise.resolve();
            }
            const p = new Promise((resolve, reject) => {
                console.log("making promise");
                let oldwatcher = this._watchers.get(key);
                if (!oldwatcher) {
                    oldwatcher = new Array(0);
                }
                oldwatcher.push(() => {
                    resolve();
                });
                this._watchers.set(key, oldwatcher);
                setTimeout(() => {
                    reject(`TimeOut loading type ${type} & name ${name}`);
                }, timeout);
            });
            return p;
        };
        this.OnTypeLoad = (typename, callback) => {
            let oldwatcher = this._type_watchers.get(typename);
            if (!oldwatcher) {
                oldwatcher = new Array(0);
            }
            oldwatcher.push(callback);
        };
        this.Get = (type, name) => {
            const key = [type, name].toString();
            return this._factories.get(key.toString());
        };
        this.GetAll = (type) => {
            const facts = Array(0);
            this._factories.forEach((fact, [_type, _]) => {
                if (type !== _type) {
                    return;
                }
                facts.push(fact);
            });
            return facts;
        };
        this.InstanceAll = (type, opts) => {
            this._factories.forEach((fact, key) => {
                const [_type, _] = key.split(',');
                if (type !== _type) {
                    return;
                }
                fact(opts);
            });
        };
        this.Instance = (type, name, opts) => {
            const key = [type, name].toString();
            this._factories.get(key)(opts);
        };
        this._factories = new Map();
        this._watchers = new Map();
        this._type_watchers = new Map();
    }
}
exports.Registry = Registry;
exports.initFactory = () => {
    if (window["__registry__"]) {
        console.warn("Registry already loaded, skipping...");
        return;
    }
    const r = new Registry();
    r.RegisterFactory("loader.factory", "std.loader", async (opts) => {
        await opts.registry.WatchLoad("plug.factory", opts.entry, 2000);
        const factory = opts.registry.Get("plug.factory", opts.entry);
        if (!factory) {
            console.warn("could not load plug factory");
            return;
        }
    });
    console.log("GLOBAL_REGISTRY =>", r);
    window["__registry__"] = r;
    window["__register_factory__"] = r.RegisterFactory;
};
exports.startExecFactory = async (opts) => {
    console.log("Before starting factory", opts);
    const registry = window["__registry__"];
    if (!registry) {
        console.warn("registry not found");
        return;
    }
    if (!opts.exec_loader) {
        opts.exec_loader = "std.loader";
    }
    try {
        await registry.WatchLoad("loader.factory", opts.exec_loader, 10000);
    }
    catch (error) {
        console.warn("could not load, error occured:", error);
        return;
    }
    const plugFactory = registry.Get("loader.factory", opts.exec_loader);
    if (!opts.target) {
        opts.target = document.body;
    }
    plugFactory({
        plug: opts.plug,
        agent: opts.agent,
        entry: opts.entry,
        env: opts.env,
        registry: registry,
        target: opts.target,
        payload: opts.payload
    });
};
