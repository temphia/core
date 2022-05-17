"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.clear_authed_data = exports.get_authed_data = exports.get_current_authed = exports.set_authed_data = exports.get_update_authed_all = exports.update_authed_all = void 0;
const ALL_AUTHED = '__temphia_authed_all';
exports.update_authed_all = (name, slug) => {
    let all = [];
    try {
        let all = JSON.parse(localStorage.getItem(ALL_AUTHED));
    }
    catch (error) {
        all = [];
    }
    all.push({
        name,
        slug
    });
    localStorage.setItem(ALL_AUTHED, JSON.stringify(all));
};
exports.get_update_authed_all = () => {
    try {
        return JSON.parse(localStorage.getItem(ALL_AUTHED));
    }
    catch (error) {
    }
    return [];
};
const key = (tenant_id) => `temphia_authed_${tenant_id}`;
const currentTenant = `temphia_current_tenant`;
exports.set_authed_data = (data) => {
    localStorage.setItem(key(data.tenant_id), JSON.stringify(data));
    exports.update_authed_all(data.tenant_id, data.tenant_id);
    localStorage.setItem(currentTenant, data.tenant_id);
};
exports.get_current_authed = () => {
    const tenant_id = localStorage.getItem(currentTenant);
    return JSON.parse(localStorage.getItem(key(tenant_id)));
};
exports.get_authed_data = (tenant_id) => {
    return JSON.parse(localStorage.getItem(key(tenant_id)));
};
exports.clear_authed_data = (tenant_id) => {
    localStorage.removeItem(key(tenant_id));
};
