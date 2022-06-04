<script>
  import DraggableCard from "./draggable_card.svelte";
  import Card from "./elements/card/card.svelte";
  import Gallery from "./elements/gallary/gallery.svelte";
  import Textbox from "./elements/textbox/textbox.svelte";
  import Todo from "./elements/todo/todo.svelte";

  const blocks = [
    {
      name: "The Catalyzer",
      type: "card",
      data: {
        subtext: "CATEGORY-1",
        contents: `Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.`,
      },
    },
    {
      name: "my list",
      type: "todo",
      data: [],
    },
    {
      name: "Report",
      type: "richtext",
      data: [],
    },
    {
      name: "Header",
      type: "simpletext",
      data: "This is one",
    },

    {
      name: "gallary1",
      type: "gallery",
      data: [],
    },
    {
      name: "My Group 1",
      type: "group",
      data: [
        {
          name: "my list 2",
          type: "todo",
          data: [],
        },
        {
          name: "Header2",
          type: "simpletext",
          data: "This is two",
        },
      ],
    },
  ];

  const links = [
    {
      from: "my list",
      to: "Report",
      name: "refers",
    },
    {
      to: "The Catalyzer",
      from: "Header",
      name: "link3",
      weighted: false,
    },
    {
      to: "The Catalyzer",
      from: "Header",
      name: "link3",
      weighted: true,
    },
  ];

  $: __block_pos = {};

  $: console.log("POS", __block_pos);

  const calculateLink = (fromElem, toElem) => {
    try {
      const fromCenter = [
        fromElem.top + fromElem.height / 2,
        fromElem.left + fromElem.width / 2,
      ];
      const toElemCenter = [
        toElem.top + toElem.height / 2,
        toElem.left + toElem.width / 2,
      ];

      const totalWeight =
        fromElem.height + fromElem.width + toElem.height + toElem.width;

      const fromWeight = fromElem.height + fromElem.width;
      const toWeight = toElem.height + toElem.width;

      const distance = Math.hypot(
        fromCenter[0] - toElemCenter[0],
        fromCenter[1] - toElemCenter[1]
      );

      const angle = Math.atan2(
        fromCenter[0] - toElemCenter[0],
        fromCenter[1] - toElemCenter[1]
      );

      const final = [
        toElemCenter[0],
        toElemCenter[1],
        distance,
        (angle * 180) / Math.PI,

        (fromCenter[0] + toElemCenter[0]) / 2,
        (fromCenter[1] + toElemCenter[1]) / 2,

        // wighted center
        (fromCenter[0] * toWeight + toElemCenter[0] * fromWeight) / totalWeight,
        (fromCenter[1] * toWeight + toElemCenter[1] * fromWeight) / totalWeight,
      ];

      console.log("@final =>", final);

      return final;
    } catch (error) {
      return undefined;
    }
  };
</script>

<div class="w-full h-full relative bg-blue-100">
  {#each blocks as block}
    <DraggableCard
      name={block.name}
      on:card_pos={(ev) => {
        __block_pos[ev.detail["name"]] = ev.detail;
        __block_pos = { ...__block_pos };
      }}
    >
      <svelte:fragment>
        {#if block.type === "card"}
          <Card {block} />
        {:else if block.type === "todo"}
          <Todo />
        {:else if block.type === "simpletext"}
          <div class="bg-blue-200 p-2">{block.name}</div>
          <div class="bg-gray-100 p-2">{block.data}</div>
        {:else if block.type === "richtext"}
          <Textbox />
        {:else if block.type === "gallery"}
          <Gallery />
        {:else if block.type === "group"}
          {#each block.data || [] as inner}
            <div class="p-1 border rounded">
              {#if inner.type === "card"}
                <Card {block} />
              {:else if inner.type === "todo"}
                <Todo />
              {:else if inner.type === "simpletext"}
                <div class="bg-blue-200 p-2">{inner.name}</div>
                <div class="bg-gray-100 p-2">{inner.data}</div>
              {:else if inner.type === "richtext"}
                <Textbox />
              {:else if block.type === "gallery"}
                <Gallery />
              {/if}
            </div>
          {/each}
        {/if}
      </svelte:fragment>
    </DraggableCard>
  {/each}

  {#each links as link}
    {@const pos = calculateLink(__block_pos[link.to], __block_pos[link.from])}

    {#if pos}
      <div
        class="h-1 z-10 absolute bg-gray-500 hover:bg-gray-700"
        style="top: {pos[0]}px; left: {pos[1]}px; width:{pos[2]}px; transform-origin: 0 0;  rotate: {pos[3]}deg;"
      />

      {#if link["weighted"]}
        <div
          class="z-10 mr-5 p-1 absolute bg-blue-600 hover:bg-blue-700 text-white rounded"
          style="top: {pos[6]}px; left: {pos[7]}px;"
        >
          {link.name}
        </div>
      {:else}
        <div
          class="z-10 mr-5 p-1 absolute bg-gray-600 hover:bg-gray-700 text-white rounded"
          style="top: {pos[4]}px; left: {pos[5]}px;"
        >
          {link.name}
        </div>
      {/if}
    {/if}
  {/each}
</div>
