import { ApiBase } from "./base";

export class BprintAPI extends ApiBase {
    constructor(url: string, user_token: string) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin", "bprint"]
        })
    }

    async bprint_list() {
        return this.get("/bprint")
    }
    async bprint_create(data: any) {
        return this.post("/bprint", data)
    }
    async bprint_get(id: string) {
        return this.get(`/bprint/${id}`)
    }
    async bprint_update(id: string, data: any) {
        return this.post(`/bprint/${id}`, data)
    }
    async bprint_remove(id: string) {
        return this.delete(`/bprint/${id}`)
    }
    async bprint_install(id: string, opts: any) {
        return this.post(`/bprint/${id}/install`, opts)
    }
    async bprint_list_files(id: string) {
        return this.get(`/bprint/${id}/file`)
    }
    async bprint_get_file(id: string, file: string) {
        return this.get(`/bprint/${id}/file/${file}`)
    }
    async bprint_new_file(id: string, file: string, data: any) {
        return this.post(`/bprint/${id}/file/${file}`, data)
    }
    async bprint_update_file(id: string, file: string, data: any) {
        return this.patch(`/bprint/${id}/file/${file}`, data)
    }

    async bprint_del_file(id: string, file: string) {
        return this.delete(`/bprint/${id}/file/${file}`)
    }
    async bprint_import(data: any) {
        return this.post(`/import_bprint`, data)
    }

}


export class UserAPI extends ApiBase {
    constructor(url: string, user_token: string) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin", "user"]
        })
    }

    async list_users(group?: string) {
        return this.get(`/user${group ? `?user_group=` + group : ''}`)
    }
    async add_user(data: any) {
        return this.post(`/user`, data)
    }

    async get_user_by_id(id: string) {
        return this.get(`/user/${id}`)
    }
    async update_user(id: string, data: any) {
        return this.get(`/user/${id}`, data)
    }
    async remove_user(id: string) {
        return this.get(`/user/${id}`)
    }

    async list_user_group() {
        return this.get(`/user_group`)
    }
    async add_user_group(data: any) {
        return this.post(`/user_group`, data)
    }
    async get_user_group(gid: string) {
        return this.get(`/user_group/${gid}`)
    }

    async update_user_group(gid: string, data: any) {
        return this.post(`/user_group/${gid}`, data)
    }
    async remove_user_group(gid: string) {
        return this.delete(`/user_group/${gid}`)
    }

    // auth

    async user_group_list_auth(gid: string) {
        return this.get(`/user_auth/${gid}`)
    }

    async user_group_add_auth(gid: string, data: any) {
        return this.post(`/user_auth/${gid}`, data)
    }

    async user_group_get_auth(gid: string, id: number) {
        return this.get(`/user_auth/${gid}/${id}`)
    }

    async user_group_update_auth(gid: string, id: number, data: any) {
        return this.post(`/user_auth/${gid}/${id}`, data)

    }
    async user_group_remove_auth(gid: string, id: number) {
        return this.delete(`/user_auth/${gid}/${id}`)
    }

    // hook

    async user_group_list_hook(gid: string) {
        return this.get(`/user_hook/${gid}`)
    }

    async user_group_add_hook(gid: string, data: any) {
        return this.post(`/user_hook/${gid}`, data)
    }

    async user_group_get_hook(gid: string, id: number) {
        return this.get(`/user_hook/${gid}/${id}`)
    }

    async user_group_update_hook(gid: string, id: number, data: any) {
        return this.post(`/user_hook/${gid}/${id}`, data)
    }

    async user_group_remove_hook(gid: string, id: number) {
        return this.get(`/user_hook/${gid}/${id}`)
    }

    // plug

    async user_group_list_plug(gid: string) {
        return this.get(`/user_plug/${gid}`)
    }

    async user_group_add_plug(gid: string, data: any) {
        return this.post(`/user_plug/${gid}`, data)
    }

    async user_group_get_plug(gid: string, id: number) {
        return this.get(`/user_plug/${gid}/${id}`)
    }

    async user_group_update_plug(gid: string, id: number, data: any) {
        return this.post(`/user_plug/${gid}/${id}`, data)
    }

    async user_group_remove_plug(gid: string, id: number) {
        return this.get(`/user_plug/${gid}/${id}`)
    }

    // data

    async user_group_list_data(gid: string) {
        return this.get(`/user_data/${gid}`)
    }

    async user_group_add_data(gid: string, data: any) {
        return this.post(`/user_data/${gid}`, data)
    }

    async user_group_get_data(gid: string, id: number) {
        return this.get(`/user_data/${gid}/${id}`)
    }

    async user_group_update_data(gid: string, id: number, data: any) {
        return this.post(`/user_data/${gid}/${id}`, data)
    }

    async user_group_remove_data(gid: string, id: number) {
        return this.get(`/user_data/${gid}/${id}`)
    }



    // fixme => user perm stuff
}


export class PlugAPI extends ApiBase {
    constructor(url: string, user_token: string) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin", "plug"]
        })
    }

    async list_plug() {
        return this.get(`/plug`)
    }

    async new_plug(data: string) {
        return this.post(`/plug`, data)
    }

    async update_plug(id: string, data: any) {
        return this.post(`/plug/${id}`, data)
    }

    async get_plug(pid: string) {
        return this.get(`/plug/${pid}`)
    }
    async del_plug(pid: string) {
        return this.delete(`/plug/${pid}`)
    }

    async list_agent(pid: string) {
        return this.get(`/plug/${pid}/agent`)
    }

    async new_agent(pid: string, data: any) {
        return this.post(`/plug/${pid}/agent`, data)
    }

    async update_agent(pid: string, aid: string, data: any) {
        return this.post(`/plug/${pid}/agent/${aid}`, data)
    }

    async get_agent(pid: string, aid: string) {
        return this.get(`/plug/${pid}/agent/${aid}`)
    }
    async del_agent(pid: string, aid: string) {
        return this.delete(`/plug/${pid}/agent/${aid}`)
    }
}


export class CabinetAPI extends ApiBase {
    constructor(url: string, user_token: string, source: string) {
        super({
            url: url,
            user_token: user_token,
            path: ["cabinet", source]
        })
    }
    async list_root() {
        return this.get(`/cabinet`)
    }
    async list_folder(folder: string) {
        return this.get(`/cabinet/${folder}`)
    }
    async new_folder(folder: string) {
        return this.post(`/cabinet/${folder}`)
    }
    async get_file(folder: string, file: string) {
        return this.get(`/cabinet/${folder}/file/${file}`)
    }
    async upload_file(folder: string, file: string, data) {
        return this.post(`/cabinet/${folder}/file/${file}`, data)
    }
    async delete_file(folder: string, file: string) {
        return this.delete(`/cabinet/${folder}/file/${file}`)
    }

    async get_folder_ticket(folder: string) {
        return this.post(`/cabinet/${folder}/ticket`)
    }

}

export class DtableAPI extends ApiBase {
    constructor(url: string, user_token: string, source: string, group: string) {
        super({
            url: url,
            user_token: user_token,
            path: ["dtable", source, group]
        })
    }

    async load_group() {
        return this.get(`/dgroup_load`)
    }

    // dtable
    async list_tables() {
        return this.get(`/dtable`)
    }

    async add_table(data: any) {
        return this.post(`/dtable`, data)
    }

    async edit_table(tid: string, data: any) {
        return this.patch(`/dtable/${tid}`, data)
    }

    async get_table(tid: string) {
        return this.get(`/dtable/${tid}`)
    }

    async delete_table(tid: string) {
        return this.delete(`/dtable/${tid}`)
    }

    async list_columns(tid: string) {
        return this.get(`/dtable/${tid}/column`)
    }
    async add_column(tid: string, data: any) {
        return this.post(`/dtable/${tid}/column`, data)
    }

    async get_column(tid: string, cid: string) {
        return this.get(`/dtable/${tid}/column/${cid}`)
    }

    async edit_column(tid: string, cid: string, data: any) {
        return this.patch(`/dtable/${tid}/column/${cid}`, data)
    }

    async delete_column(tid: string, cid: string) {
        return this.delete(`/dtable/${tid}/column/${cid}`)
    }

    // view stuff

    async list_view(tid: string) {
        return this.get(`/dtable/${tid}/view`)
    }

    async new_view(tid: string, data: any) {
        return this.post(`/dtable/${tid}/view`, data)
    }

    async modify_view(tid: string, id: number, data: any) {
        return this.post(`/dtable/${tid}/view/${id}`, data)
    }

    async get_view(tid: string, id: number) {
        return this.get(`/dtable/${tid}/view/${id}`)
    }

    async del_view(tid: string, id: number) {
        return this.delete(`/dtable/${tid}/view/${id}`)
    }

    // hook stuff
    async list_hook(tid: string) {
        return this.get(`/dtable/${tid}/hook`)
    }

    async new_hook(tid: string, data: any) {
        return this.post(`/dtable/${tid}/hook`, data)
    }

    async modify_hook(tid: string, id: number, data: any) {
        return this.post(`/dtable/${tid}/hook/${id}`, data)
    }

    async get_hook(tid: string, id: number) {
        return this.get(`/dtable/${tid}/hook/${id}`)
    }

    async del_hook(tid: string, id: number) {
        return this.delete(`/dtable/${tid}/hook/${id}`)
    }

    // dtable ops

    async new_row(tid: string, data: any) {
        return this.post(`/dtable_ops/${tid}/row`, data)
    }
    async get_row(tid: string, rid: number) {
        return this.get(`/dtable_ops/${tid}/row/${rid}`)
    }
    async update_row(tid: string, rid: number, data: any) {
        return this.post(`/dtable_ops/${tid}/row/${rid}`, data)
    }
    async delete_row(tid: string, rid: number) {
        return this.delete(`/dtable_ops/${tid}/row/${rid}`)
    }
    async simple_query(tid: string, data?: any) {
        if (!data) {
            data = {}
        }
        return this.post(`/dtable_ops/${tid}/simple_query`, data)
    }

    async fts_query(tid: string, str: string) {
        return this.post(`/dtable_ops/${tid}/fts_query`, {
            "qstr": str
        })
    }

    async ref_load(tid: string, data: any) {
        return this.post(`/dtable_ops/${tid}/ref_load`, data)
    }

    async ref_resolve(tid: string, data: any) {
        return this.post(`/dtable_ops/${tid}/ref_resolve`, data)
    }

    async rev_ref_load(tid: string, data) {
        return this.post(`/dtable_ops/${tid}/rev_ref_load`, data)
    }

    async list_activity(tid: string, rowid: number) {
        return this.get(`/dtable_ops/${tid}/activity/${rowid}`)
    }

    async comment_row(tid: string, rowid: number, msg: string) {
        return this.post(`/dtable_ops/${tid}/activity/${rowid}`, {
            "message": msg,
        })
    }

}


export class DynAPI extends ApiBase {
    constructor(url: string, user_token: string) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin"]
        })
    }

    async list_group(source: string) {
        return this.get(`/dgroup/${source}`)
    }

    async get_group(source: string, group: string) {
        return this.get(`/dgroup/${source}/${group}`)
    }

    async new_group(source: string, data: any) {
        return this.post(`/dgroup/${source}`, data)
    }

    async edit_group(source: string, gid: string, data: any) {
        return this.patch(`/dgroup/${source}/${gid}`, data)
    }

    async delete_group(source: string, gid: string) {
        return this.delete(`/dgroup/${source}/${gid}`)
    }
}



export class BasicAPI extends ApiBase {
    constructor(url: string, user_token: string) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin"]
        })
    }

    async list_cabinet_sources() {
        return this.get(`/cabinet_sources`)
    }
    async list_dgroup_sources() {
        return this.get(`/dgroup`)
    }

    async message_user(data: any) {
        return this.post("/self/message_user", data)
    }

    async get_user_info(userid: string) {
        return this.get(`/self/get_user_info/${userid}`)
    }

    async get_self_info() {
        return this.get("/self/get_self_info")
    }

    async update_self_info(data: any) {
        return this.post("/self/get_self_info", data)
    }

    async self_change_email(data: any) {
        return this.post("/self/change_email", data)
    }

    async self_change_auth(data: any) {
        return this.post("/self/change_auth", data)
    }

    async list_messages(data: any) {
        return this.post("/self/list_messages", data)
    }

    async modify_messages(data: any) {
        return this.post("/self/modify_messages", data)
    }

    async dtable_change(data: any) {
        return this.post("/self/dtable_change", data)
    }

    get_session_token() {
        return this._session_token
    }
}

export class RepoAPI {
    basic_api: BasicAPI
    constructor(bapi: BasicAPI) {
        this.basic_api = bapi
    }
    async repo_sources() {
        return this.basic_api.get(`/repo`)
    }

    async repo_list(source: string) {
        return this.basic_api.get(`/repo/${source}`)
    }
    async repo_get(source: string, group: string, slug: string) {
        return this.basic_api.get(`/repo/${source}/${group}/${slug}`)
    }
    async repo_get_file(source: string, slug: string, file: string) {
        return this.basic_api.get(`/repo/${source}/${slug}/${file}`)
    }
}

export class ResourceAPI extends ApiBase {
    constructor(url: string, user_token: string) {
        super({
            url: url,
            user_token: user_token,
            path: ["resource"]
        })
    }

    async agent_resources_list(data: any) {
        return this.post("/agent_resources", data)
    }

    async resource_list() {
        return this.get("/resource")
    }

    async resource_create(data: any) {
        return this.post("/resource", data)
    }

    async resource_get(slug: string) {
        return this.get(`/resource/${slug}`)
    }

    async resource_update(slug: string, data: any) {
        return this.post(`/resource/${slug}`, data)
    }

    async resource_remove(slug: string) {
        return this.delete(`/resource/${slug}`)
    }
}


export class EngineAPI extends ApiBase {
    constructor(url: string, user_token: string) {
        super({
            url: url,
            user_token: user_token,
            path: ["engine"]
        })
    }

    async launcher_json(plug: string, agent: string, data: any) {
        return this.post(`/engine/${plug}/${agent}/launcher/json`, data)
    }

    async referer_ticket(plug: string, agent: string) {
        return this.get(`/engine/${plug}/${agent}/referer_ticket`)
    }
}