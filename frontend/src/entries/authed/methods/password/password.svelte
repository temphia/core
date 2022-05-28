<script>
  export let emailuser = "";
  export let password = "";
</script>

<div class="mt-4">
  <label
    class="block mb-2 text-sm font-medium text-gray-600"
    for="LoggingEmailAddress">Email Address / Username</label
  >
  <input
    id="LoggingEmailAddress"
    class="block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md focus:border-blue-500 focus:outline-none focus:ring"
    bind:value={emailuser}
    type="email"
  />
</div>

<div class="mt-4">
  <div class="flex justify-between">
    <label
      class="block mb-2 text-sm font-medium text-gray-600"
      for="loggingPassword">Password</label
    >
    <a href="#" class="text-xs text-gray-500 hover:underline"
      >Forget Password?</a
    >
  </div>

  <input
    id="loggingPassword"
    class="block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md focus:border-blue-500 focus:outline-none focus:ring"
    type="password"
    bind:value={password}
  />
</div>

<div class="mt-8">
  <button
    on:click={() => console.log(emailuser, password)}
    class="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-gray-600 rounded hover:bg-gray-900"
  >
    Login
  </button>
</div>
