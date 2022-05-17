interface AuthedData {
    base_url: string;
    user_claim: string;
    site_token: string;
    api_url: string;
    tenant_id: string;
}
export declare const update_authed_all: (name: string, slug: string) => void;
export declare const get_update_authed_all: () => any;
export declare const set_authed_data: (data: AuthedData) => void;
export declare const get_current_authed: () => AuthedData;
export declare const get_authed_data: (tenant_id: string) => AuthedData;
export declare const clear_authed_data: (tenant_id: string) => void;
export {};
