<script lang="ts">
  import type { AppService } from "../../../../../services";
  export let app: AppService;

  let plugs = [
    {
      id: "xyz",
      name: "Abc",
      executor: "mno.pqr",
      live: true,
      dev: true,
      owner: "",
      bprint_id: "gwguihqwuq",
    },
  ];

  let agents = [
    {
      id: "xyz",
      name: "Abc",
      type: "",
      plug_id: "xyz",
    },
  ];

  let selected_plug = "";
  let selected_agent = "";

  const load = async () => {
    const papi = await app.apm.get_plug_api();
    const resp = await papi.list_plug();
    plugs = resp.data;
  };

  load();
</script>

<div class="flex flex-col w-full">
  <h3 class="mb-3 text-xl font-medium text-gray-900 dark:text-white">Plugs</h3>

  <table
    class="w-full border text-sm text-left text-gray-500 dark:text-gray-400 overflow-auto"
  >
    <thead
      class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
    >
      <tr>
        <th scope="col" class="p-1" />
        <th scope="col" class="px-2 py-1"> Id </th>
        <th scope="col" class="px-2 py-1"> Name </th>
        <th scope="col" class="px-2 py-1"> Executor </th>
        <th scope="col" class="px-2 py-1"> Owner</th>
        <th scope="col" class="px-2 py-1"> Bprint Id</th>
      </tr>
    </thead>
    <tbody>
      {#each plugs as plug}
        <tr
          class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
        >
          <td class="w-4 p-4">
            <div class="flex items-center">
              <input
                id="checkbox-table-search-1"
                type="checkbox"
                class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
              />
              <label for="checkbox-table-search-1" class="sr-only"
                >checkbox</label
              >
            </div>
          </td>
          <th
            scope="row"
            class="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap"
          >
            {plug.id}
          </th>
          <td class="px-6 py-4"> {plug.name} </td>
          <td class="px-6 py-4"> {plug.executor} </td>
          <td class="px-6 py-4"> {plug.owner} </td>
          <td class="px-6 py-4"> {plug.bprint_id} </td>
        </tr>
      {/each}
    </tbody>
  </table>

  <h3 class="mb-3 text-xl font-medium text-gray-900 dark:text-white">Agents</h3>

  <table
    class="w-full border text-sm text-left text-gray-500 dark:text-gray-400"
  >
    <thead
      class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
    >
      <tr>
        <th scope="col" class="p-1" />
        <th scope="col" class="px-2 py-1"> Id </th>
        <th scope="col" class="px-2 py-1"> Name </th>
        <th scope="col" class="px-2 py-1"> Type </th>
        <th scope="col" class="px-2 py-1"> Plug Id</th>
      </tr>
    </thead>
    <tbody>
      {#each agents as agent}
        <tr
          class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
        >
          <td class="w-4 p-4">
            <div class="flex items-center">
              <input
                id="checkbox-table-search-1"
                type="checkbox"
                class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
              />
              <label for="checkbox-table-search-1" class="sr-only"
                >checkbox</label
              >
            </div>
          </td>
          <th
            scope="row"
            class="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap"
          >
            {agent.id}
          </th>
          <td class="px-6 py-4"> {agent.name} </td>
          <td class="px-6 py-4"> {agent.type} </td>
          <td class="px-6 py-4"> {agent.plug_id} </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
