<script lang="ts">
  import { getContext } from "svelte";

  import { AutoTable, FloatingAdd } from "../../../../../components";
  import type { AppService } from "../../../../../services";
  import Layout from "../layout.svelte";
  import { DynAdminAPI } from "./dtable2";
  export let source;
  export let group;
  export let table;

  const app: AppService = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let columns = [];

  dynapi.load_tables_column(source, group, table).then((resp) => {
    columns = resp.data;
  });
</script>

<Layout current_item={"dtable"}>
  <AutoTable
    action_key="slug"
    actions={[
      {
        Name: "Edit",
        Action: (col) => {
          app.navigator.goto_column_edit(source, group, table, col);
        },
      },

      {
        Name: "Delete",
        Action: null,
        Class: "bg-red-400",
      },
    ]}
    key_names={[
      ["name", "Name"],
      ["slug", "Slug"],
      ["ctype", "Column type"],
      ["description", "Description"],
    ]}
    color={["ctype"]}
    datas={columns}
  />
</Layout>

<FloatingAdd onClick={() => {}} />
