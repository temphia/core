<script lang="ts">
  import { getContext } from "svelte";
  import type { AppService } from "../../../../../../services";
  import AgentEditor from "./_agent_editor.svelte";
  
  export let pid;
  export let aid;

  let agent;

  const app: AppService = getContext("__app__");

  app.apm.get_plug_api().then(async (api) => {
    const resp = await api.list_agent(pid);
    agent = resp.data;
  });
</script>

{#if agent}
  <AgentEditor {agent} />
{/if}
