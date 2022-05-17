export declare class OperatorAPI {
    token: string;
    baseURL: string;
    tenantURL: string;
    constructor(token: string, baseURL: string);
    create_tenant: (data: object) => Promise<any>;
    list_tenant: () => Promise<any>;
    update_tenant: (data: object) => Promise<any>;
    delete_tenant: (id: string) => Promise<any>;
    header: () => {
        'Content-Type': string;
        Authorization: string;
    };
}
export declare const OpLogin: (baseURL: string, user: string, password: string) => Promise<Response>;
