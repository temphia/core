<script lang="ts">
  import { getContext, onMount } from "svelte";
  import type { AppService } from "../../../../services";
  import type { PlugExec } from "../../../../services/engine";
  export let plugid = "";
  export let agentid = "";

  let launchRef;

  const app: AppService = getContext("__app__");
  let exec: PlugExec;

  onMount(async () => {
    const engine = app.engine_service;
    exec = await engine.instance_stdplug(plugid, agentid);
    exec.run(launchRef, {})
  });

</script>

<div class="h-full w-full" bind:this={launchRef} />
