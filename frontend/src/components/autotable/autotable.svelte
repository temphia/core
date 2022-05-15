<script>
  export let actions = [];
  export let key_names = [];
  export let datas = [];
  export let action_key = "";
  export let color = [];

  const hashCode = (str) => {
    let hash = 77;
    for (var i = 0; i < str.length; i++) {
      hash = str.charCodeAt(i) + ((hash << 5) - hash);
    }
    return hash;
  };

  const color_it = (key, str) => {
    console.log(key, str);
    if (!color.includes(key)) {
      return "";
    }
    return `background: hsl(${hashCode(str) % 360}, 100%, 80%)`;
  };
</script>

<div class="overflow-x-auto p-2">
  <table class="table-auto border-collapse w-full bg-white shadow rounded-xl">
    <thead>
      <tr
        class="rounded-lg text-sm font-medium text-gray-700 text-left"
        style="font-size: 0.9674rem"
      >
        {#each key_names as key_name}
          <th class="px-2 py-1" style="background-color:#f8f8f8"
            >{key_name[1]}</th
          >
        {/each}
        <th class="px-2 py-2" style="background-color:#f8f8f8"> Actions </th>
      </tr>
    </thead>
    <tbody class="text-sm font-normal text-gray-700">
      {#each datas as data}
        <tr
          class="hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700"
        >
          {#each key_names as key_name}
            <td class="px-3 py-1">
              <span
                class="p-1 rounded-lg"
                style={color_it(key_name[0], data[key_name[0]] || "")}
                >{data[key_name[0]] || ""}</span
              >
            </td>
          {/each}

          <td class="px-3 py-1">
            {#each actions as action}
              <button
                on:click={() => action.Action(data[action_key])}
                class="p-1 m-1 text-sm font-semibold text-white rounded transform hover:scale-110 {action.Class ||
                  'bg-blue-400'}">{action.Name}</button
              >
            {/each}
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
