export class FolderAPI {
    ticket: string
    base_url: string
    constructor(base_url: string, ticket: string) {
        this.ticket = ticket
        this.base_url = base_url
    }

    async list() {
        const resp = await fetch(`${this.base_url}/ticket_cabinet/${this.ticket}`)
        return resp.json()
    }

    async upload_file(file: string, data?: any) {
        const resp = await fetch(`${this.base_url}/ticket_cabinet/${this.ticket}/${file}`, {
            method: "POST",
            body: data,
        })
        return resp.json()
    }

    get_file_link(file: string): string {
        return `${this.base_url}/ticket_cabinet/${this.ticket}/${file}`
    }

    get_file_preview_link(file: string): string {
        return `${this.base_url}/ticket_cabinet/${this.ticket}/preview/${file}`
    }
}