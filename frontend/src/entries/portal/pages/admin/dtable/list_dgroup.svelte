<script lang="ts">
  import { AutoTable, FloatingAdd } from "../../../../../components";
  import Layout from "../layout.svelte";

  import { getContext } from "svelte";
  import { DynAdminAPI } from "./dtable2";
  import type { AppService } from "../../../../../services";

  let groups = [];
  let sources = [];
  let current_source = "";
  let loaded = false;

  const app: AppService = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  dynapi.load_sources().then((resp) => {
    sources = resp.data;
    if (!sources.length) {
      return;
    }

    current_source = sources[0];
  });

  $: {
    if (current_source !== "") {
      dynapi.load_groups(current_source).then((resp) => {
        groups = resp.data;
        loaded = true;
      });
    }
  }
</script>

<Layout current_item={"dtable"}>
  <div class="bg-white rounded p-2 m-2 flex flex-col">
    <div class="flex justify-end">
      <label class="p-1">
        <span class="uppercase">Sources</span>

        <select class="p-2 border bg-gray-50" bind:value={current_source}>
          {#each sources as source}
            <option>{source}</option>
          {/each}
        </select>
      </label>
    </div>

    {#if loaded}
      <AutoTable
        action_key="slug"
        actions={[
          {
            Class: "bg-green-400",
            Name: "explore",
            Action: (id) => {
              dynapi.goto_dgroup(current_source, id);
            },
          },
          {
            Name: "Edit",
            Action: (grp) => {
              app.navigator.goto_dgroup_edit(current_source, grp);
            },
          },
          {
            Name: "Delete",
            Class: "bg-red-400",
            Action: async (grp) => {
              await dynapi.delete_dgroup(current_source, grp);
              dynapi.load_groups(current_source).then((resp) => {
                groups = resp.data;
                loaded = true;
              });
            },
          },
        ]}
        key_names={[
          ["name", "Name"],
          ["description", "Description"],
          ["source_db", "source"],
        ]}
        datas={groups}
      />
    {/if}
  </div>
</Layout>

<FloatingAdd onClick={() => {}} />
