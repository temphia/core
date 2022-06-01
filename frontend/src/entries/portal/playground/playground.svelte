<script>
  import { getContext } from "svelte";
  import PlugPick from "../pages/admin/core/plug_pick.svelte";
  import Board from "./board.svelte";
  import Builder from "./builder.svelte";

  const { open, close } = getContext("simple-modal");

  const app = getContext("__app__");
</script>

<button
  on:click={() => {
    open(PlugPick, { app });
  }}>Pickup</button
>

<Builder />
