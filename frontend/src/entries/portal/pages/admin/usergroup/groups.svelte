<script lang="ts">
  import { getContext } from "svelte";
  import Layout from "../layout.svelte";
  import { AutoTable, FloatingAdd } from "../../../../../components";
  import type { AppService } from "../../../../../services";

  let groups = [];
  const app: AppService = getContext("__app__");

  const load = async () => {
    const api = await app.apm.get_user_api();
    const resp = await api.list_user_group();
    groups = resp.data;
  };

  load();
</script>

<Layout current_item={"user_groups"}>
  <AutoTable
    color={["slug"]}
    action_key="slug"
    actions={[
      {
        Name: "Explore",
        Class: "bg-green-400",
        Action: app.navigator.goto_admin_user_by_group,
      },
      {
        Name: "Edit",
        Action: async (id) => app.navigator.goto_admin_usergroup_page(id),
      },
      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: async (id) => {
          const api = await app.apm.get_user_api();
          api.remove_user_group(id);
        },
      },
    ]}
    key_names={[
      ["name", "Name"],
      ["slug", "Slug"],
    ]}
    datas={groups}
  />
</Layout>

<FloatingAdd onClick={app.navigator.goto_admin_new_usergroup_page} />
