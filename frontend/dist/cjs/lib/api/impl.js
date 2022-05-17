"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.EngineAPI = exports.ResourceAPI = exports.RepoAPI = exports.BasicAPI = exports.DynAPI = exports.DtableAPI = exports.CabinetAPI = exports.PlugAPI = exports.UserAPI = exports.BprintAPI = void 0;
const base_1 = require("./base");
class BprintAPI extends base_1.ApiBase {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin", "bprint"]
        });
    }
    async bprint_list() {
        return this.get("/bprint");
    }
    async bprint_create(data) {
        return this.post("/bprint", data);
    }
    async bprint_get(id) {
        return this.get(`/bprint/${id}`);
    }
    async bprint_update(id, data) {
        return this.post(`/bprint/${id}`, data);
    }
    async bprint_remove(id) {
        return this.delete(`/bprint/${id}`);
    }
    async bprint_install(id, opts) {
        return this.post(`/bprint/${id}/install`, opts);
    }
    async bprint_list_files(id) {
        return this.get(`/bprint/${id}/file`);
    }
    async bprint_get_file(id, file) {
        return this.get(`/bprint/${id}/file/${file}`);
    }
    async bprint_new_file(id, file, data) {
        return this.post(`/bprint/${id}/file/${file}`, data);
    }
    async bprint_update_file(id, file, data) {
        return this.patch(`/bprint/${id}/file/${file}`, data);
    }
    async bprint_del_file(id, file) {
        return this.delete(`/bprint/${id}/file/${file}`);
    }
    async bprint_import(data) {
        return this.post(`/import_bprint`, data);
    }
}
exports.BprintAPI = BprintAPI;
class UserAPI extends base_1.ApiBase {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin", "user"]
        });
    }
    async list_users(group) {
        return this.get(`/user${group ? `?user_group=` + group : ''}`);
    }
    async add_user(data) {
        return this.post(`/user`, data);
    }
    async get_user_by_id(id) {
        return this.get(`/user/${id}`);
    }
    async update_user(id, data) {
        return this.get(`/user/${id}`, data);
    }
    async remove_user(id) {
        return this.get(`/user/${id}`);
    }
    async list_user_group() {
        return this.get(`/user_group`);
    }
    async add_user_group(data) {
        return this.post(`/user_group`, data);
    }
    async get_user_group(gid) {
        return this.get(`/user_group/${gid}`);
    }
    async update_user_group(gid, data) {
        return this.post(`/user_group/${gid}`, data);
    }
    async remove_user_group(gid) {
        return this.delete(`/user_group/${gid}`);
    }
    // auth
    async user_group_list_auth(gid) {
        return this.get(`/user_auth/${gid}`);
    }
    async user_group_add_auth(gid, data) {
        return this.post(`/user_auth/${gid}`, data);
    }
    async user_group_get_auth(gid, id) {
        return this.get(`/user_auth/${gid}/${id}`);
    }
    async user_group_update_auth(gid, id, data) {
        return this.post(`/user_auth/${gid}/${id}`, data);
    }
    async user_group_remove_auth(gid, id) {
        return this.delete(`/user_auth/${gid}/${id}`);
    }
    // hook
    async user_group_list_hook(gid) {
        return this.get(`/user_hook/${gid}`);
    }
    async user_group_add_hook(gid, data) {
        return this.post(`/user_hook/${gid}`, data);
    }
    async user_group_get_hook(gid, id) {
        return this.get(`/user_hook/${gid}/${id}`);
    }
    async user_group_update_hook(gid, id, data) {
        return this.post(`/user_hook/${gid}/${id}`, data);
    }
    async user_group_remove_hook(gid, id) {
        return this.get(`/user_hook/${gid}/${id}`);
    }
    // plug
    async user_group_list_plug(gid) {
        return this.get(`/user_plug/${gid}`);
    }
    async user_group_add_plug(gid, data) {
        return this.post(`/user_plug/${gid}`, data);
    }
    async user_group_get_plug(gid, id) {
        return this.get(`/user_plug/${gid}/${id}`);
    }
    async user_group_update_plug(gid, id, data) {
        return this.post(`/user_plug/${gid}/${id}`, data);
    }
    async user_group_remove_plug(gid, id) {
        return this.get(`/user_plug/${gid}/${id}`);
    }
    // data
    async user_group_list_data(gid) {
        return this.get(`/user_data/${gid}`);
    }
    async user_group_add_data(gid, data) {
        return this.post(`/user_data/${gid}`, data);
    }
    async user_group_get_data(gid, id) {
        return this.get(`/user_data/${gid}/${id}`);
    }
    async user_group_update_data(gid, id, data) {
        return this.post(`/user_data/${gid}/${id}`, data);
    }
    async user_group_remove_data(gid, id) {
        return this.get(`/user_data/${gid}/${id}`);
    }
}
exports.UserAPI = UserAPI;
class PlugAPI extends base_1.ApiBase {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin", "plug"]
        });
    }
    async list_plug() {
        return this.get(`/plug`);
    }
    async new_plug(data) {
        return this.post(`/plug`, data);
    }
    async update_plug(id, data) {
        return this.post(`/plug/${id}`, data);
    }
    async get_plug(pid) {
        return this.get(`/plug/${pid}`);
    }
    async del_plug(pid) {
        return this.delete(`/plug/${pid}`);
    }
    async list_agent(pid) {
        return this.get(`/plug/${pid}/agent`);
    }
    async new_agent(pid, data) {
        return this.post(`/plug/${pid}/agent`, data);
    }
    async update_agent(pid, aid, data) {
        return this.post(`/plug/${pid}/agent/${aid}`, data);
    }
    async get_agent(pid, aid) {
        return this.get(`/plug/${pid}/agent/${aid}`);
    }
    async del_agent(pid, aid) {
        return this.delete(`/plug/${pid}/agent/${aid}`);
    }
}
exports.PlugAPI = PlugAPI;
class CabinetAPI extends base_1.ApiBase {
    constructor(url, user_token, source) {
        super({
            url: url,
            user_token: user_token,
            path: ["cabinet", source]
        });
    }
    async list_root() {
        return this.get(`/cabinet`);
    }
    async list_folder(folder) {
        return this.get(`/cabinet/${folder}`);
    }
    async new_folder(folder) {
        return this.post(`/cabinet/${folder}`);
    }
    async get_file(folder, file) {
        return this.get(`/cabinet/${folder}/file/${file}`);
    }
    async upload_file(folder, file, data) {
        return this.post(`/cabinet/${folder}/file/${file}`, data);
    }
    async delete_file(folder, file) {
        return this.delete(`/cabinet/${folder}/file/${file}`);
    }
    async get_folder_ticket(folder) {
        return this.post(`/cabinet/${folder}/ticket`);
    }
}
exports.CabinetAPI = CabinetAPI;
class DtableAPI extends base_1.ApiBase {
    constructor(url, user_token, source, group) {
        super({
            url: url,
            user_token: user_token,
            path: ["dtable", source, group]
        });
    }
    async load_group() {
        return this.get(`/dgroup_load`);
    }
    // dtable
    async list_tables() {
        return this.get(`/dtable`);
    }
    async add_table(data) {
        return this.post(`/dtable`, data);
    }
    async edit_table(tid, data) {
        return this.patch(`/dtable/${tid}`, data);
    }
    async get_table(tid) {
        return this.get(`/dtable/${tid}`);
    }
    async delete_table(tid) {
        return this.delete(`/dtable/${tid}`);
    }
    async list_columns(tid) {
        return this.get(`/dtable/${tid}/column`);
    }
    async add_column(tid, data) {
        return this.post(`/dtable/${tid}/column`, data);
    }
    async get_column(tid, cid) {
        return this.get(`/dtable/${tid}/column/${cid}`);
    }
    async edit_column(tid, cid, data) {
        return this.patch(`/dtable/${tid}/column/${cid}`, data);
    }
    async delete_column(tid, cid) {
        return this.delete(`/dtable/${tid}/column/${cid}`);
    }
    // view stuff
    async list_view(tid) {
        return this.get(`/dtable/${tid}/view`);
    }
    async new_view(tid, data) {
        return this.post(`/dtable/${tid}/view`, data);
    }
    async modify_view(tid, id, data) {
        return this.post(`/dtable/${tid}/view/${id}`, data);
    }
    async get_view(tid, id) {
        return this.get(`/dtable/${tid}/view/${id}`);
    }
    async del_view(tid, id) {
        return this.delete(`/dtable/${tid}/view/${id}`);
    }
    // hook stuff
    async list_hook(tid) {
        return this.get(`/dtable/${tid}/hook`);
    }
    async new_hook(tid, data) {
        return this.post(`/dtable/${tid}/hook`, data);
    }
    async modify_hook(tid, id, data) {
        return this.post(`/dtable/${tid}/hook/${id}`, data);
    }
    async get_hook(tid, id) {
        return this.get(`/dtable/${tid}/hook/${id}`);
    }
    async del_hook(tid, id) {
        return this.delete(`/dtable/${tid}/hook/${id}`);
    }
    // dtable ops
    async new_row(tid, data) {
        return this.post(`/dtable_ops/${tid}/row`, data);
    }
    async get_row(tid, rid) {
        return this.get(`/dtable_ops/${tid}/row/${rid}`);
    }
    async update_row(tid, rid, data) {
        return this.post(`/dtable_ops/${tid}/row/${rid}`, data);
    }
    async delete_row(tid, rid) {
        return this.delete(`/dtable_ops/${tid}/row/${rid}`);
    }
    async simple_query(tid, data) {
        if (!data) {
            data = {};
        }
        return this.post(`/dtable_ops/${tid}/simple_query`, data);
    }
    async fts_query(tid, str) {
        return this.post(`/dtable_ops/${tid}/fts_query`, {
            "qstr": str
        });
    }
    async ref_load(tid, data) {
        return this.post(`/dtable_ops/${tid}/ref_load`, data);
    }
    async ref_resolve(tid, data) {
        return this.post(`/dtable_ops/${tid}/ref_resolve`, data);
    }
    async rev_ref_load(tid, data) {
        return this.post(`/dtable_ops/${tid}/rev_ref_load`, data);
    }
    async list_activity(tid, rowid) {
        return this.get(`/dtable_ops/${tid}/activity/${rowid}`);
    }
    async comment_row(tid, rowid, msg) {
        return this.post(`/dtable_ops/${tid}/activity/${rowid}`, {
            "message": msg,
        });
    }
}
exports.DtableAPI = DtableAPI;
class DynAPI extends base_1.ApiBase {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin"]
        });
    }
    async list_group(source) {
        return this.get(`/dgroup/${source}`);
    }
    async get_group(source, group) {
        return this.get(`/dgroup/${source}/${group}`);
    }
    async new_group(source, data) {
        return this.post(`/dgroup/${source}`, data);
    }
    async edit_group(source, gid, data) {
        return this.patch(`/dgroup/${source}/${gid}`, data);
    }
    async delete_group(source, gid) {
        return this.delete(`/dgroup/${source}/${gid}`);
    }
}
exports.DynAPI = DynAPI;
class BasicAPI extends base_1.ApiBase {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin"]
        });
    }
    async list_cabinet_sources() {
        return this.get(`/cabinet_sources`);
    }
    async list_dgroup_sources() {
        return this.get(`/dgroup`);
    }
    async message_user(data) {
        return this.post("/self/message_user", data);
    }
    async get_user_info(userid) {
        return this.get(`/self/get_user_info/${userid}`);
    }
    async get_self_info() {
        return this.get("/self/get_self_info");
    }
    async update_self_info(data) {
        return this.post("/self/get_self_info", data);
    }
    async self_change_email(data) {
        return this.post("/self/change_email", data);
    }
    async self_change_auth(data) {
        return this.post("/self/change_auth", data);
    }
    async list_messages(data) {
        return this.post("/self/list_messages", data);
    }
    async modify_messages(data) {
        return this.post("/self/modify_messages", data);
    }
    async dtable_change(data) {
        return this.post("/self/dtable_change", data);
    }
    get_session_token() {
        return this._session_token;
    }
}
exports.BasicAPI = BasicAPI;
class RepoAPI {
    constructor(bapi) {
        this.basic_api = bapi;
    }
    async repo_sources() {
        return this.basic_api.get(`/repo`);
    }
    async repo_list(source) {
        return this.basic_api.get(`/repo/${source}`);
    }
    async repo_get(source, group, slug) {
        return this.basic_api.get(`/repo/${source}/${group}/${slug}`);
    }
    async repo_get_file(source, slug, file) {
        return this.basic_api.get(`/repo/${source}/${slug}/${file}`);
    }
}
exports.RepoAPI = RepoAPI;
class ResourceAPI extends base_1.ApiBase {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["resource"]
        });
    }
    async agent_resources_list(data) {
        return this.post("/agent_resources", data);
    }
    async resource_list() {
        return this.get("/resource");
    }
    async resource_create(data) {
        return this.post("/resource", data);
    }
    async resource_get(slug) {
        return this.get(`/resource/${slug}`);
    }
    async resource_update(slug, data) {
        return this.post(`/resource/${slug}`, data);
    }
    async resource_remove(slug) {
        return this.delete(`/resource/${slug}`);
    }
}
exports.ResourceAPI = ResourceAPI;
class EngineAPI extends base_1.ApiBase {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["engine"]
        });
    }
    async launcher_json(plug, agent, data) {
        return this.post(`/engine/${plug}/${agent}/launcher/json`, data);
    }
    async referer_ticket(plug, agent) {
        return this.get(`/engine/${plug}/${agent}/referer_ticket`);
    }
}
exports.EngineAPI = EngineAPI;
