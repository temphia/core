<script lang="ts">
    import Layout from "../layout.svelte";
</script>

<Layout>
    <div class="h-full w-full overflow-auto">
        <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
            <div
                class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg"
            >
                <div class="text-2xl text-indigo-900">Organization</div>

                <div class="flex-col flex py-3">
                    <label class="pb-2 text-gray-700 font-semibold">Name</label>
                    <input
                        type="text"
                        disabled
                        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
                        placeholder="John"
                    />
                </div>

                <div class="flex-col flex py-3 relative">
                    <label class="pb-2 text-gray-700 font-semibold">Slug</label>
                    <input
                        type="text"
                        value=""
                        disabled
                        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
                    />
                </div>

                <div class="flex py-3">
                    <button class="p-2 bg-blue-400 m-1 w-20 text-white rounded"
                        >Save</button
                    >
                </div>
            </div>
        </div>
    </div>
</Layout>
