<script lang="ts">
  import { getContext } from "svelte";

  import { AutoTable, FloatingAdd } from "../../../../../../components";
  import type { AppService } from "../../../../../../services";
  import Layout from "../../layout.svelte";

  export let pid;
  export let aid;

  let resources = [];

  const app: AppService = getContext("__app__");

  app.apm.get_resource_api().then(async (rapi) => {
    const resp = await rapi.resource_list();
    resources = resp.data;
  });
</script>

<Layout>
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Edit",
        Action: (id) => app.navigator.goto_admin_resource_edit(id),
      },
      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: null,
      },
    ]}
    key_names={[
      ["name", "Name"],
      ["id", "Id"],
      ["type", "Type"],
      ["schema", "Schema"],
    ]}
    datas={resources}
  />
</Layout>

<FloatingAdd
  onClick={() => {
    app.navigator.goto_admin_resource_new(pid);
  }}
/>
