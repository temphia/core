import axios, { AxiosInstance } from "axios";

export interface LoginResponse {
    message?: string
    user_token?: string
    status_ok: boolean
    redirrect_to?: string
}

export interface AuthenticatorOptions {
    api_base_url: string
    site_token: string
    tenant_id: string
}

export class Authenticator {
    _api_base_url: string
    _site_token: string
    _tenant_id: string
    _http_client: AxiosInstance

    constructor(opts: AuthenticatorOptions) {
        this._site_token = opts.site_token
        this._tenant_id = opts.tenant_id
        this._api_base_url = opts.api_base_url

        this._http_client = axios.create({
            baseURL: this._api_base_url,
        })
    }

    async LoginWithPassword(user_ident: string, password: string): Promise<LoginResponse> {
        let resp = await this._http_client.post("/auth/login", {
            tenant_id: this._tenant_id,
            user_idendity: user_ident,
            password: password,
            site_token: this._site_token
        })
        if (resp.status == 200) {
            return {
                status_ok: true,
                user_token: resp.data["token"] || "",
            }
        }
        return {
            message: resp.data.message,
            status_ok: false,
        }
    }
}