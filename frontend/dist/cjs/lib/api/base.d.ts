import { AxiosResponse, AxiosInstance, AxiosRequestConfig } from "axios";
interface BaseOptions {
    url: string;
    user_token: string;
    path: string[];
    service_opts?: any;
}
export declare class ApiBase {
    _user_token: string;
    _session_token: string;
    _http: AxiosInstance;
    _api_base_url: string;
    _raw_http: AxiosInstance;
    _service_options: object;
    _service_path: string[];
    _service_resp_payload: any;
    constructor(opts: BaseOptions);
    init(): Promise<void>;
    refresh_token(): Promise<AxiosResponse<any>>;
    intercept_request(config: AxiosRequestConfig): AxiosRequestConfig;
    intercept_request_err(error: any): Promise<never>;
    get<T = any, R = AxiosResponse<T>>(url: string, config?: AxiosRequestConfig): Promise<R>;
    post<T = any, R = AxiosResponse<T>>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R>;
    put<T = any, R = AxiosResponse<T>>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R>;
    patch<T = any, R = AxiosResponse<T>>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R>;
    delete<T = any, R = AxiosResponse<T>>(url: string, config?: AxiosRequestConfig): Promise<R>;
}
export {};
