import { AxiosInstance } from "axios";
export interface LoginResponse {
    message?: string;
    user_token?: string;
    status_ok: boolean;
    redirrect_to?: string;
}
export interface AuthenticatorOptions {
    api_base_url: string;
    site_token: string;
    tenant_id: string;
}
export declare class Authenticator {
    _api_base_url: string;
    _site_token: string;
    _tenant_id: string;
    _http_client: AxiosInstance;
    constructor(opts: AuthenticatorOptions);
    LoginWithPassword(user_ident: string, password: string): Promise<LoginResponse>;
}
