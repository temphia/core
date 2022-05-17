import { Navigator } from "./navigator";
import { Notification } from "./notification";
import { ApiManager } from "./apm";
import { DataGroupService } from "./dyn";
import { FolderAPI } from "../lib";
import { EngineService } from "./engine";
declare global {
    interface Window {
        showModal(c: string, p: string): any;
        closeModal(): any;
    }
}
export interface Toaster {
    success(message: string): void;
    error(message: string): void;
}
interface AppOptions {
    url_base: string;
    api_url: string;
    tenant_id: string;
    site_token: string;
    user_claim: string;
    simple_modal_open: any;
    simple_modal_close: any;
    toaster: Toaster;
}
export declare class AppService {
    url_base: string;
    tenant_id: string;
    site_token: string;
    user_claim: string;
    api_url: string;
    apm: ApiManager;
    navigator: Navigator;
    noti: Notification;
    toaster: Toaster;
    engine_service: EngineService;
    _simple_modal_open: any;
    _simple_modal_close: any;
    _current_data_service: DataGroupService;
    _cabinet_sources: string[];
    _dyn_sources: string[];
    _folder_tickets: Map<string, FolderAPI>;
    _store_sources: string[];
    _quick_apps: object[];
    constructor(opts: AppOptions);
    init(): Promise<void>;
    build_api_manager(claim: string): Promise<void>;
    user_profile_image_link: (user_id: string) => string;
    get_base_url(): string;
    log_out(): void;
    get_data_service: (source: string, group: string) => Promise<DataGroupService>;
    simple_modal_open: (compo: any, opts: any) => void;
    simple_modal_close: () => void;
    big_modal_open: (_compo: any, _props: any) => void;
    big_modal_close: () => void;
    get_dyn_sources: () => Promise<any>;
    get_cabinet_sources: () => Promise<any>;
    get_store_sources: () => Promise<any>;
    get_folder_api: (source: string, folder: string) => Promise<FolderAPI>;
    get_quick_apps: () => Promise<object[]>;
    is_mobile: () => boolean;
}
export {};
