import { AppService } from "."
import { get_current_authed, } from "../lib/authstore"
import type { Toaster } from "./app"

export const buildApp = (modal_open: any, modal_close: any, toaster: Toaster) => {
    const data = get_current_authed()
    const __app = new AppService({
        site_token: data.site_token,
        tenant_id: data.tenant_id,
        url_base: data.base_url,
        user_claim: data.user_claim,
        api_url: data.api_url,
        simple_modal_close: modal_close,
        simple_modal_open: modal_open,
        toaster,
    })
    return __app
}