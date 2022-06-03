<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  let pos1 = 0,
    pos2 = 0,
    pos3 = 0,
    pos4 = 0;

  let elmnt;

  function dragMouseDown(e) {
    e = e || window.event;
    e.preventDefault();
    // get the mouse cursor position at startup:
    pos3 = e.clientX;
    pos4 = e.clientY;
    document.onmouseup = closeDragElement;
    // call a function whenever the cursor moves:
    document.onmousemove = elementDrag;
  }

  function elementDrag(e) {
    e = e || window.event;
    e.preventDefault();
    // calculate the new cursor position:
    pos1 = pos3 - e.clientX;
    pos2 = pos4 - e.clientY;
    pos3 = e.clientX;
    pos4 = e.clientY;
    // set the element's new position:
    elmnt.style.top = elmnt.offsetTop - pos2 + "px";
    elmnt.style.left = elmnt.offsetLeft - pos1 + "px";
  }

  function closeDragElement() {
    // stop moving when mouse button is released:
    document.onmouseup = null;
    document.onmousemove = null;
  }
</script>

<div
  class="absolute border bg-white shadow rounded"
  bind:this={elmnt}
  style="min-width: 5rem; min-height: 5rem;"
>
  <div
    class="h-2 cursor-pointer w-full bg-yellow-100 hover:bg-yellow-300"
    on:mousedown={dragMouseDown}
  />

  <div
    class="h-2 w-2 rounded-full absolute -right-1 top-1/2 bg-red-200 hover:bg-red-400"
  />

  <div class="p-2">
    <slot />
  </div>
</div>
