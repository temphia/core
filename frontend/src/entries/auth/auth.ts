import { set_authed_data } from "../../lib";
import { Authenticator } from "../../lib/api";
import type { SiteData } from "../../lib/core/types";

let auth: Authenticator
let siteData: SiteData
export let doLogin = async (user, pass) => {
    if (!auth) {
        siteData = window["__temphia_site_data__"] as SiteData
        if (!siteData) {
            console.warn("site data not loaded")
            return
        }
        auth = new Authenticator({
            api_base_url: siteData.api_url,
            site_token: siteData.site_token,
            tenant_id: siteData.tenant_id,
        })
    }

    let resp = await auth.LoginWithPassword(user, pass)
    if (!resp.status_ok) {
        return resp.message
    }

    set_authed_data({
        api_url: siteData.api_url,
        base_url: window.location.origin,
        site_token: siteData.site_token,
        tenant_id: siteData.tenant_id,
        user_claim: resp.user_token
    })

    if (!resp.redirrect_to) {
        window.location.href = `${window.location.origin}/console`;
        return ""
    }

    if (resp.redirrect_to.startsWith("http://") || resp.redirrect_to.startsWith("https://")) {
        window.location.href = resp.redirrect_to;
        return ""
    }

    window.location.href = `${window.location.origin}/${resp.redirrect_to}`;
    return ""
};

