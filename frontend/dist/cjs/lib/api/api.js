"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ApiBuilder = void 0;
const sockd_1 = require("../sockd");
const impl_1 = require("./impl");
class ApiBuilder {
    constructor(opts) {
        this._api_base_url = opts.api_base_url;
        if (!opts.skip_url_modify) {
            this._api_base_url = `${opts.api_base_url}/${opts.tenant_id}/v1/`;
        }
        this._user_token = opts.user_token;
        this._tenant_id = opts.tenant_id;
    }
    async get_sockd_api() {
        const sockd = new sockd_1.Sockd(`${this._api_base_url}/self/user_ws?token=${this._basic_api.get_session_token()}`);
        await sockd.init();
        return sockd;
    }
    async get_dyn_api() {
        const sapi = new impl_1.DynAPI(this._api_base_url, this._user_token);
        await sapi.init();
        return sapi;
    }
    async get_basic_api() {
        const basic_api = new impl_1.BasicAPI(this._api_base_url, this._user_token);
        await basic_api.init();
        this._basic_api = basic_api;
        return basic_api;
    }
    async get_dtable_api(source, group) {
        const sapi = new impl_1.DtableAPI(this._api_base_url, this._user_token, source, group);
        await sapi.init();
        return sapi;
    }
    async get_cabinet_api(source) {
        const sapi = new impl_1.CabinetAPI(this._api_base_url, this._user_token, source);
        await sapi.init();
        return sapi;
    }
    async get_plug_api() {
        const sapi = new impl_1.PlugAPI(this._api_base_url, this._user_token);
        await sapi.init();
        return sapi;
    }
    async get_user_api() {
        const sapi = new impl_1.UserAPI(this._api_base_url, this._user_token);
        await sapi.init();
        return sapi;
    }
    get_repo_api() {
        const rapi = new impl_1.RepoAPI(this._basic_api);
        return rapi;
    }
    async get_bprint_api() {
        const bapi = new impl_1.BprintAPI(this._api_base_url, this._user_token);
        await bapi.init();
        return bapi;
    }
    async get_resource_api() {
        const rapi = new impl_1.ResourceAPI(this._api_base_url, this._user_token);
        await rapi.init();
        return rapi;
    }
    async get_engine_api() {
        const eapi = new impl_1.EngineAPI(this._api_base_url, this._user_token);
        await eapi.init();
        return eapi;
    }
}
exports.ApiBuilder = ApiBuilder;
