<script lang="ts">
  import { AutoTable, FloatingAdd } from "../../../../../../components";
  import Layout from "../layout.svelte";
  import type { AppService } from "../../../../../../services";
  import { getContext } from "svelte";

  const app: AppService = getContext("__app__");
  let domains = [];

  const load = async () => {
    const tapi = await app.apm.get_tenant_id();
    const resp = await tapi.list_tenant_domain();
    domains = resp.data;
  };

  load();
</script>

<Layout>
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Edit",
        Action: (id) => {},
      },
      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: async (pid) => {},
      },
    ]}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
    ]}
    datas={domains}
  />
</Layout>

<FloatingAdd onClick={() => {}} />
