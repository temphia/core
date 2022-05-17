"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.hslColor = exports.numHash = exports.strHash = exports.generateId = void 0;
exports.generateId = () => Math.random().toString(36).slice(2);
exports.strHash = (str) => {
    let hash = 0;
    for (let i = 0; i < str.length; i++) {
        const char = str.charCodeAt(i);
        hash = (hash << 5) - hash + char;
        hash &= hash; // Convert to 32bit integer
    }
    return new Uint32Array([hash])[0].toString(36);
};
const pp = ".*(D#D01e-u0_ue819g_!UJ123456789023";
exports.numHash = (str) => {
    let hash = 77;
    for (var i = 0; i < str.length; i++) {
        hash = str.charCodeAt(i) + ((hash << 6) - hash);
        hash = pp.charCodeAt(i) ^ hash;
    }
    return hash;
};
exports.hslColor = (str) => {
    return `background: hsl(${exports.numHash(str) % 360}, 100%, 80%)`;
};
// const assetsUrls = {
//     "leaflet_css": "https://unpkg.com/leaflet@1.6.0/dist/leaflet.css",
//     "leaflet_js": "https://unpkg.com/leaflet@1.6.0/dist/leaflet.js",
//     "flatpicker_css": "",
// }
