import { Link, navigate } from "svelte-routing";

export class Navigator {
    current: string
    base_url: string
    nav_options?: any

    constructor(base_url: string) {
        this.base_url = base_url
    }

    goto = (path: string, opts?: any) => {
        this.nav_options = opts
        this.current = path
        navigate(path, { replace: true })
    }

    goto_dtable_source = (source: string) => {
        this.goto(`/z/portal/dtable/${source}`)
    }

    goto_dtable_group = (source: string, group: string) => {
        this.goto(`/z/portal/dtable/${source}/${group}`)
    }

    goto_dtable = (source: string, group: string, table: string, opts?: any) => {
        this.goto(`/z/portal/dtable/${source}/${group}/${table}`, opts)
    }

    goto_repo_store = () => {
        this.goto(`/z/portal/store`)
    }

    goto_repo_item = (repo: string, group: string, item: string) => {
        this.goto(`/z/portal/store/${repo}/${group}/${item}`)
    }
    goto_cabinet_source = (src: string) => {
        this.goto(`/z/portal/cabinet/${src}`)
    }

    goto_cabinet_folder = (src: string, folder: string) => {
        this.goto(`/z/portal/cabinet/${src}/${folder}`)
    }

    goto_cabinet_file = (src: string, folder: string, file: string) => {
        this.goto(`/z/portal/cabinet/${src}/${folder}/${file}`)
    }

    // admin pages
    goto_admin_bprints_page = () => {
        this.goto(`/z/portal/admin/bprints`)
    }

    goto_admin_bprint_page = (bid: string) => {
        this.goto(`/z/portal/admin/bprints/${bid}`)
    }

    goto_admin_plugs_page = () => {
        this.goto("/z/portal/admin/plugs")
    }

    goto_admin_plug_page = (id: string) => {
        this.goto(`/z/portal/admin/plugs/${id}`)
    }

    goto_admin_agents_page = (pid: string) => {
        this.goto(`/z/portal/admin/plugs/${pid}/agents`)
    }
    goto_admin_agent_page = (pid: string, aid: string) => {
        this.goto(`/z/portal/admin/plugs/${pid}/agents/${aid}`)
    }

    goto_admin_agent_resources_page = (pid: string, aid: string) => {
        this.goto(`/z/portal/admin/plugs/${pid}/agents/${aid}/resources`)
    }

    goto_admin_resources_page = () => {
        this.goto("/z/portal/admin/resources")
    }

    goto_admin_resource_new = (plug?: string) => {
        this.goto(`/z/portal/admin/resources/new`)
    }

    goto_admin_resource_edit = (id: string) => {
        this.goto(`/z/portal/admin/resources/edit/${id}`)
    }

    // users

    goto_admin_users_page = () => {
        this.goto("/z/portal/admin/users")
    }

    goto_admin_new_user_page = () => {
        this.goto("/z/portal/admin/new_user")
    }

    goto_admin_user_page = (id: string) => {
        this.goto(`/z/portal/admin/users/${id}`)
    }
    goto_admin_usergroups_page = () => {
        this.goto("/z/portal/admin/user_groups")
    }
    goto_admin_usergroup_page = (id: string) => {
        this.goto(`/z/portal/admin/user_groups/${id}`)
    }

    goto_admin_new_usergroup_page = () => {
        this.goto(`/z/portal/admin/new_user_groups`)
    }

    goto_admin_user_by_group = (id: string) => {
        this.goto(`/z/portal/admin/users_by_group/${id}`)
    }

    // user group auth

    goto_admin_user_auth_new = (gid: string) => {
        this.goto(`/z/portal/admin/user_group_auth/${gid}/new`)
    }

    goto_admin_user_auth_edit = (gid: string, id: string) => {
        this.goto(`/z/portal/admin/user_group_auth/${gid}/edit/${id}`)
    }

    // user group qapp

    goto_admin_user_hook_new = (gid: string) => {
        this.goto(`/z/portal/admin/user_group_hook/${gid}/new`)
    }

    goto_admin_user_hook_edit = (gid: string, id: string) => {
        this.goto(`/z/portal/admin/user_group_hook/${gid}/edit/${id}`)
    }

    goto_admin_user_plug_new = (gid: string) => {
        this.goto(`/z/portal/admin/user_group_plug/${gid}/new`)
    }

    goto_admin_user_plug_edit = (gid: string, id: string) => {
        this.goto(`/z/portal/admin/user_group_plug/${gid}/edit/${id}`)
    }

    goto_admin_user_data_new = (gid: string) => {
        this.goto(`/z/portal/admin/user_group_data/${gid}/new`)
    }

    goto_admin_user_data_edit = (gid: string, id: string) => {
        this.goto(`/z/portal/admin/user_group_data/${gid}/edit/${id}`)
    }


    // dsource

    goto_admin_dtable_builder = (bid?: string) => {
        if (bid) {
            this.goto(`/z/portal/admin/builder/builder/${bid}`)
            return
        }
        this.goto(`/z/portal/admin/builder/builder`)
    }

    goto_admin_dsource_page = () => {
        this.goto("/z/portal/admin/dtable")
    }

    goto_admin_dgroup_page = (source: string, group: string) => {
        this.goto(`/z/portal/admin/dtable/${source}/${group}`)
    }

    goto_admin_dtable_page = (source: string, group: string, table: string) => {
        this.goto(`/z/portal/admin/dtable/${source}/${group}/${table}`)
    }

    goto_dgroup_edit = (source: string, group: string) => {
        this.goto(`/z/portal/admin/dtable_edit/${source}/${group}`)
    }

    goto_dtable_edit = (source: string, group: string, table: string) => {
        this.goto(`/z/portal/admin/dtable_edit/${source}/${group}/${table}`)
    }

    goto_column_edit = (source: string, group: string, table: string, column: string) => {
        this.goto(`/z/portal/admin/dtable_edit/${source}/${group}/${table}/${column}`)
    }

    goto_views = (source: string, group: string, table: string) => {
        this.goto(`/z/portal/admin/table_views/${source}/${group}/${table}`)
    }

    goto_add_view = (source: string, group: string, table: string) => {
        this.goto(`/z/portal/admin/table_views/${source}/${group}/${table}/new`)
    }

    goto_edit_view = (source: string, group: string, table: string, id: string) => {
        this.goto(`/z/portal/admin/table_views/${source}/${group}/${table}/edit/${id}`)
    }

    goto_hooks = (source: string, group: string, table: string) => {
        this.goto(`/z/portal/admin/table_hooks/${source}/${group}/${table}`)
    }

    goto_add_hook = (source: string, group: string, table: string) => {
        this.goto(`/z/portal/admin/table_hooks/${source}/${group}/${table}/new`)
    }

    goto_edit_hook = (source: string, group: string, table: string, id: string) => {
        this.goto(`/z/portal/admin/table_hooks/${source}/${group}/${table}/edit/${id}`)
    }

    extern_plug_launch = (plug: string, agent: string) => {
        window.history.pushState('', '', '?' + "referer_token=fixmewithactualtoken");
        window.open(`${this.base_url}/engine/${plug}/${agent}/launcher/html`, '_blank');
    }

    iframe_plug_launch = (plug: string, agent: string) => {
        this.goto(`/z/portal/launcher/${plug}/${agent}`)
    }

    launcher_apps = () => {
        this.goto("/z/portal/apps_launcher")
    }

    goto_self_profile = () => {
        this.goto(`/z/portal/self_profile`)
    }

    goto_user_profile = (user) => {
        this.goto(`/z/portal/user_profile/${user}`)
    }

    goto_org_profile = () => {
        this.goto("/z/portal/about_ns")
    }

    goto_admin_org_edit = () => {
        this.goto("/z/portal/admin/ns_edit")
    }
}