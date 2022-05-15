<script lang="ts">
  import { getContext } from "svelte";

  import type { AppService } from "../../../../../../services";

  import ObjectEditor from "../_object_editor.svelte";
  export let agent = {};
  export let mod_agent = {};
  export let modified = false;

  const app: AppService = getContext("__app__");
</script>

<ObjectEditor
  data={agent}
  bind:mod_data={mod_agent}
  bind:modified
  name="Agent Editor"
  onSave={async () => {
    const papi = await app.apm.get_plug_api();
    await papi.update_agent(agent["plug_id"], agent["id"], mod_agent);
  }}
  schema={{
    primary: "id",
    text: [
      "name",
      "type",
      "invoke_policy",
      "plug_id",
      "entry_name",
      "entry_script",
      "entry_style",
    ],
    object: ["resources", "serve_files", "ext_scripts", "extra_meta"],
  }}
/>
