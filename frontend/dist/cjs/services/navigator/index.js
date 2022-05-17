"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Navigator = void 0;
const svelte_routing_1 = require("svelte-routing");
class Navigator {
    constructor(base_url) {
        this.goto = (path, opts) => {
            this.nav_options = opts;
            this.current = path;
            svelte_routing_1.navigate(path, { replace: true });
        };
        this.goto_dtable_source = (source) => {
            this.goto(`/console/dtable/${source}`);
        };
        this.goto_dtable_group = (source, group) => {
            this.goto(`/console/dtable/${source}/${group}`);
        };
        this.goto_dtable = (source, group, table, opts) => {
            this.goto(`/console/dtable/${source}/${group}/${table}`, opts);
        };
        this.goto_repo_store = () => {
            this.goto(`/console/store`);
        };
        this.goto_repo_item = (repo, group, item) => {
            this.goto(`/console/store/${repo}/${group}/${item}`);
        };
        this.goto_cabinet_source = (src) => {
            this.goto(`/console/cabinet/${src}`);
        };
        this.goto_cabinet_folder = (src, folder) => {
            this.goto(`/console/cabinet/${src}/${folder}`);
        };
        this.goto_cabinet_file = (src, folder, file) => {
            this.goto(`/console/cabinet/${src}/${folder}/${file}`);
        };
        // admin pages
        this.goto_admin_bprints_page = () => {
            this.goto(`/console/admin/bprints`);
        };
        this.goto_admin_bprint_page = (bid) => {
            this.goto(`/console/admin/bprints/${bid}`);
        };
        this.goto_admin_plugs_page = () => {
            this.goto("/console/admin/plugs");
        };
        this.goto_admin_plug_page = (id) => {
            this.goto(`/console/admin/plugs/${id}`);
        };
        this.goto_admin_agents_page = (pid) => {
            this.goto(`/console/admin/plugs/${pid}/agents`);
        };
        this.goto_admin_agent_page = (pid, aid) => {
            this.goto(`/console/admin/plugs/${pid}/agents/${aid}`);
        };
        this.goto_admin_agent_resources_page = (pid, aid) => {
            this.goto(`/console/admin/plugs/${pid}/agents/${aid}/resources`);
        };
        this.goto_admin_resources_page = () => {
            this.goto("/console/admin/resources");
        };
        this.goto_admin_resource_new = (plug) => {
            this.goto(`/console/admin/resources/new`);
        };
        this.goto_admin_resource_edit = (id) => {
            this.goto(`/console/admin/resources/edit/${id}`);
        };
        // users
        this.goto_admin_users_page = () => {
            this.goto("/console/admin/users");
        };
        this.goto_admin_new_user_page = () => {
            this.goto("/console/admin/new_user");
        };
        this.goto_admin_user_page = (id) => {
            this.goto(`/console/admin/users/${id}`);
        };
        this.goto_admin_usergroups_page = () => {
            this.goto("/console/admin/user_groups");
        };
        this.goto_admin_usergroup_page = (id) => {
            this.goto(`/console/admin/user_groups/${id}`);
        };
        this.goto_admin_new_usergroup_page = () => {
            this.goto(`/console/admin/new_user_groups`);
        };
        this.goto_admin_user_by_group = (id) => {
            this.goto(`/console/admin/users_by_group/${id}`);
        };
        // user group auth
        this.goto_admin_user_auth_new = (gid) => {
            this.goto(`/console/admin/user_group_auth/${gid}/new`);
        };
        this.goto_admin_user_auth_edit = (gid, id) => {
            this.goto(`/console/admin/user_group_auth/${gid}/edit/${id}`);
        };
        // user group qapp
        this.goto_admin_user_hook_new = (gid) => {
            this.goto(`/console/admin/user_group_hook/${gid}/new`);
        };
        this.goto_admin_user_hook_edit = (gid, id) => {
            this.goto(`/console/admin/user_group_hook/${gid}/edit/${id}`);
        };
        this.goto_admin_user_plug_new = (gid) => {
            this.goto(`/console/admin/user_group_plug/${gid}/new`);
        };
        this.goto_admin_user_plug_edit = (gid, id) => {
            this.goto(`/console/admin/user_group_plug/${gid}/edit/${id}`);
        };
        this.goto_admin_user_data_new = (gid) => {
            this.goto(`/console/admin/user_group_data/${gid}/new`);
        };
        this.goto_admin_user_data_edit = (gid, id) => {
            this.goto(`/console/admin/user_group_data/${gid}/edit/${id}`);
        };
        // dsource
        this.goto_admin_dtable_builder = (bid) => {
            if (bid) {
                this.goto(`/console/admin/builder/builder/${bid}`);
                return;
            }
            this.goto(`/console/admin/builder/builder`);
        };
        this.goto_admin_dsource_page = () => {
            this.goto("/console/admin/dtable");
        };
        this.goto_admin_dgroup_page = (source, group) => {
            this.goto(`/console/admin/dtable/${source}/${group}`);
        };
        this.goto_admin_dtable_page = (source, group, table) => {
            this.goto(`/console/admin/dtable/${source}/${group}/${table}`);
        };
        this.goto_dgroup_edit = (source, group) => {
            this.goto(`/console/admin/dtable_edit/${source}/${group}`);
        };
        this.goto_dtable_edit = (source, group, table) => {
            this.goto(`/console/admin/dtable_edit/${source}/${group}/${table}`);
        };
        this.goto_column_edit = (source, group, table, column) => {
            this.goto(`/console/admin/dtable_edit/${source}/${group}/${table}/${column}`);
        };
        this.goto_views = (source, group, table) => {
            this.goto(`/console/admin/table_views/${source}/${group}/${table}`);
        };
        this.goto_add_view = (source, group, table) => {
            this.goto(`/console/admin/table_views/${source}/${group}/${table}/new`);
        };
        this.goto_edit_view = (source, group, table, id) => {
            this.goto(`/console/admin/table_views/${source}/${group}/${table}/edit/${id}`);
        };
        this.goto_hooks = (source, group, table) => {
            this.goto(`/console/admin/table_hooks/${source}/${group}/${table}`);
        };
        this.goto_add_hook = (source, group, table) => {
            this.goto(`/console/admin/table_hooks/${source}/${group}/${table}/new`);
        };
        this.goto_edit_hook = (source, group, table, id) => {
            this.goto(`/console/admin/table_hooks/${source}/${group}/${table}/edit/${id}`);
        };
        this.extern_plug_launch = (plug, agent) => {
            window.history.pushState('', '', '?' + "referer_token=fixmewithactualtoken");
            window.open(`${this.base_url}/engine/${plug}/${agent}/launcher/html`, '_blank');
        };
        this.iframe_plug_launch = (plug, agent) => {
            this.goto(`/console/launcher/${plug}/${agent}`);
        };
        this.launcher_apps = () => {
            this.goto("/console/apps_launcher");
        };
        this.goto_self_profile = () => {
            this.goto(`/console/self_profile`);
        };
        this.goto_user_profile = (user) => {
            this.goto(`/console/user_profile/${user}`);
        };
        this.goto_org_profile = () => {
            this.goto("/console/about_ns");
        };
        this.goto_admin_org_edit = () => {
            this.goto("/console/admin/ns_edit");
        };
        this.base_url = base_url;
    }
}
exports.Navigator = Navigator;
