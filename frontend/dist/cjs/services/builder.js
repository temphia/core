"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildApp = void 0;
const _1 = require(".");
const authstore_1 = require("../lib/authstore");
exports.buildApp = (modal_open, modal_close, toaster) => {
    const data = authstore_1.get_current_authed();
    const __app = new _1.AppService({
        site_token: data.site_token,
        tenant_id: data.tenant_id,
        url_base: data.base_url,
        user_claim: data.user_claim,
        api_url: data.api_url,
        simple_modal_close: modal_close,
        simple_modal_open: modal_open,
        toaster,
    });
    return __app;
};
