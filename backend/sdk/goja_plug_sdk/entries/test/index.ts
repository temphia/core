import { core, registerHandler } from "../../lib/index";

function main(event) {
    core.log("This is test" + JSON.stringify(event) )
}

registerHandler(main)