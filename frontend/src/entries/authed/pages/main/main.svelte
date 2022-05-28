<script lang="ts">
  import Password from "../../methods/password/password.svelte";

  import AuthLayout from "./layout.svelte";

  const auth_methods = [
    {
      name: "password",
      type: "password",
    },
    {
      name: "Login With Google",
      type: "oauth",
      provider: "google",
      url: "",
    },
  ];
</script>

<AuthLayout>
  <div class="w-full px-6 py-8">
    <h2 class="text-2xl font-semibold text-center text-gray-700 mb-5">Temphia User Login</h2>
    <p class="text-xl text-center text-gray-600 mb-5">Select preferred method</p>

    <select class="w-full p-2 rounded">
      {#each auth_methods as am}
        <option> {am.name}</option>
      {/each}
    </select>

    <Password />

    <div class="flex items-center justify-between mt-4">
      <span class="w-1/5 border-b md:w-1/4" />

      <a href="#" class="text-xs text-gray-500 uppercase hover:underline"
        >or sign up</a
      >

      <span class="w-1/5 border-b md:w-1/4" />
    </div>
  </div>
</AuthLayout>
