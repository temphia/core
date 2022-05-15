<script lang="ts">
  import { getContext } from "svelte";
  import type { AppService } from "../../../../services";
  const app: AppService = getContext("__app__");

  app.get_dyn_sources().then((resp) => {
    if (!resp) {
      return;
    }
    app.navigator.goto_dtable_source(resp[0]);
  });
</script>
