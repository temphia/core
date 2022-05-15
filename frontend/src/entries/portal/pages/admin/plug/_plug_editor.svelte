<script lang="ts">
  import { getContext } from "svelte";
  import type { AppService } from "../../../../../services";
  import ObjectEditor from "./_object_editor.svelte";
  export let plug = {};
  export let mod_plug = {};
  export let modified = false;

  const app: AppService = getContext("__app__");

</script>

<ObjectEditor
  data={plug}
  bind:mod_data={mod_plug}
  bind:modified
  name="Plug Editor"
  onSave={async () => {
    const papi = await app.apm.get_plug_api()
    await papi.update_plug(plug["id"], mod_plug)
  }}
  schema={{
    primary: "id",
    text: ["name", "executor", "owner", "bprint_id"],
    object: ["handlers", "extra_meta"],
    bool: ["live", "dev"],
  }}
/>
