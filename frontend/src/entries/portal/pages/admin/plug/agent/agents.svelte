<script lang="ts">
  import { getContext } from "svelte";
  import { AutoTable, FloatingAction } from "../../../../../../components";
  import type { AppService } from "../../../../../../services";

  import Layout from "../../layout.svelte";

  export let pid;
  let agents = [];

  const app: AppService = getContext("__app__");

  const load = async () => {
    const papi = await app.apm.get_plug_api();
    const resp = await papi.list_agent(pid);
    agents = resp.data;
  };

  load();
</script>

<Layout>
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Edit",
        Action: async (aid) => app.navigator.goto_admin_agent_page(pid, aid),
      },
      {
        Name: "Extern Execute",
        Class: "bg-green-400",
        Action: (agent_id) => {
          app.navigator.extern_plug_launch(pid, agent_id);
        },
      },
      {
        Name: "Iframe Execute",
        Class: "bg-green-400",
        Action: (agent_id) => {
          app.navigator.iframe_plug_launch(pid, agent_id);
        },
      },
      {
        Name: "Resources",
        Class: "bg-lime-400",
        Action: (agent_id) => {
          app.navigator.goto_admin_agent_resources_page(pid, agent_id);
        },
      },

      {
        Name: "Connections",
        Class: "bg-lime-400",
        Action: null,
      },
      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: null,
      },
    ]}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["type", "Type"],
      ["plug_id", "Plug Id"],
    ]}
    datas={agents}
  />
</Layout>

<FloatingAction />
