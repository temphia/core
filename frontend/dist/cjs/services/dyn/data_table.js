"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.HookExecutor = exports.DataTableService = void 0;
const store_1 = require("svelte/store");
const folder_1 = require("../../lib/api/folder");
const data_types_1 = require("./data_types");
const roweditor_1 = require("./roweditor");
class DataTableService {
    constructor(opts) {
        this.init = async () => {
            const data = await this.do_query({ ...data_types_1.defaultViewData(), load_extra_meta: true });
            if (!data) {
                console.warn("Could not fetch rows");
                return;
            }
            this.saveData(data, false);
        };
        this.reset = async () => {
            await this.init();
        };
        this.reload = async () => {
            const data = await this.do_query({ ...data_types_1.defaultViewData(), load_extra_meta: true });
            if (!data) {
                console.warn("Could not fetch rows");
                return;
            }
            this.saveData(data, false);
        };
        this.saveData = (data, append) => {
            this.store.set_rows_data(this.dtable, data, append);
        };
        this.fetchRowLatest = async (rowid) => {
            const resp = await this.api.get_row(this.dtable, rowid);
            if (resp.status !== 200) {
                return;
            }
            this.store.set_row_data(this.dtable, resp.data);
        };
        this.deleteRow = async (rowid) => {
            let resp = await this.api.delete_row(this.dtable, rowid);
            if (resp.status !== 200) {
                console.warn("could not delete row", resp);
                return;
            }
        };
        this.applyView = async (name, view) => {
            const data = await this.do_query({ ...data_types_1.defaultViewData(), ...view, load_extra_meta: false });
            if (!data) {
                console.warn("Could not fetch rows");
                return;
            }
            this.saveData(data, false);
        };
        this.reachedTop = async () => {
            console.log("TOP REACHED");
            if (this.skipLoading()) {
                return;
            }
            console.log("FETCH MORE");
        };
        this.reachedButtom = async () => {
            console.log("@start_fetch_more");
            const navdata = store_1.get(this.navStore);
            if (navdata.last_page) {
                if (this.skipLoading()) {
                    console.warn("already last page");
                    return;
                }
            }
            const data = await this.do_query({
                ...navdata.active_view, page: navdata.active_page + 1,
                load_extra_meta: false
            });
            if (!data) {
                console.warn("Could not fetch rows");
                return;
            }
            this.saveData(data, true);
            console.log("@end_fetch_more");
        };
        this.skipLoading = () => {
            if (this.loading) {
                return true;
            }
            const now = new Date().valueOf();
            if ((now - this.lastLoading) < (1000 * 10)) {
                return true;
            }
            this.lastLoading = now;
            return false;
        };
        this.saveDirtyRow = () => {
            return this._saveRow();
        };
        this._saveRow = async () => {
            const dirtyData = store_1.get(this.dirtyStore);
            if (dirtyData.rowid === 0) {
                const resp = await this.api.new_row(this.dtable, {
                    cells: dirtyData.data,
                });
                return resp.data;
            }
            const rowid = dirtyData.rowid;
            const data = store_1.get(this.store.states);
            const version = data[this.dtable].indexed_rows[rowid]["__version"];
            const resp = await this.api.update_row(this.dtable, rowid, {
                cells: dirtyData.data,
                version
            });
            if (resp.status !== 200) {
                console.warn("could not update row", resp);
                return resp.data;
            }
            this.store.set_row_data(this.dtable, resp.data);
            return resp.data;
        };
        this.do_query = async (query) => {
            this.navStore.update((old) => ({ ...old, loading: true }));
            const resp = await this.api.simple_query(this.dtable, query);
            if (resp.status !== 200) {
                this.navStore.update((old) => ({ ...old, loading: false, loading_error: resp.data }));
                return;
            }
            let last_page = false;
            if (query.count > resp.data["rows"].length) {
                last_page = true;
            }
            const active_filter_conds = query.filter_conds;
            this.navStore.update((old) => ({
                ...old,
                loading: false,
                active_view: {
                    count: query.count,
                    filter_conds: active_filter_conds,
                    main_column: "",
                    search_term: query.search_term,
                    selects: query.selects
                },
                last_page,
                active_page: query.page,
                loading_error: "",
            }));
            return resp.data;
        };
        this.close = () => {
            this.hook_executor.close();
        };
        this.ref_load = async (data) => {
            return this.api.ref_load(this.dtable, data);
        };
        this.ref_resolve_pri = async (ref_type, target_table, target_column, ids) => {
            return this.api.ref_resolve(this.dtable, {
                type: ref_type,
                target: target_table,
                column: target_column,
                row_ids: ids,
            });
        };
        this.list_activity = async (rowId) => {
            return this.api.list_activity(this.dtable, rowId);
        };
        this.rev_ref_load = async (target_table, target_column, rowid) => {
            return this.api.rev_ref_load(this.dtable, {
                "current_table": this.dtable,
                "target_table": target_table,
                "column": target_column,
                "current_item": rowid,
            });
        };
        this.comment_row = async (rowId, message) => {
            return this.api.comment_row(this.dtable, rowId, message);
        };
        this.set_ref_callback = (fn) => {
            this.hook_executor.get_target_ref = fn;
        };
        this.api = opts.api;
        this.dtable = opts.current_table;
        this.store = opts.store;
        this.groupOpts = opts;
        this.folderAPI = new folder_1.FolderAPI(opts.api._api_base_url, opts.cabinet_ticket);
        this.navStore = store_1.writable({
            loading: true,
            lastTry: null,
            loading_error: "",
            active_page: 0,
            last_page: false,
            active_view: data_types_1.defaultViewData(),
            views: {},
        });
        this.navStore.subscribe((val) => console.log("NAV_STORE @=> ", val));
        this.navStore.subscribe((navd) => this.loading = navd.loading);
        this.lastLoading = 0;
        this.navStore.subscribe((data) => console.log(data));
        this.dirtyStore = store_1.writable({ data: {}, rowid: 0 });
        this.dirtyStore.subscribe((data) => console.log("DIRTY DATA", data));
        this.row_editor = new roweditor_1.RowEditor(this.dirtyStore);
        this.hook_executor = new HookExecutor(opts.engine_service, this);
    }
}
exports.DataTableService = DataTableService;
class HookExecutor {
    constructor(e, dts) {
        this.execute_hook = async (hook) => {
            const hid = hook["id"];
            let pexec = this._active_execs[hid];
            if (!pexec) {
                pexec = await this._engine.instance_dataplug(hook);
                this._active_execs.set(hid, pexec);
                pexec.set_handler((xid, action, data) => {
                    this.on_message(hid, data); // fixme => send action too
                });
            }
            // run or re_run
            await pexec.run(this.get_target_ref(), {});
        };
        this.close = () => {
            this._active_execs.forEach((pexec) => {
                pexec.close();
            });
        };
        // commands
        this.on_message = (hookid, data) => {
            console.log("HOOK@", hookid, "DATA", data);
        };
        this.command_data_hello = (message) => {
            console.log(message);
        };
        this._engine = e;
        this._table_service = dts;
        this._active_execs = new Map();
    }
}
exports.HookExecutor = HookExecutor;
