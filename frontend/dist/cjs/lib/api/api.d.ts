import { Sockd } from "../sockd";
import { DynAPI, RepoAPI, BasicAPI, DtableAPI, CabinetAPI, PlugAPI, UserAPI, BprintAPI, ResourceAPI, EngineAPI } from "./impl";
export interface Options {
    api_base_url: string;
    user_token: string;
    tenant_id: string;
    skip_url_modify?: boolean;
}
export declare class ApiBuilder {
    _api_base_url: string;
    _user_token: string;
    _tenant_id: string;
    _basic_api: BasicAPI;
    _admin_user_api: UserAPI;
    constructor(opts: Options);
    get_sockd_api(): Promise<Sockd>;
    get_dyn_api(): Promise<DynAPI>;
    get_basic_api(): Promise<BasicAPI>;
    get_dtable_api(source: string, group: string): Promise<DtableAPI>;
    get_cabinet_api(source: string): Promise<CabinetAPI>;
    get_plug_api(): Promise<PlugAPI>;
    get_user_api(): Promise<UserAPI>;
    get_repo_api(): RepoAPI;
    get_bprint_api(): Promise<BprintAPI>;
    get_resource_api(): Promise<ResourceAPI>;
    get_engine_api(): Promise<EngineAPI>;
}
