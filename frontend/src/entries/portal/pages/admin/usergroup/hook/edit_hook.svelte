<script lang="ts">
  import { getContext } from "svelte";
  import type { AppService } from "../../../../../../services";
  import { KvEditor } from "../../../../../common";
  import Layout from "../../layout.svelte";

  export let gid = "";
  export let id = "";

  let data;
  let mod_data = {};

  const app: AppService = getContext("__app__");

  let getMetaData;
  let meta_modified;
  let modified;

  const modifyField = (field: string) => (ev) => {
    modified = true;
    mod_data[field] = ev.target["value"];
  };

  const load = async () => {
    const uapi = await app.apm.get_user_api();
    const resp = await uapi.user_group_get_hook(gid, Number(id));
    if (resp.status !== 200) {
      console.log("Err ", resp);
      return;
    }
    data = resp.data;
  };

  const save = async () => {
    const uapi = await app.apm.get_user_api();
    const resp = await uapi.user_group_update_hook(gid, Number(id), mod_data);
    if (resp.status !== 200) {
      console.log("Err ", resp);
      return;
    }
    app.navigator.goto_admin_usergroup_page(gid);
  };

  load();
</script>

<Layout current_item={"user_groups"} loading={data === null}>
  <div class="w-full h-full p-10">
    <div class="bg-white p-2">
      {#if data}
        <div class="text-2xl text-indigo-900">Edit Hook</div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Stage Name</label>

          <input
            type="text"
            value={data["stage_name"] || ""}
            on:change={modifyField("stage_name")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Plug</label>
          <input
            type="text"
            value={data["plug_id"] || ""}
            on:change={modifyField("plug_id")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Agent</label>
          <textarea
            value={data["agent_id"] || ""}
            on:change={modifyField("agent_id")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Autoload</label>
          <input
            checked={data["autoload"] || false}
            on:change={(ev) => {
              modified = true;
              mod_data["autoload"] = ev.target["checked"];
            }}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
          <KvEditor
            data={data["extra_meta"] || {}}
            bind:getData={getMetaData}
            bind:modified={meta_modified}
          />
        </div>

        <div class="flex justify-end">
          <button
            on:click={save}
            class="p-2 bg-blue-400 hover:bg-blue-600 m-1 w-20 text-white rounded"
            >Save</button
          >
        </div>
      {/if}
    </div>
  </div>
</Layout>
