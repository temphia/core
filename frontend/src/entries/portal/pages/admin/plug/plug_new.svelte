<script lang="ts">
  import { getContext } from "svelte";

  import type { AppService } from "../../../../../services";
  import Kveditor from "../../../../common/kveditor.svelte";

  export let id = "";
  export let name = "";
  export let executor = "";
  export let live = false;
  export let dev = false;
  export let owner = "";
  export let bprint_id = "";
  export let handlers = {};
  export let extra_meta = {};

  const app: AppService = getContext("__app__");

  let getHandlerData;
  let getMetaData;

  const save = async () => {
      
  };
</script>

<div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
  <div class="p-5 bg-white w-full ">
    <div class="text-2xl text-indigo-900">Resource</div>

    <div class="flex-col flex py-3">
      <label class="pb-2 text-gray-700 font-semibold">Id</label>
      <input
        type="text"
        bind:value={id}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="res1"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Name</label>
      <input
        type="text"
        bind:value={name}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="signal"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Executor</label>
      <input
        type="text"
        bind:value={executor}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="signal"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Live</label>
      <input
        type="checkbox"
        bind:checked={live}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Dev</label>
      <input
        type="checkbox"
        bind:checked={dev}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Owner</label>
      <input
        type="text"
        bind:value={owner}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Owner</label>
      <input
        type="text"
        bind:value={owner}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Bprint</label>
      <input
        type="text"
        bind:value={bprint_id}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Handlers</label>
      <Kveditor data={{}} bind:getData={getHandlerData} />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
      <Kveditor data={{}} bind:getData={getMetaData} />
    </div>

    <div class="flex py-3">
      <button
        on:click={save}
        class="p-2 bg-blue-400 m-1 w-20 text-white rounded">Save</button
      >
    </div>
  </div>
</div>
