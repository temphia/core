<script lang="ts">
  import { getContext } from "svelte";
  import { KeyPrimary } from "../../../../../../../lib/core/dyn";
  import type { DataTableService } from "../../../../../../../services";
  import type { Column } from "../../manager/dtypes";
  import RefPanel from "./ref_panel.svelte";

  const { open, close } = getContext("simple-modal");

  export let value;
  export let column: Column;
  export let onChange: (value: any) => void;
  export let manager: DataTableService;

  const loader = (cursor: number) => {
    return manager.ref_load({
      column: column.slug,
      type: column.ref_type,
      target: column.ref_target,
      object: column.ref_object,
      cursor_row_id: cursor,
    });
  };

  const openPanel = () => {
    open(RefPanel, {
      loader,
      onRowSelect: (row: object) => {
        console.log("ROW", row);
        onChange(row[KeyPrimary]);

        // fixme => ref_copy here

        close();
      },
    });
  };

  $: __value = value;
</script>

<div class="flex w-full">
  <input
    type="text"
    disabled
    value={__value}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1"
  />
  <button on:click={openPanel}>
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-6 w-6"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
      stroke-width="2"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"
      />
    </svg>
  </button>
</div>
