<script lang="ts">
  import Tailwind from "../common/_tailwind.svelte";
  import { doLogin } from "./auth";
  import AuthLayout from "./auth_layout.svelte";

  export let emailuser = "";
  export let password = "";
</script>

<AuthLayout>
  <div class="w-full px-6 py-8">
    <h2 class="text-2xl font-semibold text-center text-gray-700">Temphia</h2>

    <p class="text-xl text-center text-gray-600">User Login</p>

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
        on:click={() => doLogin(emailuser, password)}
        class="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-gray-600 rounded hover:bg-gray-900"
      >
        Login
      </button>
    </div>

    <div class="flex items-center justify-between mt-4">
      <span class="w-1/5 border-b lg:w-1/4" />

      <a
        href="#"
        class="text-xs text-center text-gray-500 uppercase hover:underline"
        >or</a
      >

      <span class="w-1/5 border-b lg:w-1/4" />
    </div>

    <a
      href="#"
      class="flex items-center justify-center mt-4 text-gray-600 rounded-lg shadow-md hover:bg-gray-200"
    >
      <div class="px-4 py-3" />

      <span class="w-5/6 px-4 py-3 font-bold text-center"
        >Other Sign In Options
      </span>
    </a>

    <div class="flex items-center justify-between mt-4">
      <span class="w-1/5 border-b md:w-1/4" />

      <a href="#" class="text-xs text-gray-500 uppercase hover:underline"
        >or sign up</a
      >

      <span class="w-1/5 border-b md:w-1/4" />
    </div>
  </div>
</AuthLayout>

<Tailwind />
