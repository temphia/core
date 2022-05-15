<script lang="ts">
  import PlugEditor from "./_plug_editor.svelte";
  import type { AppService } from "../../../../../services";
  import { getContext } from "svelte";
  export let id;

  const app: AppService = getContext("__app__");

  let plug;

  app.apm.get_plug_api().then(async (papi) => {
    const resp = await papi.get_plug(id);
    plug = resp.data;
  });
</script>

{#if plug}
  <PlugEditor {plug} />
{/if}
