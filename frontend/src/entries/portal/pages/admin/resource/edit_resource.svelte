<script lang="ts">
  import EditResource from "./_edit_resource.svelte";
  import Layout from "../layout.svelte";
  import type { AppService } from "../../../../../services";
  import { getContext } from "svelte";

  export let id = "";

  const app: AppService = getContext("__app__");
  let data;

  const load = async () => {
    const rapi = await app.apm.get_resource_api();
    const resp = await rapi.resource_get(id);
    data = resp.data;
  };
  
  load()

</script>

<Layout current_item="resources">
  {#if data}
    <EditResource {data} />
  {/if}
</Layout>
