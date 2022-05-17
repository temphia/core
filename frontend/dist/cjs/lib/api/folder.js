"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.FolderAPI = void 0;
class FolderAPI {
    constructor(base_url, ticket) {
        this.ticket = ticket;
        this.base_url = base_url;
    }
    async list() {
        const resp = await fetch(`${this.base_url}/ticket_cabinet/${this.ticket}`);
        return resp.json();
    }
    async upload_file(file, data) {
        const resp = await fetch(`${this.base_url}/ticket_cabinet/${this.ticket}/${file}`, {
            method: "POST",
            body: data,
        });
        return resp.json();
    }
    get_file_link(file) {
        return `${this.base_url}/ticket_cabinet/${this.ticket}/${file}`;
    }
    get_file_preview_link(file) {
        return `${this.base_url}/ticket_cabinet/${this.ticket}/preview/${file}`;
    }
}
exports.FolderAPI = FolderAPI;
