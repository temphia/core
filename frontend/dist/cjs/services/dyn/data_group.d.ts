import type { DtableAPI } from "../../lib/api/impl";
import { CommonStore } from "./store";
import { DataTableService } from "./data_table";
import type { EngineService } from "../engine";
import type { FilterItem } from "./data_types";
export interface GroupOptions {
    tables: object[];
    group: string;
    cabinet_ticket: string;
    sockd_ticket: string;
}
export declare class DataGroupService {
    source: string;
    group: string;
    groupAPI: DtableAPI;
    store: CommonStore;
    engine_service: EngineService;
    tmanagers: Map<string, DataTableService>;
    options: GroupOptions;
    constructor(source: string, group: string, gapi: DtableAPI, es: EngineService);
    init: () => Promise<void>;
    get_table_service: (table: string, opts: FilterItem) => Promise<DataTableService>;
    default_table: () => string;
    close: () => Promise<void>;
}
