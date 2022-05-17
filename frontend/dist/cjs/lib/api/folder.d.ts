export declare class FolderAPI {
    ticket: string;
    base_url: string;
    constructor(base_url: string, ticket: string);
    list(): Promise<any>;
    upload_file(file: string, data?: any): Promise<any>;
    get_file_link(file: string): string;
    get_file_preview_link(file: string): string;
}
