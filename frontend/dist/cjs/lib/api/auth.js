"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Authenticator = void 0;
const axios_1 = require("axios");
class Authenticator {
    constructor(opts) {
        this._site_token = opts.site_token;
        this._tenant_id = opts.tenant_id;
        this._api_base_url = opts.api_base_url;
        this._http_client = axios_1.default.create({
            baseURL: this._api_base_url,
        });
    }
    async LoginWithPassword(user_ident, password) {
        let resp = await this._http_client.post("/auth/login", {
            tenant_id: this._tenant_id,
            user_idendity: user_ident,
            password: password,
            site_token: this._site_token
        });
        if (resp.status == 200) {
            return {
                status_ok: true,
                user_token: resp.data["token"] || "",
            };
        }
        return {
            message: resp.data.message,
            status_ok: false,
        };
    }
}
exports.Authenticator = Authenticator;
