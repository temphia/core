"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Notification = void 0;
const store_1 = require("svelte/store");
class Notification {
    constructor(opts) {
        this.handler = (message) => {
            switch (message.payload["type"]) {
                case "new":
                    this.state.update((old) => {
                        return {
                            ...old,
                            messages: [...old.messages, message.payload["data"]],
                        };
                    });
                    break;
                default:
                    console.log("@message =>", message);
                    break;
            }
        };
        this.init = async () => {
            this.state.update((old) => ({ ...old, loading: true }));
            const resp = await this.basicAPI.list_messages({});
            if (resp.status !== 200) {
                console.warn("Error happend", resp);
                return;
            }
            this.state.update((old) => {
                return {
                    ...old,
                    cursor: 0,
                    loading: false,
                    messages: resp.data,
                };
            });
        };
        this.set_read_notifications = async (id) => {
            await this.basicAPI.modify_messages({
                ops: "read",
                ids: [id],
            });
            return this.init();
        };
        this.delete_notification = async (id) => {
            await this.basicAPI.modify_messages({
                ops: "delete",
                ids: [id],
            });
            return this.init();
        };
        this.toggle_notification = () => {
            this.is_open.update((old) => !old);
        };
        const room = opts.sockdMuxer.get_notification_room();
        room.onServer(this.handler);
        this.basicAPI = opts.basicAPI;
        this.state = store_1.writable({
            cursor: 0,
            loading: false,
            messages: [],
        });
        this.isPendingRead = store_1.derived([this.state], ([state]) => {
            let pending = false;
            state.messages.forEach((msg) => {
                if (!msg["read"]) {
                    pending = true;
                    return;
                }
            });
            return pending;
        });
        this.is_open = store_1.writable(false);
    }
}
exports.Notification = Notification;
