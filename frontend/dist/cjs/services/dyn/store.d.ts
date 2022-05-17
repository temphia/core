import { Writable } from "svelte/store";
export interface State {
    indexed_column: {
        [_: string]: object;
    };
    column_order: string[];
    reverse_ref_column: object[];
    rows: number[];
    indexed_rows: {
        [_: number]: object;
    };
    sparse_rows: number[];
    remote_dirty: {
        [_: number]: true;
    };
    views: object[];
    hooks: object[];
}
export declare class CommonStore {
    states: Writable<{
        [_: string]: State;
    }>;
    constructor();
    set_rows_data: (table: string, data: any, append: boolean) => void;
    set_row_data: (table: string, data: object) => void;
    set_sockd_change: (data: any) => void;
    generate_column_order: (columns: {
        [_: string]: object;
    }) => string[];
    destroy: () => void;
}
