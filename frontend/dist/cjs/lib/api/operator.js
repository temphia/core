"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.OpLogin = exports.OperatorAPI = void 0;
class OperatorAPI {
    constructor(token, baseURL) {
        this.create_tenant = async (data) => {
            let response = await fetch(this.tenantURL, {
                method: 'POST',
                headers: this.header(),
                body: JSON.stringify(data)
            });
            if (response.ok) {
                return response.json();
            }
            return response.text();
        };
        this.list_tenant = async () => {
            let response = await fetch(this.tenantURL, {
                method: 'GET',
                headers: this.header(),
            });
            if (response.ok) {
                return response.json();
            }
            return response.text();
        };
        this.update_tenant = async (data) => {
            let response = await fetch(this.tenantURL, {
                method: 'PATCH',
                headers: this.header(),
                body: JSON.stringify(data)
            });
            if (response.ok) {
                return response.json();
            }
            return response.text();
        };
        this.delete_tenant = async (id) => {
            let response = await fetch(this.tenantURL, {
                method: 'DELETE',
                headers: this.header(),
                body: JSON.stringify({ slug: id })
            });
            if (response.ok) {
                return response.json();
            }
            return response.text();
        };
        // private
        this.header = () => ({
            'Content-Type': 'application/json;charset=utf-8',
            Authorization: this.token,
        });
        this.token = token;
        this.baseURL = baseURL;
        this.tenantURL = `${this.baseURL}/operator/tenant`;
    }
}
exports.OperatorAPI = OperatorAPI;
exports.OpLogin = (baseURL, user, password) => {
    return fetch(`${baseURL}/operator/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify({
            user,
            password,
        })
    });
};
