import type { DtableAPI } from "../../lib/api/impl"
import { CommonStore } from "./store"
import { DataTableService } from "./data_table"
import type { EngineService } from "../engine"
import type { FilterItem } from "./data_types"

export interface GroupOptions {
    tables: object[]
    group: string
    cabinet_ticket: string
    sockd_ticket: string
}

export class DataGroupService {
    source: string
    group: string
    groupAPI: DtableAPI
    store: CommonStore
    engine_service: EngineService

    tmanagers: Map<string, DataTableService>
    options: GroupOptions
    constructor(source: string, group: string, gapi: DtableAPI, es: EngineService) {
        this.source = source
        this.group = group
        this.groupAPI = gapi
        this.tmanagers = new Map()
        this.store = new CommonStore()
        this.engine_service = es
    }

    init = async (): Promise<void> => {
        const resp = await this.groupAPI.load_group()
        if (resp.status !== 200) {
            console.warn("err loading group", resp)
            return null
        }
        this.options = resp.data
    }


    get_table_service = async (table: string, opts: FilterItem) => {
        if (!this.options) {
            await this.init()
        }

        let svc = this.tmanagers.get(table)
        if (!svc) {
            svc = new DataTableService({
                api: this.groupAPI,
                cabinet_ticket: this.options.cabinet_ticket,
                current_table: table,
                group: this.group,
                sockd_ticket: this.options.sockd_ticket,
                tables: this.options.tables,
                store: this.store,
                engine_service: this.engine_service,
            })
            if (!opts) {
                await svc.init()
            }

            this.tmanagers.set(table, svc)
        }

        if (opts) {
            await svc.applyView("nav_with_options", {
                count: 0,
                filter_conds: [opts],
                main_column: "",
                search_term: "",
                selects: [],
            })
        }
        
        return svc
    }


    default_table = (): string => {
        return this.options.tables[0]["slug"]
    }


    close = async () => {
        this.tmanagers.forEach((manager) => manager.close())
        this.tmanagers.clear()
    }

}
