"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DataGroupService = void 0;
const store_1 = require("./store");
const data_table_1 = require("./data_table");
class DataGroupService {
    constructor(source, group, gapi, es) {
        this.init = async () => {
            const resp = await this.groupAPI.load_group();
            if (resp.status !== 200) {
                console.warn("err loading group", resp);
                return null;
            }
            this.options = resp.data;
        };
        this.get_table_service = async (table, opts) => {
            if (!this.options) {
                await this.init();
            }
            let svc = this.tmanagers.get(table);
            if (!svc) {
                svc = new data_table_1.DataTableService({
                    api: this.groupAPI,
                    cabinet_ticket: this.options.cabinet_ticket,
                    current_table: table,
                    group: this.group,
                    sockd_ticket: this.options.sockd_ticket,
                    tables: this.options.tables,
                    store: this.store,
                    engine_service: this.engine_service,
                });
                if (!opts) {
                    await svc.init();
                }
                this.tmanagers.set(table, svc);
            }
            if (opts) {
                await svc.applyView("nav_with_options", {
                    count: 0,
                    filter_conds: [opts],
                    main_column: "",
                    search_term: "",
                    selects: [],
                });
            }
            return svc;
        };
        this.default_table = () => {
            return this.options.tables[0]["slug"];
        };
        this.close = async () => {
            this.tmanagers.forEach((manager) => manager.close());
            this.tmanagers.clear();
        };
        this.source = source;
        this.group = group;
        this.groupAPI = gapi;
        this.tmanagers = new Map();
        this.store = new store_1.CommonStore();
        this.engine_service = es;
    }
}
exports.DataGroupService = DataGroupService;
