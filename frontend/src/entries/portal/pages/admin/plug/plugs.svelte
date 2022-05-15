<script lang="ts">
  import { AutoTable, FloatingAdd } from "../../../../../components";
  import Layout from "../layout.svelte";
  import PlugEditor from "./_plug_editor.svelte";
  import type { AppService } from "../../../../../services";
  import { getContext } from "svelte";

  const app: AppService = getContext("__app__");
  let plugs = [];

  const load = async () => {
    const papi = await app.apm.get_plug_api();
    const resp = await papi.list_plug();
    plugs = resp.data;
  };

  load()

</script>

<Layout>
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Execute",
        Class: "bg-blue-400",
        Action: (pid) => {
          app.navigator.iframe_plug_launch(pid, "default");
        },
      },
      {
        Name: "Agents",
        Class: "bg-green-400",
        Action: (id) => {
          app.navigator.goto_admin_agents_page(id);
        },
      },
      {
        Name: "Edit",
        Action: (id) => {
          app.navigator.goto_admin_plug_page(id);
        },
      },
      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: async (pid) => {
          const papi = await app.apm.get_plug_api();
          await papi.del_plug(pid);
          load();
        },
      },
    ]}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["executor", "Executor"],
      ["owner", "Owner"],
      ["bprint_id", "Bprint Id"],
    ]}
    color={["executor"]}
    datas={plugs}
  />
</Layout>

<FloatingAdd onClick={() => app.big_modal_open(PlugEditor, {})} />
