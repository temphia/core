<script>
  import Icon from "./icon.svelte";
  import SvelteMarkdown from "svelte-markdown";

  export let data = {
    id: "",
    name: "",
    slug: "",
    group: "",
    subgroup: "",
    description: ``,
    screenshots: [],
    icon: "",
    schema: "",
    meta: ``,
    source: "",
    tenant_id: "",
    tags: [],
    files: {},
  };

  export let readme = ``;

  export let actions = [];
</script>

<div class="h-full w-full bg-gray-200 overflow-auto">
  <div class="w-full mt-10">
    <div
      class="m-auto flex flex-col p-4 bg-white shadow rounded"
      style="max-width: 50rem;"
    >
      <div class="flex flex-wrap  border-b">
        <div class="w-1/4">
          <h2
            class="m-1 text-xs font-semibold text-lime-500 uppercase title-font"
          >
            BLUEPRINT
          </h2>
          <div class="w-32 p-2 bg-gray-100">
            <Icon src={data.icon} />
          </div>
        </div>
        <div class="w-3/4 flex flex-col ml-5 mb-4">
          <h1 class="text-gray-700 text-3xl font-medium">{data.name}</h1>
          <p
            class="mb-8 text-base leading-relaxed text-left text-blueGray-600 "
          >
            {data.description}
          </p>
        </div>
      </div>

      <div class="m-auto">
        <SvelteMarkdown source={readme} />
      </div>

      <div class="flex justify-end items-center">
        <div class="space-5">
          {#each actions as action}
            <button
              on:click={() => action.Action()}
              class="p-2 m-1 text-white text-xs inline-flex text-semibold drop-shadow-md rounded transform uppercase hover:scale-110 {action.Class ||
                'bg-blue-400'}"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5 mr-1"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M8 4H6a2 2 0 00-2 2v12a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-2m-4-1v8m0 0l3-3m-3 3L9 8m-5 5h2.586a1 1 0 01.707.293l2.414 2.414a1 1 0 00.707.293h3.172a1 1 0 00.707-.293l2.414-2.414a1 1 0 01.707-.293H20"
                />
              </svg>
              {" " + action.Name}
            </button>
          {/each}
        </div>
      </div>
    </div>
  </div>

  <div class="w-full m-5 " />
</div>
