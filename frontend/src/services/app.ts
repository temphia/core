import { Navigator } from "./navigator";
import { Notification } from "./notification"
import { ApiManager } from "./apm";
import { clear_authed_data } from "../lib/authstore";
import { DataGroupService } from "./dyn";
import { FolderAPI } from "../lib";
import { EngineService } from "./engine";


declare global {
    interface Window {
        showModal(c: string, p: string);
        closeModal()
    }
}

export interface Toaster {
    success(message: string): void
    error(message: string): void
}


interface AppOptions {
    url_base: string
    api_url: string
    tenant_id: string
    site_token: string
    user_claim: string

    simple_modal_open: any
    simple_modal_close: any
    toaster: Toaster
}

export class AppService {
    url_base: string
    tenant_id: string
    site_token: string
    user_claim: string
    api_url: string

    apm: ApiManager
    navigator: Navigator
    noti: Notification
    toaster: Toaster

    engine_service: EngineService

    _simple_modal_open: any
    _simple_modal_close: any
    _current_data_service: DataGroupService

    _cabinet_sources: string[]
    _dyn_sources: string[]
    _folder_tickets: Map<string, FolderAPI>
    _store_sources: string[]
    _quick_apps: object[]


    constructor(opts: AppOptions) {
        this.url_base = opts.url_base
        this.api_url = opts.api_url
        this.tenant_id = opts.tenant_id
        this.site_token = opts.site_token
        this.user_claim = opts.user_claim
        this.navigator = new Navigator(this.url_base)
        this.toaster = opts.toaster

        this._folder_tickets = new Map()

        this._simple_modal_open = opts.simple_modal_open
        this._simple_modal_close = opts.simple_modal_close

        console.log("TOASTER ====>", opts.toaster)

        window["debug_app_handle1"] = this
    }

    async init() {
        await this.build_api_manager(this.user_claim)

        this.noti = new Notification({
            basicAPI: this.apm.get_basic_api(),
            sockdMuxer: this.apm.get_sockd_muxer()
        })

        await this.noti.init()

        const eapi = await this.apm.get_engine_api();
        this.engine_service = new EngineService(eapi)
    }

    async build_api_manager(claim: string) {
        this.apm = new ApiManager({
            api_base_url: this.api_url,
            tenant_id: this.tenant_id,
            skip_url_modify: true,
            user_token: claim
        })
        await this.apm.init()
    }


    user_profile_image_link = (user_id: string) => {
        return `${this.url_base}/api/${this.tenant_id}/v1/user_profile_image/${user_id}`
    }


    get_base_url() {
        return this.api_url
    }

    log_out() {
        clear_authed_data(this.tenant_id) // fixme => actually use user group specific auth page
        location.pathname = "/auth"
    }

    get_data_service = async (source: string, group: string) => {
        if (this._current_data_service) {
            if (this._current_data_service.source === source && this._current_data_service.group === group) {
                return this._current_data_service
            }
            await this._current_data_service.close()
        }

        const dapi = await this.apm.get_dtable_api(source, group)
        this._current_data_service = new DataGroupService(source, group, dapi, this.engine_service)
        await this._current_data_service.init()
        return this._current_data_service
    }

    simple_modal_open = (compo: any, opts: any) => {
        this._simple_modal_open(compo, opts)
    }

    simple_modal_close = () => {
        this._simple_modal_close()
    }

    big_modal_open = (_compo, _props) => {
        window.showModal(_compo, _props)
    }

    big_modal_close = () => {
        window.closeModal()
    }


    get_dyn_sources = async () => {
        if (this._dyn_sources) {
            return this._dyn_sources
        }

        const bapi = this.apm.get_basic_api()
        const resp = await bapi.list_dgroup_sources()
        if (resp.status !== 200) {
            console.log("Err loading dyn sources", resp)
            return []
        }

        this._dyn_sources = resp.data;
        return resp.data;
    }

    get_cabinet_sources = async () => {
        if (this._cabinet_sources) {
            return this._cabinet_sources
        }
        const bapi = this.apm.get_basic_api()
        const resp = await bapi.list_cabinet_sources()
        if (resp.status !== 200) {
            console.log("Err loading cabinet sources", resp)
            return []
        }
        this._cabinet_sources = resp.data
        return resp.data
    }

    get_store_sources = async () => {
        if (this._store_sources) {
            return this._store_sources
        }
        const api = await this.apm.get_repo_api()

        const resp = await api.repo_sources()
        this._store_sources = resp.data;
        return resp.data
    }


    get_folder_api = async (source: string, folder: string) => {
        const key = `${source}__${folder}`
        if (!this._folder_tickets.has(key)) {
            const capi = await this.apm.get_cabinet_api(source)
            const fresp = await capi.get_folder_ticket(folder)
            this._folder_tickets.set(key, new FolderAPI(this.get_base_url(), fresp.data))
        }
        return this._folder_tickets.get(key)
    }

    get_quick_apps = async () => {
        if (this._quick_apps) {
            return this._quick_apps
        }

        // fixme => 
    }

    is_mobile = () => {
        return screen.width < 700
    }

    // route_promise = async (original: Promise<any>, err_msg: string, success_msg?: string) => {
    //     try {
    //         const resp = await original;
    //         if (success_msg) {
    //             this.toaster.success(success_msg)
    //         }
    //         return resp
    //     } catch (error) {
    //         this.toaster.error(err_msg)
    //     }

    // }

}