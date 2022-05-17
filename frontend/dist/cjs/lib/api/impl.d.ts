import { ApiBase } from "./base";
export declare class BprintAPI extends ApiBase {
    constructor(url: string, user_token: string);
    bprint_list(): Promise<import("axios").AxiosResponse<any>>;
    bprint_create(data: any): Promise<import("axios").AxiosResponse<any>>;
    bprint_get(id: string): Promise<import("axios").AxiosResponse<any>>;
    bprint_update(id: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    bprint_remove(id: string): Promise<import("axios").AxiosResponse<any>>;
    bprint_install(id: string, opts: any): Promise<import("axios").AxiosResponse<any>>;
    bprint_list_files(id: string): Promise<import("axios").AxiosResponse<any>>;
    bprint_get_file(id: string, file: string): Promise<import("axios").AxiosResponse<any>>;
    bprint_new_file(id: string, file: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    bprint_update_file(id: string, file: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    bprint_del_file(id: string, file: string): Promise<import("axios").AxiosResponse<any>>;
    bprint_import(data: any): Promise<import("axios").AxiosResponse<any>>;
}
export declare class UserAPI extends ApiBase {
    constructor(url: string, user_token: string);
    list_users(group?: string): Promise<import("axios").AxiosResponse<any>>;
    add_user(data: any): Promise<import("axios").AxiosResponse<any>>;
    get_user_by_id(id: string): Promise<import("axios").AxiosResponse<any>>;
    update_user(id: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    remove_user(id: string): Promise<import("axios").AxiosResponse<any>>;
    list_user_group(): Promise<import("axios").AxiosResponse<any>>;
    add_user_group(data: any): Promise<import("axios").AxiosResponse<any>>;
    get_user_group(gid: string): Promise<import("axios").AxiosResponse<any>>;
    update_user_group(gid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    remove_user_group(gid: string): Promise<import("axios").AxiosResponse<any>>;
    user_group_list_auth(gid: string): Promise<import("axios").AxiosResponse<any>>;
    user_group_add_auth(gid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    user_group_get_auth(gid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
    user_group_update_auth(gid: string, id: number, data: any): Promise<import("axios").AxiosResponse<any>>;
    user_group_remove_auth(gid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
    user_group_list_hook(gid: string): Promise<import("axios").AxiosResponse<any>>;
    user_group_add_hook(gid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    user_group_get_hook(gid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
    user_group_update_hook(gid: string, id: number, data: any): Promise<import("axios").AxiosResponse<any>>;
    user_group_remove_hook(gid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
    user_group_list_plug(gid: string): Promise<import("axios").AxiosResponse<any>>;
    user_group_add_plug(gid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    user_group_get_plug(gid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
    user_group_update_plug(gid: string, id: number, data: any): Promise<import("axios").AxiosResponse<any>>;
    user_group_remove_plug(gid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
    user_group_list_data(gid: string): Promise<import("axios").AxiosResponse<any>>;
    user_group_add_data(gid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    user_group_get_data(gid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
    user_group_update_data(gid: string, id: number, data: any): Promise<import("axios").AxiosResponse<any>>;
    user_group_remove_data(gid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
}
export declare class PlugAPI extends ApiBase {
    constructor(url: string, user_token: string);
    list_plug(): Promise<import("axios").AxiosResponse<any>>;
    new_plug(data: string): Promise<import("axios").AxiosResponse<any>>;
    update_plug(id: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    get_plug(pid: string): Promise<import("axios").AxiosResponse<any>>;
    del_plug(pid: string): Promise<import("axios").AxiosResponse<any>>;
    list_agent(pid: string): Promise<import("axios").AxiosResponse<any>>;
    new_agent(pid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    update_agent(pid: string, aid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    get_agent(pid: string, aid: string): Promise<import("axios").AxiosResponse<any>>;
    del_agent(pid: string, aid: string): Promise<import("axios").AxiosResponse<any>>;
}
export declare class CabinetAPI extends ApiBase {
    constructor(url: string, user_token: string, source: string);
    list_root(): Promise<import("axios").AxiosResponse<any>>;
    list_folder(folder: string): Promise<import("axios").AxiosResponse<any>>;
    new_folder(folder: string): Promise<import("axios").AxiosResponse<any>>;
    get_file(folder: string, file: string): Promise<import("axios").AxiosResponse<any>>;
    upload_file(folder: string, file: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    delete_file(folder: string, file: string): Promise<import("axios").AxiosResponse<any>>;
    get_folder_ticket(folder: string): Promise<import("axios").AxiosResponse<any>>;
}
export declare class DtableAPI extends ApiBase {
    constructor(url: string, user_token: string, source: string, group: string);
    load_group(): Promise<import("axios").AxiosResponse<any>>;
    list_tables(): Promise<import("axios").AxiosResponse<any>>;
    add_table(data: any): Promise<import("axios").AxiosResponse<any>>;
    edit_table(tid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    get_table(tid: string): Promise<import("axios").AxiosResponse<any>>;
    delete_table(tid: string): Promise<import("axios").AxiosResponse<any>>;
    list_columns(tid: string): Promise<import("axios").AxiosResponse<any>>;
    add_column(tid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    get_column(tid: string, cid: string): Promise<import("axios").AxiosResponse<any>>;
    edit_column(tid: string, cid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    delete_column(tid: string, cid: string): Promise<import("axios").AxiosResponse<any>>;
    list_view(tid: string): Promise<import("axios").AxiosResponse<any>>;
    new_view(tid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    modify_view(tid: string, id: number, data: any): Promise<import("axios").AxiosResponse<any>>;
    get_view(tid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
    del_view(tid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
    list_hook(tid: string): Promise<import("axios").AxiosResponse<any>>;
    new_hook(tid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    modify_hook(tid: string, id: number, data: any): Promise<import("axios").AxiosResponse<any>>;
    get_hook(tid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
    del_hook(tid: string, id: number): Promise<import("axios").AxiosResponse<any>>;
    new_row(tid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    get_row(tid: string, rid: number): Promise<import("axios").AxiosResponse<any>>;
    update_row(tid: string, rid: number, data: any): Promise<import("axios").AxiosResponse<any>>;
    delete_row(tid: string, rid: number): Promise<import("axios").AxiosResponse<any>>;
    simple_query(tid: string, data?: any): Promise<import("axios").AxiosResponse<any>>;
    fts_query(tid: string, str: string): Promise<import("axios").AxiosResponse<any>>;
    ref_load(tid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    ref_resolve(tid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    rev_ref_load(tid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    list_activity(tid: string, rowid: number): Promise<import("axios").AxiosResponse<any>>;
    comment_row(tid: string, rowid: number, msg: string): Promise<import("axios").AxiosResponse<any>>;
}
export declare class DynAPI extends ApiBase {
    constructor(url: string, user_token: string);
    list_group(source: string): Promise<import("axios").AxiosResponse<any>>;
    get_group(source: string, group: string): Promise<import("axios").AxiosResponse<any>>;
    new_group(source: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    edit_group(source: string, gid: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    delete_group(source: string, gid: string): Promise<import("axios").AxiosResponse<any>>;
}
export declare class BasicAPI extends ApiBase {
    constructor(url: string, user_token: string);
    list_cabinet_sources(): Promise<import("axios").AxiosResponse<any>>;
    list_dgroup_sources(): Promise<import("axios").AxiosResponse<any>>;
    message_user(data: any): Promise<import("axios").AxiosResponse<any>>;
    get_user_info(userid: string): Promise<import("axios").AxiosResponse<any>>;
    get_self_info(): Promise<import("axios").AxiosResponse<any>>;
    update_self_info(data: any): Promise<import("axios").AxiosResponse<any>>;
    self_change_email(data: any): Promise<import("axios").AxiosResponse<any>>;
    self_change_auth(data: any): Promise<import("axios").AxiosResponse<any>>;
    list_messages(data: any): Promise<import("axios").AxiosResponse<any>>;
    modify_messages(data: any): Promise<import("axios").AxiosResponse<any>>;
    dtable_change(data: any): Promise<import("axios").AxiosResponse<any>>;
    get_session_token(): string;
}
export declare class RepoAPI {
    basic_api: BasicAPI;
    constructor(bapi: BasicAPI);
    repo_sources(): Promise<import("axios").AxiosResponse<any>>;
    repo_list(source: string): Promise<import("axios").AxiosResponse<any>>;
    repo_get(source: string, group: string, slug: string): Promise<import("axios").AxiosResponse<any>>;
    repo_get_file(source: string, slug: string, file: string): Promise<import("axios").AxiosResponse<any>>;
}
export declare class ResourceAPI extends ApiBase {
    constructor(url: string, user_token: string);
    agent_resources_list(data: any): Promise<import("axios").AxiosResponse<any>>;
    resource_list(): Promise<import("axios").AxiosResponse<any>>;
    resource_create(data: any): Promise<import("axios").AxiosResponse<any>>;
    resource_get(slug: string): Promise<import("axios").AxiosResponse<any>>;
    resource_update(slug: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    resource_remove(slug: string): Promise<import("axios").AxiosResponse<any>>;
}
export declare class EngineAPI extends ApiBase {
    constructor(url: string, user_token: string);
    launcher_json(plug: string, agent: string, data: any): Promise<import("axios").AxiosResponse<any>>;
    referer_ticket(plug: string, agent: string): Promise<import("axios").AxiosResponse<any>>;
}
