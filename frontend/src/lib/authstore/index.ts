interface AuthedData {
    base_url: string
    user_claim: string
    site_token: string
    api_url: string
    tenant_id: string
}

const ALL_AUTHED = '__temphia_authed_all'
export const update_authed_all = (name: string, slug: string): void => {
    let all = []
    try {
        let all = JSON.parse(localStorage.getItem(ALL_AUTHED))
    } catch (error) {
        all = []
    }
    all.push({
        name,
        slug
    })

    localStorage.setItem(ALL_AUTHED, JSON.stringify(all))
}

export const get_update_authed_all = () => {
    try {
        return JSON.parse(localStorage.getItem(ALL_AUTHED))
    } catch (error) {
    }

    return []
}


const key = (tenant_id) => `temphia_authed_${tenant_id}`
const currentTenant = `temphia_current_tenant`

export const set_authed_data = (data: AuthedData): void => {
    localStorage.setItem(key(data.tenant_id), JSON.stringify(data))
    update_authed_all(data.tenant_id, data.tenant_id)
    localStorage.setItem(currentTenant, data.tenant_id)
}

export const get_current_authed = (): AuthedData => {
    const tenant_id = localStorage.getItem(currentTenant)
    return JSON.parse(localStorage.getItem(key(tenant_id)))
}

export const get_authed_data = (tenant_id: string): AuthedData => {
    return JSON.parse(localStorage.getItem(key(tenant_id)))
}

export const clear_authed_data = (tenant_id: string) => {

    localStorage.removeItem(key(tenant_id))
}
