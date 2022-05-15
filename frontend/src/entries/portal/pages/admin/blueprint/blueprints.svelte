<script lang="ts">
  import { AutoTable, FloatingAdd } from "../../../../../components";
  import Layout from "../layout.svelte";
  import Installer from "./installer/installer.svelte";
  import BprintEditor from "./editor/editor.svelte";
  import type { AppService } from "../../../../../services";
  import { getContext } from "svelte";

  const app: AppService = getContext("__app__");

  let bprints = [];

  const load = async () => {
    const api = await app.apm.get_bprint_api();
    const resp = await api.bprint_list();
    bprints = resp.data;
  };

  load();
</script>

<Layout>
  <AutoTable
    color={["group"]}
    action_key="id"
    actions={[
      {
        Name: "Instance",
        Class: "bg-green-400",
        Action: (bid) => {
          app.big_modal_open(Installer, { bid });
        },
      },
      {
        Name: "Edit",
        Action: app.navigator.goto_admin_bprint_page,
      },
      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: async (id) => {
          const api = await app.apm.get_bprint_api();
          const resp = await api.bprint_remove(id);
        },
      },
    ]}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["slug", "Slug"],
      ["type", "Type"],
      ["sub_type", "Sub Type"],
    ]}
    datas={bprints}
  />
</Layout>

<FloatingAdd onClick={() => app.big_modal_open(BprintEditor, {})} />
