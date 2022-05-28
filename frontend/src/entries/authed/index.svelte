<script>
  import Tailwind from "../common/_tailwind.svelte";
  import {
    parseParams,
    PAGE_AFTER_AUTH,
    PAGE_AUTH_MAIN,
    PAGE_EXTERNAL_AUTH,
  } from "./authed";
  import After from "./pages/after/after.svelte";
  import External from "./pages/external/external.svelte";
  import Loader from "./pages/loader/loader.svelte";
  import Main from "./pages/main/main.svelte";

  let loaded = false;
  let page_type;
  let data;

  parseParams().then((resp) => {
    page_type = resp.page_type;
    data = resp.data;
    loaded = true;
  });
</script>

{#if loaded}
  {#if page_type === PAGE_AUTH_MAIN}
    <Main />
  {:else if page_type === PAGE_AFTER_AUTH}
    <After />
  {:else if page_type === PAGE_EXTERNAL_AUTH}
    <External />
  {/if}
{:else}
  <Loader />
{/if}

<Tailwind />
