<script lang="ts">
  import { getContext } from "svelte";
  import type { AppService } from "../../../../../../services";
  import Kveditor from "../../../../../common/kveditor.svelte";
  import Layout from "../../layout.svelte";

  export let gid = "";

  const app: AppService = getContext("__app__");

  let name = "";
  let payload = "";
  let type = "";
  let policy = "";

  let getMetaData;

  const save = async () => {
    const uapi = await app.apm.get_user_api();
    const resp = await uapi.user_group_add_auth(gid, {
      name,
      payload,
      type,
      policy,
    });
    if (resp.status !== 200) {
      console.log("Err ", resp);
      return;
    }
    app.navigator.goto_admin_usergroups_page();
  };
</script>

<Layout current_item={"user_groups"}>
  <div class="w-full h-full p-10">
    <div class="bg-white p-2">
      <div class="text-2xl text-indigo-900">Auth Provider</div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Name</label>
        <input
          type="text"
          bind:value={name}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Type</label>
        <input
          type="text"
          bind:value={type}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Payload</label>
        <textarea
          type="text"
          bind:value={payload}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Policy</label>
        <textarea
          type="text"
          bind:value={policy}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
        <Kveditor data={{}} bind:getData={getMetaData} />
      </div>

      <div class="flex justify-end">
        <button
          on:click={save}
          class="p-2 bg-blue-400 hover:bg-blue-600 m-1 w-20 text-white rounded"
          >Save</button
        >
      </div>
    </div>
  </div>
</Layout>
