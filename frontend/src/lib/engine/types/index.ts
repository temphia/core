export const MODE_IFRAME = "IFRAME"
export const MODE_RAW_DOM = "RAW_DOM"
export const MODE_SUB_ORIGIN = "SUB_ORIGIN"

export interface LoaderOptions {
    token: string
    entry: string
    exec_loader: string
    plug: string
    agent: string
    base_url: string
    parent_secret?: string
    startup_payload?: any
}
