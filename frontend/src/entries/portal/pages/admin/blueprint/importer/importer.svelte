<script lang="ts">
  import { getContext } from "svelte";
  import type { AppService } from "../../../../../../services";
  import Importer from "./_importer.svelte";

  export let data;
  export let group;
  export let source;

  const app: AppService = getContext("__app__");

  const importBprint = async (_data: any) => {
    const bapi = await app.apm.get_bprint_api();

    return bapi.bprint_import(_data);
  };
</script>

<Importer {data} {source} import_func={importBprint} />
