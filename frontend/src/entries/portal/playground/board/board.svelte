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
      name: "my list",
      type: "richtext",
      data: [],
    },
    {
      name: "Header",
      type: "simpletext",
      data: "This is one",
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
</script>

<div class="w-full h-full relative bg-blue-100">
  {#each blocks as block}
    <DraggableCard>
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
              {/if}
            </div>
          {/each}
        {/if}
      </svelte:fragment>
    </DraggableCard>
  {/each}
</div>
