import { core } from "./core";
import { plugkv } from "./plugkv";
import { Request, Response } from "./http";
import { CabFolder } from "./cabinet";
import { SockdRoom } from "./sockd";

// this is noop method to prevent entry function for treeshaking
// deadcode eleminiation or whater by rollup/typescript
const registerHandler = (handler: any) => {
  JSON.stringify(handler);
};

export {
  core,
  plugkv,
  registerHandler,
  Request,
  Response,
  CabFolder,
  SockdRoom,
};
