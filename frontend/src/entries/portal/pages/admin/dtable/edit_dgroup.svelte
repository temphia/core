<script lang="ts">
  import Layout from "../layout.svelte";
  import DGroup from "./_dgroup.svelte";

  import { getContext } from "svelte";
  import { DynAdminAPI } from "./dtable2";
  import type { AppService } from "../../../../../services";

  export let source;
  export let group;

  const app: AppService = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let loading = true;
  let data;

  dynapi.get_group(source, group).then((resp) => {
    data = resp.data;
    loading = false;
  });
</script>

<Layout current_item={"dtable"} {loading}>
  <DGroup {data} {source} />
</Layout>
