<script>
  import Modal from "svelte-simple-modal";
  import { Router, Route } from "svelte-routing";
  import MainLayout from "./layout/layout.svelte";
  import BigModal from "../../components/modal/big.svelte";
  import {
    UserProfile,
    SelfProfile,
    Start,
    CabinetLoader,
    CabinetSource,
    CabinetFolder,
    CabinetFile,
    DgroupLoader,
    Dgroups,
    Dgroup,
    Dtable,
    Store,
    StoreItem,
    AdminPlugs,
    AdminPlug,
    AdminPlugAgents,
    AdminPlugAgent,
    AdminAgentResources,
    AdminResources,
    AdminNewResource,
    AdminEditResource,
    AdminUsers,
    AdminUser,
    AdminBrint,
    AdminBrints,
    AdminUserGroups,
    AdminUserGroup,
    AdminNewUser,
    AdminUserByGroup,
    AdminNewUserGroup,
    AdminTenant,
    ListColumn,
    ListDgroup,
    ListDtable,
    EditColumn,
    EditDgroup,
    EditDtable,
    ListViews,
    ViewEdit,
    ViewNew,
    ListHooks,
    HookEdit,
    HookNew,
    UserGroupAuthNew,
    UserGroupAuthEdit,
    UserGroupPlugNew,
    UserGroupPlugEdit,
    UserGroupHookNew,
    UserGroupHookEdit,
    UserGroupDataNew,
    UserGroupDataEdit,
    AdminDtableBuilder,
    AboutTenant,
    AdminTenantEdit,
    Launcher,
    AppsLauncher,
  } from "./pages";
  import Playground from "./playground/playground.svelte";
  import Tailwind from "../common/_tailwind.svelte";
  import EditGroup from "./pages/admin/usergroup/_edit_group.svelte";

  let url = "/console";
</script>

<Router {url}>
  <Modal>
    <MainLayout>
      <BigModal />

      <Route path="/console" component={Start} />
      <Route path="/console/playground" component={Playground} />

      <!-- dtable -->
      <Route path="/console/dtable_load" component={DgroupLoader} />
      <Route path="/console/dtable/:source" let:params>
        <Dgroups source={params.source} />
      </Route>
      <Route path="/console/dtable/:source/:dgroup" let:params>
        <Dgroup source={params.source} dgroup={params.dgroup} />
      </Route>
      <Route path="/console/dtable/:source/:dgroup/:dtable" let:params>
        <Dtable
          source={params.source}
          dgroup={params.dgroup}
          dtable={params.dtable}
        />
      </Route>

      <!-- plugs -->

      <!-- <Route path="/console/plugs" component={Plugs} />
      <Route path="/console/plug/:plugid" let:params>
        <Plug plugid={params.plugid} />
      </Route> -->

      <!-- store stuff -->
      <Route path="/console/store" component={Store} />
      <Route path="/console/store/:source/:group/:item" let:params>
        <StoreItem
          item={params.item}
          source={params.source}
          group={params.group}
        />
      </Route>

      <Route path="/console/launcher/:plugid/:agentid" let:params>
        <Launcher plugid={params.plugid} agentid={params.agentid} />
      </Route>

      <Route
        path="/console/apps_launcher"
        let:params
        component={AppsLauncher}
      />

      <!-- Admin -->

      <Route path="/console/admin/bprints" component={AdminBrints} />
      <Route path="/console/admin/bprints/:id" let:params>
        <AdminBrint bid={params.id} />
      </Route>

      <Route path="/console/admin/plugs" component={AdminPlugs} />
      <Route path="/console/admin/plugs/:id" let:params>
        <AdminPlug id={params.id} />
      </Route>

      <Route path="/console/admin/plugs/:pid/agents" let:params>
        <AdminPlugAgents pid={params.pid} aid={params.aid} />
      </Route>
      <Route path="/console/admin/plugs/:pid/agents/:aid" let:params>
        <AdminPlugAgent pid={params.pid} aid={params.aid} />
      </Route>

      <Route path="/console/admin/plugs/:pid/agents/:aid/resources" let:params>
        <AdminAgentResources pid={params.pid} aid={params.aid} />
      </Route>

      <Route path="/console/admin/resources" component={AdminResources} />
      <Route path="/console/admin/resources/new" let:params>
        <AdminNewResource />
      </Route>

      <Route path="/console/admin/resources/edit/:id" let:params>
        <AdminEditResource id={params.id} />
      </Route>

      <Route path="/console/admin/users" component={AdminUsers} />
      <Route path="/console/admin/users/:id" let:params>
        <AdminUser id={params.id} />
      </Route>

      <Route path="/console/admin/new_user" component={AdminNewUser} />

      <Route path="/console/admin/users_by_group/:id" let:params>
        <AdminUserByGroup id={params.id} />
      </Route>

      <Route path="/console/admin/user_groups" component={AdminUserGroups} />
      <Route path="/console/admin/user_groups/:id" let:params>
        <AdminUserGroup id={params.id} />
      </Route>

      <Route
        path="/console/admin/new_user_groups"
        component={AdminNewUserGroup}
      />

      <Route path="/console/admin/dtable" let:params component={ListDgroup} />
      <Route path="/console/admin/dtable/:source/:group" let:params>
        <ListDtable source={params.source} group={params.group} />
      </Route>
      <Route path="/console/admin/dtable/:source/:group/:table" let:params>
        <ListColumn
          source={params.source}
          group={params.group}
          table={params.table}
        />
      </Route>

      <Route path="/console/admin/table_views/:source/:group/:table" let:params>
        <ListViews
          source={params.source}
          group={params.group}
          table={params.table}
        />
      </Route>

      <Route
        path="/console/admin/table_views/:source/:group/:table/new"
        let:params
      >
        <ViewNew
          source={params.source}
          group={params.group}
          table={params.table}
        />
      </Route>

      <Route
        path="/console/admin/table_views/:source/:group/:table/edit/:id"
        let:params
      >
        <ViewEdit
          source={params.source}
          group={params.group}
          table={params.table}
          id={params.id}
        />
      </Route>

      <Route path="/console/admin/table_hooks/:source/:group/:table" let:params>
        <ListHooks
          source={params.source}
          group={params.group}
          table={params.table}
        />
      </Route>

      <Route
        path="/console/admin/table_hooks/:source/:group/:table/new"
        let:params
      >
        <HookNew
          source={params.source}
          group={params.group}
          table={params.table}
        />
      </Route>

      <Route
        path="/console/admin/table_hooks/:source/:group/:table/edit/:id"
        let:params
      >
        <HookEdit
          source={params.source}
          group={params.group}
          table={params.table}
          id={params.id}
        />
      </Route>

      <Route path="/console/admin/user_group_auth/:gid/new" let:params>
        <UserGroupAuthNew gid={params.gid} />
      </Route>
      <Route path="/console/admin/user_group_auth/:gid/edit/:id" let:params>
        <UserGroupAuthEdit id={params.id} gid={params.gid} />
      </Route>

      <Route path="/console/admin/user_group_hook/:gid/new" let:params>
        <UserGroupHookNew gid={params.gid} />
      </Route>
      <Route path="/console/admin/user_group_hook/:gid/edit/:id" let:params>
        <UserGroupHookEdit id={params.id} gid={params.gid} />
      </Route>

      <Route path="/console/admin/user_group_plug/:gid/new" let:params>
        <UserGroupPlugNew gid={params.gid} />
      </Route>
      <Route path="/console/admin/user_group_plug/:gid/edit/:id" let:params>
        <UserGroupPlugEdit id={params.id} gid={params.gid} />
      </Route>

      <Route path="/console/admin/user_group_data/:gid/new" let:params>
        <UserGroupDataNew gid={params.gid} />
      </Route>
      <Route path="/console/admin/user_group_data/:gid/edit/:id" let:params>
        <UserGroupDataEdit id={params.id} gid={params.gid} />
      </Route>

      <Route path="/console/admin/dtable_edit/:source/:group" let:params>
        <EditDgroup source={params.source} group={params.group} />
      </Route>

      <Route
        path="/console/admin/builder/builder"
        component={AdminDtableBuilder}
      />

      <Route path="/console/admin/builder/:bid" let:params>
        <AdminDtableBuilder bid={params.bid} />
      </Route>

      <Route path="/console/admin/dtable_edit/:source/:group/:table" let:params>
        <EditDtable
          source={params.source}
          group={params.group}
          table={params.table}
        />
      </Route>

      <Route
        path="/console/admin/dtable_edit/:source/:group/:table/:column"
        let:params
      >
        <EditColumn
          source={params.source}
          group={params.group}
          table={params.table}
          column={params.column}
        />
      </Route>

      <Route path="/console/admin/ns" component={AdminTenant} />
      <Route path="/console/admin/ns_edit" component={AdminTenantEdit} />
      <Route path="/console/about_ns" component={AboutTenant} />

      <Route path="/console/cabinet_load" component={CabinetLoader} />
      <Route path="/console/cabinet/:source" let:params>
        <CabinetSource source={params.source} />
      </Route>
      <Route path="/console/cabinet/:source/:folder" let:params>
        <CabinetFolder source={params.source} folder={params.folder} />
      </Route>

      <Route path="/console/cabinet/:source/:folder/:file" let:params>
        <CabinetFile
          source={params.source}
          folder={params.folder}
          file={params.file}
        />
      </Route>

      <Route path="/console/store" component={CabinetLoader} />
      <Route path="/console/self_profile" component={SelfProfile} />

      <Route path="/console/user_profile/:user" let:params>
        <UserProfile user={params.user} />
      </Route>
    </MainLayout>
  </Modal>
</Router>

<Tailwind />
