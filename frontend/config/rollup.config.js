import { config } from "./rollup.config.base";

const production = !!process.env.PRODUCTION;
const entryFile= process.env.ENTRY_FILE;
export default config( entryFile, production)