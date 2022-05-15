import {
    _state_kv_get,
    _state_kv_tx_finish,
    _state_kv_tx_start,
    StateValue
} from "./intrinsic"


export class StateKv {
    txid: number
    done: boolean

    constructor() {
        this.txid = 0
        this.done = false
    }
    
    txStart(): StateKv {
        if (this.done) {
            throw new Error('use of finished txn handle');
        }

        if (this.txid !== 0) {
            throw new Error('nested txn not supported');
        }
        const txid = _state_kv_tx_start();
        if (typeof txid !== "number") {
            throw new Error('could not start txn');
        }
        const skv = new StateKv()
        skv.txid = txid
        return skv
    }

    txFinish(): boolean {
        if (this.txid === 0) {
            // its already zero
        }

        _state_kv_tx_finish(this.txid)
        return false
    }

    Get(key: string): StateValue {
        const val = _state_kv_get(this.txid, key)
        if (typeof val === "string") {
            throw new Error(val);
        }
        return val
    }


    Set() { }
    Del() { }
    List() { }
    ListByAudience() { }
}