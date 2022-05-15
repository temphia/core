<script lang="ts">
  import { getContext } from "svelte";
  import { onMount } from "svelte";
  import { StoreItem } from "../../../../components";
  import type { AppService } from "../../../../services";
  import Importer from "../admin/blueprint/importer/importer.svelte";

  export let item;
  export let source;
  export let group;

  const app: AppService = getContext("__app__");

  let data = null;
  let repo_api = null;

  onMount(async () => {
    repo_api = await app.apm.get_repo_api();
    let resp = await repo_api.repo_get(source, group, item);
    data = resp.data;
  });
</script>

{#if data}
  <StoreItem
    {data}
    actions={[
      {
        Name: "Import",
        Class: "bg-green-400",
        Action: async () => {
          app.big_modal_open(Importer, { data, group, source });
        },
      },
    ]}
  />
{/if}
