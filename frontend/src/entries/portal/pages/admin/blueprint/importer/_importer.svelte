<script lang="ts">
  import { getContext } from "svelte";

  import { CEditor, ToggleButton } from "../../../../../../components";
  import type { AppService } from "../../../../../../services";
  import Installer from "../installer/installer.svelte";
  import ImporterLayout from "../_layout.svelte";

  export let data;
  export let import_func;
  export let source;

  const app: AppService = getContext("__app__");

  let import_opts_target_id = "";
  let import_opts_schema_modified = false;
  let install_next = false;
  let import_done = false;
  let editor;

  $: _all_files = data.ffiles || [];
  $: _skip_files = [];
  let mesage = "importing blueprint...";

  const pick_file = (file) => () => {
    const index = _skip_files.indexOf(file);
    if (index > -1) {
      _skip_files.splice(index, 1);
    } else {
      _skip_files.push(file);
    }
    _skip_files = [..._skip_files];
  };

  const final_func = () => {
    if (mesage.startsWith("error") || !install_next) {
      app.big_modal_close();
    } else {
      app.big_modal_open(Installer, { bid: import_opts_target_id });
    }
  };
  const importer_func = async () => {
    const resp = await import_func({
      source,
      group: data.group,
      slug: data.slug,
      skip_files: _skip_files,
      target_id: import_opts_target_id,
      schema_modified: false, //import_opts_schema_modified,
      schema: "", //editor.getValue(),
    });

    if (resp.status === 200) {
      import_opts_target_id = resp.data;
      mesage = `blueprint imported sucessfully: ${resp.data}`;
    } else {
      mesage = `error while importing blueprint: ${resp.data}`;
    }
    import_done = true;
  };
</script>

<ImporterLayout
  description={data["description"]}
  files={data["files"]}
  type={data["type"]}
  name={data["name"]}
  slug={data["slug"]}
  subtype={data["sub_type"]}
  source=""
  {final_func}
  last_page_func={importer_func}
>
  <div slot="options">
    <div class="w-full shadow rounded bg-white">
      <div
        class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
      >
        <div>
          <h2
            class="text-lg uppercase text-gray-900 font-medium title-font mb-4"
          >
            Slug
          </h2>
          <p
            class="leading-relaxed selection:bg-red-200 text-base text-gray-400"
          >
            It could contain alphanumberic characters and not special
            characters. If not given, it will be generated automatically
          </p>
        </div>

        <div>
          <input
            type="text"
            bind:value={import_opts_target_id}
            class="px-4 py-2 border focus:ring-gray-500 focus:border-gray-900 w-full sm:text-sm border-gray-300 rounded-md focus:outline-none text-gray-600"
            placeholder="Optional"
          />
        </div>
      </div>

      <div
        class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
      >
        <div>
          <p class="text-gray-600 uppercase">Install Next</p>
        </div>

        <div>
          <ToggleButton bind:checked={install_next} label="" />
        </div>
      </div>

      {#if _all_files.length > 0}
        <div class="flex-col flex py-3">
          <label class="leading-loose">Files</label>

          <table class="min-w-full leading-normal border overflow-auto">
            <thead>
              <tr>
                <th
                  class="px-1 py-2 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                />
                <th
                  class="px-5 py-2 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                >
                  Name
                </th>
                <th
                  class="px-5 py-2 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                >
                  File
                </th>
                <th
                  class="px-5 py-2 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                >
                  Action
                </th>
              </tr>
            </thead>
            <tbody>
              {#each _all_files as file}
                <tr class="hover:bg-gray-200">
                  <td
                    class="px-1 py-1 border-b border-gray-200 bg-white text-sm"
                  >
                    <input
                      type="checkbox"
                      class="form-checkbox h-5 w-5 p-2 text-gray-600"
                      checked={!_skip_files.includes(file)}
                      on:click={pick_file(file)}
                    />
                  </td>
                  <td
                    class="px-5 py-1 border-b border-gray-200 bg-white text-sm"
                  >
                    <div class="flex items-center">
                      <div class="flex-shrink-0 w-10 h-10">
                        <div class="w-full h-full rounded-full">
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            class="h-10 w-10 text-gray-500"
                            viewBox="0 0 20 20"
                            fill="currentColor"
                          >
                            <path
                              fill-rule="evenodd"
                              d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z"
                              clip-rule="evenodd"
                            />
                          </svg>
                        </div>
                      </div>
                      <div class="ml-3">
                        <p class="text-gray-900 whitespace-no-wrap">
                          {file}
                        </p>
                      </div>
                    </div>
                  </td>
                  <td
                    class="px-5 py-1 border-b border-gray-200 bg-white text-sm"
                  >
                    <p class="text-gray-700">{file}</p>
                  </td>
                  <td
                    class="px-5 py-1 border-b border-gray-200 bg-white text-sm"
                  >
                    <span
                      class="relative inline-block px-3 py-1 font-semibold text-red-900 leading-tight"
                    >
                      <span
                        aria-hidden
                        class="absolute inset-0 bg-red-200 opacity-50 rounded-full"
                      />
                      <a class="relative" href="#">Preview</a>
                    </span>
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      {/if}
    </div>
  </div>

  <div slot="final">
    <div class="flex items-center justify-center w-full h-full">
      <div
        class="flex justify-center items-center space-x-1 text-lg text-gray-700"
      >
        {#if !import_done}
          <svg
            fill="none"
            class="w-20 h-20 animate-spin"
            viewBox="0 0 32 32"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              clip-rule="evenodd"
              d="M15.165 8.53a.5.5 0 01-.404.58A7 7 0 1023 16a.5.5 0 011 0 8 8 0 11-9.416-7.874.5.5 0 01.58.404z"
              fill="currentColor"
              fill-rule="evenodd"
            />
          </svg>
        {/if}
        <div>{mesage}</div>
      </div>
    </div>
  </div>
</ImporterLayout>
