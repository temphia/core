<script>
  import { getContext } from "svelte";
  import PlugPick from "../pages/admin/core/plug_pick.svelte";
  

  const { open, close } = getContext("simple-modal");

  const app = getContext("__app__");
</script>

<button
  on:click={() => {
    open(PlugPick, { app });
  }}>Pickup</button
>
