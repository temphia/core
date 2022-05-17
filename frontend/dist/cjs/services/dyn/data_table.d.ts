import { Writable } from "svelte/store";
import type { DtableAPI } from "../../lib/api/impl";
import { FolderAPI } from "../../lib/api/folder";
import { DirtyData, NavData, ViewData } from "./data_types";
import { RowEditor } from "./roweditor";
import type { CommonStore } from "./store";
import type { EngineService, PlugExec } from "../engine";
export interface TableOptions {
    tables: object[];
    group: string;
    cabinet_ticket: string;
    sockd_ticket: string;
    api: DtableAPI;
    current_table: string;
    store: CommonStore;
    engine_service: EngineService;
}
export declare class DataTableService {
    api: DtableAPI;
    dtable: string;
    store: CommonStore;
    tableData: object;
    dirtyStore: Writable<DirtyData>;
    navStore: Writable<NavData>;
    lastLoading: number;
    groupOpts: TableOptions;
    folderAPI: FolderAPI;
    loading: boolean;
    row_editor: RowEditor;
    hook_executor: HookExecutor;
    constructor(opts: TableOptions);
    init: () => Promise<void>;
    reset: () => Promise<void>;
    reload: () => Promise<void>;
    saveData: (data: object, append: boolean) => void;
    fetchRowLatest: (rowid: number) => Promise<void>;
    deleteRow: (rowid: number) => Promise<void>;
    applyView: (name: string, view: ViewData) => Promise<void>;
    reachedTop: () => Promise<void>;
    reachedButtom: () => Promise<void>;
    private skipLoading;
    saveDirtyRow: () => Promise<any>;
    private _saveRow;
    do_query: (query: {
        count: number;
        filter_conds: object[];
        page: number;
        selects: string[];
        search_term: string;
        load_extra_meta: boolean;
    }) => Promise<any>;
    close: () => void;
    ref_load: (data: any) => Promise<import("axios").AxiosResponse<any>>;
    ref_resolve_pri: (ref_type: string, target_table: string, target_column: string, ids: number[]) => Promise<import("axios").AxiosResponse<any>>;
    list_activity: (rowId: number) => Promise<import("axios").AxiosResponse<any>>;
    rev_ref_load: (target_table: string, target_column: string, rowid: number) => Promise<import("axios").AxiosResponse<any>>;
    comment_row: (rowId: number, message: string) => Promise<import("axios").AxiosResponse<any>>;
    set_ref_callback: (fn: () => HTMLElement) => void;
}
export declare class HookExecutor {
    _engine: EngineService;
    _table_service: DataTableService;
    _active_execs: Map<number, PlugExec>;
    get_target_ref: () => HTMLElement;
    constructor(e: EngineService, dts: DataTableService);
    execute_hook: (hook: object) => Promise<void>;
    close: () => void;
    on_message: (hookid: number, data: any) => void;
    command_data_hello: (message: string) => void;
}
