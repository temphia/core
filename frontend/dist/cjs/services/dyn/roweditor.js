"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.RowEditor = void 0;
class RowEditor {
    constructor(store) {
        // row stuff
        this.startModifyRow = (row) => {
            this._callbacks.clear();
            this._dirtyStore.set({ rowid: row, data: {} });
        };
        this.startNewRow = () => {
            this._callbacks.clear();
            this._dirtyStore.set({ rowid: 0, data: {} });
        };
        this.setValue = (_filed, value) => {
            this._dirtyStore.update((old) => ({ ...old, data: { ...old.data, [_filed]: value } }));
        };
        this.clearDirtyRow = () => {
            this._dirtyStore.set({ rowid: 0, data: {} });
        };
        this._dirtyStore = store;
        this._callbacks = new Map();
    }
    RegisterBeforeSave(field, callback) {
        this._callbacks.set(field, callback);
    }
    OnChange(_filed, _value) {
        this.setValue(_filed, _value);
    }
    setRefCopy(column, value) {
        this._dirtyStore.update((old) => ({ ...old, data: { ...old.data, [column]: value } }));
    }
    beforeSave() {
        this._callbacks.forEach((val) => val());
    }
}
exports.RowEditor = RowEditor;
