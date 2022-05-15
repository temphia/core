<script lang="ts">
  import { CEditor } from "../../../../../components";
  export let name = "";
  export let data;
  export let mod_data;
  export let modified = false;
  export let onSave;
  export let schema = {
    primary: "id",
    text: [],
    object: [],
    bool: [],
  };

  let editor;
  let mode_raw = false;
  let collapsed = false;

  const toggle = () => {
    collapsed = !collapsed;
  };

  const extract_editor = () => {
    if (editor && modified) {
      try {
        const value = JSON.parse(editor.getValue());
        if (typeof value !== "object" || !value) {
          return;
        }
        data = { ...data, ...value };
      } catch (error) {}
    }
  };

  const toggleTab = () => {
    if (mode_raw) {
      extract_editor();
    }
    mode_raw = !mode_raw;
  };

  const get = (name) => data[name] || "";
  const getBool = (name) => data[name] || false;

  const getObject = (name) => {
    const inner = data[name];
    if (!inner) {
      return "";
    }
    return JSON.stringify(inner, null, 4) || "{}";
  };

  const set = (name) => (ev) => {
    mod_data = { ...mod_data, [name]: ev.target.value };
    modified = true;
  };

  const setBool = (name) => (ev) => {
    mod_data = { ...mod_data, [name]: ev.target.checked };
    modified = true;
  };
</script>

<div class="w-full h-full flex">
  <div class="absolute bottom-1 z-10 p-1">
    <button
      on:click={toggle}
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 rounded"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 6h16M4 12h16M4 18h7"
        />
      </svg>
    </button>
  </div>
  {#if !collapsed}
    <div class="flex-none w-64 flex flex-col bg-white border-r">
      <div class="flex items-center justify-center h-14 border-b">
        <div>{name}</div>
      </div>
      <div class="overflow-y-auto flex flex-col overflow-x-hidden">
        <ul class="flex flex-col py-4 space-y-1">
          <li class="px-5">
            <div class="flex flex-row items-center h-8">
              <div class="text-sm font-light tracking-wide text-gray-500">
                System
              </div>

              {#if modified}
                <span
                  class="px-2 py-0.5 ml-auto text-xs font-medium tracking-wide text-red-500 bg-red-50 rounded-full"
                  >Modified</span
                >
              {/if}
            </div>
          </li>
          <li>
            <a
              href="#"
              on:click={toggleTab}
              class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-indigo-500 pr-6  {!mode_raw
                ? 'border-indigo-600'
                : ''}"
            >
              <span class="inline-flex justify-center items-center ml-4">
                <svg
                  class="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                  ><path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
                  /></svg
                >
              </span>
              <span class="ml-2 text-sm tracking-wide truncate">Common</span>
            </a>
          </li>
          <li>
            <a
              href="#"
              on:click={toggleTab}
              class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-indigo-500 pr-6  {mode_raw
                ? 'border-indigo-600'
                : ''}"
            >
              <span class="inline-flex justify-center items-center ml-4">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-6 w-6"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"
                  />
                </svg>
              </span>
              <span class="ml-2 text-sm tracking-wide truncate">Raw</span>
            </a>
          </li>
        </ul>
      </div>
    </div>
  {/if}
  <div style="width: inherit;">
    {#if mode_raw}
      <CEditor
        code={JSON.stringify(data, null, 4) || "{}"}
        bind:editor
        container_style="height:100vh;"
        on:change={() => {
          modified = true;
        }}
      />
    {:else}
      <div class=" h-full w-full bg-indigo-100 p-10 overflow-auto">
        <div class="p-5 bg-white w-full ">
          <div class="text-2xl text-indigo-900">Plug</div>
          <div class="flex-col flex py-3">
            <label class="pb-2 text-gray-700 font-semibold">Id</label>
            <input
              type="text"
              value={get(schema.primary)}
              disabled={get(schema.primary) !== ""}
              on:change={set(schema.primary)}
              class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
              placeholder="idxx"
            />
          </div>

          {#each schema.text || [] as item}
            <div class="flex-col flex py-3 relative">
              <label for="" class="pb-2 text-gray-700 font-semibold capitalize"
                >{item}</label
              >
              <input
                type="text"
                value={get(item)}
                on:change={set(item)}
                class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
              />
            </div>
          {/each}

          {#each schema.object || [] as item}
            <div class="flex-col flex py-3 relative">
              <label for="" class="pb-2 text-gray-700 font-semibold capitalize"
                >{item}</label
              >
              <textarea
                type="text"
                disabled
                value={getObject(item)}
                class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
              />
            </div>
          {/each}

          {#each schema.bool || [] as item}
            <div class="flex-col flex py-3 relative">
              <label for="" class="pb-2 text-gray-700 font-semibold capitalize"
                >{item}</label
              >
              <input
                type="checkbox"
                value={getBool(item)}
                on:change={setBool(item)}
                class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
              />
            </div>
          {/each}

          <div class="flex justify-end py-3">
            <button
              on:click={onSave}
              class="p-2 bg-blue-400 hover:bg-blue-600 m-1 w-20 text-white rounded"
              >Save</button
            >
          </div>
        </div>
      </div>
    {/if}
  </div>
</div>
