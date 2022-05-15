<script lang="ts">
  import { getContext } from "svelte";
  import { AutoTable, FloatingAdd } from "../../../../../components";
  import type { AppService } from "../../../../../services";
  import Layout from "../layout.svelte";

  let resources = [
    // {
    //   namespace: "demo3",
    //   id: "c6vd0ttmecapm6na51g0",
    //   name: "External Ping",
    //   type: "slot",
    // },
  ];

  const app: AppService = getContext("__app__");

  const load = async () => {
    const rapi = await app.apm.get_resource_api();
    const resp = await rapi.resource_list();
    resources = resp.data;
  };
  load();
</script>

<Layout current_item="resources">
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
        Action: async (id) => {
          const rapi = await app.apm.get_resource_api();
          await rapi.resource_remove(id);
          load();
        },
      },
    ]}
    key_names={[
      ["name", "Name"],
      ["id", "Id"],
      ["type", "Type"],
      ["schema", "Schema"],
    ]}
    color={["type"]}
    datas={resources}
  />
</Layout>

<FloatingAdd onClick={app.navigator.goto_admin_resource_new} />
