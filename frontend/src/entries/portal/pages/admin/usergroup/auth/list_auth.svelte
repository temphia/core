<script lang="ts">
  import { AutoTable } from "../../../../../../components";
  import type { AppService } from "../../../../../../services";

  export let gid = "";
  export let app: AppService;

  let data = [];
  const load = async () => {
    const uapi = await app.apm.get_user_api();
    const resp = await uapi.user_group_list_auth(gid);
    data = resp.data;
  };

  load();
</script>

<AutoTable
  action_key="id"
  actions={[
    {
      Name: "Edit",
      Action: (id) => {
        app.navigator.goto_admin_user_auth_edit(gid, id);
      },
    },
    {
      Name: "Delete",
      Class: "bg-red-400",
      Action: async (id) => {
        const uapi = await app.apm.get_user_api();
        const resp = await uapi.user_group_remove_auth(gid, Number(id));
        if (resp.status !== 200) {
          console.log("Err ", resp);
          return;
        }
        app.navigator.goto_admin_users_page();
      },
    },
  ]}
  key_names={[
    ["id", "Id"],
    ["name", "Name"],
    ["type", "Type"],
  ]}
  datas={data}
/>
