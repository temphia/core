<script lang="ts">
  import { getContext } from "svelte";
  import type { AppService } from "../../../../services";

  export let source;
  export let dgroup;

  const app: AppService = getContext("__app__");

  app.get_data_service(source, dgroup).then((resp) => {
    const table = resp.default_table();
    app.navigator.goto_dtable(source, dgroup, table)
  });
  
</script>

<div>Loading Group...</div>
