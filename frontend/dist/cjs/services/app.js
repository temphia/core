"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AppService = void 0;
const navigator_1 = require("./navigator");
const notification_1 = require("./notification");
const apm_1 = require("./apm");
const authstore_1 = require("../lib/authstore");
const dyn_1 = require("./dyn");
const lib_1 = require("../lib");
const engine_1 = require("./engine");
class AppService {
    constructor(opts) {
        this.user_profile_image_link = (user_id) => {
            return `${this.url_base}/api/${this.tenant_id}/v1/user_profile_image/${user_id}`;
        };
        this.get_data_service = async (source, group) => {
            if (this._current_data_service) {
                if (this._current_data_service.source === source && this._current_data_service.group === group) {
                    return this._current_data_service;
                }
                await this._current_data_service.close();
            }
            const dapi = await this.apm.get_dtable_api(source, group);
            this._current_data_service = new dyn_1.DataGroupService(source, group, dapi, this.engine_service);
            await this._current_data_service.init();
            return this._current_data_service;
        };
        this.simple_modal_open = (compo, opts) => {
            this._simple_modal_open(compo, opts);
        };
        this.simple_modal_close = () => {
            this._simple_modal_close();
        };
        this.big_modal_open = (_compo, _props) => {
            window.showModal(_compo, _props);
        };
        this.big_modal_close = () => {
            window.closeModal();
        };
        this.get_dyn_sources = async () => {
            if (this._dyn_sources) {
                return this._dyn_sources;
            }
            const bapi = this.apm.get_basic_api();
            const resp = await bapi.list_dgroup_sources();
            if (resp.status !== 200) {
                console.log("Err loading dyn sources", resp);
                return [];
            }
            this._dyn_sources = resp.data;
            return resp.data;
        };
        this.get_cabinet_sources = async () => {
            if (this._cabinet_sources) {
                return this._cabinet_sources;
            }
            const bapi = this.apm.get_basic_api();
            const resp = await bapi.list_cabinet_sources();
            if (resp.status !== 200) {
                console.log("Err loading cabinet sources", resp);
                return [];
            }
            this._cabinet_sources = resp.data;
            return resp.data;
        };
        this.get_store_sources = async () => {
            if (this._store_sources) {
                return this._store_sources;
            }
            const api = await this.apm.get_repo_api();
            const resp = await api.repo_sources();
            this._store_sources = resp.data;
            return resp.data;
        };
        this.get_folder_api = async (source, folder) => {
            const key = `${source}__${folder}`;
            if (!this._folder_tickets.has(key)) {
                const capi = await this.apm.get_cabinet_api(source);
                const fresp = await capi.get_folder_ticket(folder);
                this._folder_tickets.set(key, new lib_1.FolderAPI(this.get_base_url(), fresp.data));
            }
            return this._folder_tickets.get(key);
        };
        this.get_quick_apps = async () => {
            if (this._quick_apps) {
                return this._quick_apps;
            }
            // fixme => 
        };
        this.is_mobile = () => {
            return screen.width < 700;
        };
        this.url_base = opts.url_base;
        this.api_url = opts.api_url;
        this.tenant_id = opts.tenant_id;
        this.site_token = opts.site_token;
        this.user_claim = opts.user_claim;
        this.navigator = new navigator_1.Navigator(this.url_base);
        this.toaster = opts.toaster;
        this._folder_tickets = new Map();
        this._simple_modal_open = opts.simple_modal_open;
        this._simple_modal_close = opts.simple_modal_close;
        console.log("TOASTER ====>", opts.toaster);
        window["debug_app_handle1"] = this;
    }
    async init() {
        await this.build_api_manager(this.user_claim);
        this.noti = new notification_1.Notification({
            basicAPI: this.apm.get_basic_api(),
            sockdMuxer: this.apm.get_sockd_muxer()
        });
        await this.noti.init();
        const eapi = await this.apm.get_engine_api();
        this.engine_service = new engine_1.EngineService(eapi);
    }
    async build_api_manager(claim) {
        this.apm = new apm_1.ApiManager({
            api_base_url: this.api_url,
            tenant_id: this.tenant_id,
            skip_url_modify: true,
            user_token: claim
        });
        await this.apm.init();
    }
    get_base_url() {
        return this.api_url;
    }
    log_out() {
        authstore_1.clear_authed_data(this.tenant_id); // fixme => actually use user group specific auth page
        location.pathname = "/auth";
    }
}
exports.AppService = AppService;
