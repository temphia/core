import type { Writable } from "svelte/store";
import type { DirtyData } from "./data_types";
export declare type Callback = () => void;
export declare class RowEditor {
    _dirtyStore: Writable<DirtyData>;
    _callbacks: Map<string, Callback>;
    constructor(store: Writable<DirtyData>);
    RegisterBeforeSave(field: string, callback: Callback): void;
    OnChange(_filed: string, _value: any): void;
    startModifyRow: (row: number) => void;
    startNewRow: () => void;
    setValue: (_filed: string, value: any) => void;
    clearDirtyRow: () => void;
    setRefCopy(column: string, value: any): void;
    beforeSave(): void;
}
