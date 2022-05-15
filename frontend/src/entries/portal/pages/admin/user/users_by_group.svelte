<script lang="ts">
  import { getContext } from "svelte";
  import { FloatingAdd } from "../../../../../components";
  import type { AppService } from "../../../../../services";
  import Layout from "../layout.svelte";
  import Users from "./users.svelte";

  export let id;

  const app: AppService = getContext("__app__");

  const newUser = app.navigator.goto_admin_new_user_page;
</script>

<Layout current_item={"user_groups"}>
  <Users user_group={id} />
</Layout>

<FloatingAdd onClick={newUser} />
