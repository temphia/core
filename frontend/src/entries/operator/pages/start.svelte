<script lang="ts">
  const data = {
    cluster_name: "xyz",
    coredb: { vendor: "postgres", name: "Xyz" },
    dyndbs: [
      { vendor: "postgres", name: "Xyz" },
      { vendor: "postgres", name: "Xyz" },
    ],
    repo_sources: [],
    cabinet_sources:[
      { vendor: "local_fs", name: "local" },
    ],
    nodes: [
      {
        node_id: "ax8y2bi2",
        mem_max: "1024000",
        mem_used: "89000",
        cpu_max: "40%",
        tags: [],
      },
      {
        node_id: "aw6y2byr",
        mem_max: "1024000",
        mem_used: "89000",
        cpu_max: "40%",
        tags: [],
      },
    ],
  };
</script>

<div class="h-full w-full flex-1 flex flex-col overflow-hidden">
  <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-200">
    <div class="container mx-auto px-6 py-8">
      <h3 class="text-gray-700 text-3xl font-medium">Home</h3>

      <div class="flex flex-col mt-8 space-y-4">
        <div class="w-full bg-white rounded p-2">
          <table class="w-full border">
            <thead>
              <tr
                class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal"
              >
                <th class="py-3 px-6 text-left">Properties</th>
                <th class="py-3 px-6 text-left">Value</th>
              </tr>
            </thead>
            <tbody class="text-gray-600 text-sm font-light">
              <tr class="border-b border-gray-200 hover:bg-gray-100">
                <th class="py-3 px-6 text-left">Cluster Name</th>
                <th class="py-3 px-6 text-left">{data["cluster_name"]}</th>
              </tr>

              <tr class="border-b border-gray-200 hover:bg-gray-100">
                <td class="py-3 px-6 text-left">Core Database</td>
                <td class="py-3 px-6 text-left">
                  <div>
                    Vendor: {data["coredb"]["vendor"]}
                  </div>
                  <div>
                    Name: {data["coredb"]["name"]}
                  </div>
                </td>
              </tr>
              <tr class="border-b border-gray-200 hover:bg-gray-100">
                <td class="py-3 px-6 text-left">Nodes</td>
                <td class="py-3 px-6 text-left">
                  {#each data["nodes"] as node}
                    <div class="p-2 border">
                      <div>
                        Node Id: {node["node_id"]}
                      </div>
                      <div>
                        CPU Max: {node["cpu_max"]}
                      </div>
                      <div>
                        Mem Max: {node["mem_max"]}
                      </div>
                      <div>
                        Mem Used: {node["mem_used"]}
                      </div>
                      <div>
                        Tags: {#each node["tags"] || [] as tag} {tag}{/each}
                      </div>
                    </div>
                  {/each}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </main>
</div>
