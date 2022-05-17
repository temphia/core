"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ApiManager = void 0;
const api_1 = require("../lib/api");
const yootils_1 = require("yootils");
const sockd_1 = require("./sockd");
class ApiManager {
    constructor(opts) {
        this.init = async () => {
            this._basic_api = await this._api_builder.get_basic_api();
            this._sockd = await this._api_builder.get_sockd_api();
            await yootils_1.sleep(500);
            this._sockd_muxer = new sockd_1.SockdService(this._sockd);
        };
        this.get_sockd_muxer = () => {
            return this._sockd_muxer;
        };
        this._api_builder = new api_1.ApiBuilder(opts);
        this._dtable_apis = new Map();
        this._cabinet_apis = new Map();
    }
    async get_dyn_api() {
        if (!this._dyn_api) {
            this._dyn_api = await this._api_builder.get_dyn_api();
        }
        return this._dyn_api;
    }
    get_basic_api() {
        return this._basic_api;
    }
    async get_dtable_api(source, group) {
        const key = `${source}__${group}`;
        if (!this._dtable_apis.has(key)) {
            this._dtable_apis.set(key, await this._api_builder.get_dtable_api(source, group));
        }
        return this._dtable_apis.get(key);
    }
    async get_cabinet_api(source) {
        if (!this._cabinet_apis.has(source)) {
            this._cabinet_apis.set(source, await this._api_builder.get_cabinet_api(source));
        }
        return this._cabinet_apis.get(source);
    }
    async get_plug_api() {
        if (!this._plug_api) {
            this._plug_api = await this._api_builder.get_plug_api();
        }
        return this._plug_api;
    }
    async get_user_api() {
        if (!this._user_api) {
            this._user_api = await this._api_builder.get_user_api();
        }
        return this._user_api;
    }
    get_repo_api() {
        return this._api_builder.get_repo_api();
    }
    async get_bprint_api() {
        if (!this._bprint_api) {
            this._bprint_api = await this._api_builder.get_bprint_api();
        }
        return this._bprint_api;
    }
    async get_resource_api() {
        if (!this._repo_api) {
            this._resource_api = await this._api_builder.get_resource_api();
        }
        return this._resource_api;
    }
    async get_engine_api() {
        if (!this._engine_api) {
            this._engine_api = await this._api_builder.get_engine_api();
        }
        return this._engine_api;
    }
}
exports.ApiManager = ApiManager;
