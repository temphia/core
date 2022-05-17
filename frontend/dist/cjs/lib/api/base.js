"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ApiBase = void 0;
const axios_1 = require("axios");
class ApiBase {
    constructor(opts) {
        this._user_token = opts.user_token;
        this._api_base_url = opts.url;
        this._service_options = opts.service_opts;
        this._session_token = "";
        this._http = null;
        this._service_path = opts.path;
        this.intercept_request = this.intercept_request.bind(this);
        this.intercept_request_err = this.intercept_request_err.bind(this);
        this._raw_http = axios_1.default.create({
            baseURL: opts.url,
        });
    }
    async init() {
        let resp = await this.refresh_token();
        console.log("@@@@@", resp.data);
        this._service_resp_payload = resp.data["service_payload"] || null;
        this._session_token = resp.data.token;
        this._http = axios_1.default.create({
            headers: { "Authorization": this._session_token },
            baseURL: this._api_base_url,
        });
        this._http.interceptors.request.use(this.intercept_request, this.intercept_request_err);
    }
    async refresh_token() {
        return this._raw_http.post(`/auth/refresh`, {
            "user_token": this._user_token,
            "options": this._service_options,
            "path": this._service_path,
        });
    }
    intercept_request(config) {
        return config;
    }
    intercept_request_err(error) {
        // fixme => if error is 401, refresh the token
        return Promise.reject(error);
    }
    get(url, config) {
        return this._http.get(url, config);
    }
    post(url, data, config) {
        return this._http.post(url, data, config);
    }
    put(url, data, config) {
        return this._http.put(url, data, config);
    }
    patch(url, data, config) {
        return this._http.patch(url, data, config);
    }
    delete(url, config) {
        return this._http.delete(url, config);
    }
}
exports.ApiBase = ApiBase;
