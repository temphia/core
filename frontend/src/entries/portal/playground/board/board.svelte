<script>
  import DraggableCard from "./draggable_card.svelte";

  var pos1 = 0,
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

<div class="w-full h-full relative bg-blue-100">
  <DraggableCard>
    <svelte:fragment>
      <div class="bg-blue-200 p-2">Header</div>
      <div class="bg-gray-100 p-2">This is one</div>
    </svelte:fragment>
  </DraggableCard>

  <DraggableCard>
    <svelte:fragment>


        <div class="title font-bold text-lg">Todo Application</div>
        <div class="flex items-center text-sm mt-2">
          <button >
            <svg class="w-3 h-3 mr-3 focus:outline-none" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" stroke="currentColor">
              <path d="M12 4v16m8-8H4"></path>
            </svg>
          </button>
          <span>Click to add task</span>
        </div>
        <input type="text" placeholder="what is your plan for today" class=" rounded-sm shadow-sm px-4 py-2 border border-gray-200 w-full mt-4" x-model="todoTitle" x-ref="addTask" >
  
        <!-- todo list -->
        <ul class="todo-list mt-4">
  
          <template x-for="todo in todos" :key="todo.id">
            <li class="flex justify-between items-center mt-3" >
              <div class="flex items-center" >
                <input type="checkbox" name="" id="" x-model="todo.isComplete">
                <div class="capitalize ml-3 text-sm font-semibold" x-text="todo.title"></div>
              </div>
              <div>
                <button>
                  <svg class=" w-4 h-4 text-gray-600 fill-current" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" stroke="currentColor">
                    <path d="M6 18L18 6M6 6l12 12"></path>
                  </svg>
                </button>
              </div>
            </li>
          </template>
  
        </ul>

    </svelte:fragment>
  </DraggableCard>

  <DraggableCard>
    <svelte:fragment>
      <div class="h-full w-96 rounded-xl shadow-cla-violate bg-gradient-to-r from-pink-50 to-red-50 overflow-hidden">
        <img class="lg:h-48 md:h-36 w-full object-cover object-center scale-110 transition-all duration-400 hover:scale-100" src="https://images.unsplash.com/photo-1624628639856-100bf817fd35?ixlib=rb-1.2.1&amp;ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8M2QlMjBpbWFnZXxlbnwwfHwwfHw%3D&amp;auto=format&amp;fit=crop&amp;w=600&amp;q=60" alt="blog">
        <div class="p-6">
          <h2 class="tracking-widest text-xs title-font font-medium text-gray-400 mb-1">CATEGORY-1</h2>
          <h1 class="title-font text-lg font-medium text-gray-600 mb-3">The Catalyzer</h1>
          <p class="leading-relaxed mb-3">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p>
          <div class="flex items-center flex-wrap ">
            <button class="bg-gradient-to-r from-orange-300 to-amber-400 hover:scale-105 drop-shadow-md shadow-cla-violate px-4 py-1 rounded-lg">Learn more</button>
            
          </div>
        </div>
      </div>
    </svelte:fragment>
  </DraggableCard>


  <DraggableCard>

    <div class="w-96">


    <div class="mb-2 pb-2">
      <h3 class="font-semibold text-lg text-gray-600">Bar Chart</h3>
      <p class="text-sm text-gray-500">Profit Share between customers</p>
    </div>
    <div id="chartdiv" class="w-full" style="height: 240px">
      <div style="width: 100%; height: 100%; position: relative;">
        <svg version="1.1"  role="group" 
          style="width: 100%; height: 100%; overflow: visible;"><defs><clipPath id="id-46"><rect width="568" height="240"></rect></clipPath><linearGradient id="gradient-id-69" x1="1%" x2="99%" y1="59%" y2="41%"><stop stop-color="#474758" offset="0"></stop><stop stop-color="#474758" stop-opacity="1" offset="0.75"></stop><stop stop-color="#3cabff" stop-opacity="1" offset="0.755"></stop></linearGradient><filter id="filter-id-78" filterUnits="objectBoundingBox" width="200%" height="200%" x="-50%" y="-50%"></filter><filter id="filter-id-99" filterUnits="objectBoundingBox" width="200%" height="200%" x="-50%" y="-50%"></filter><clipPath id="id-132"><path d="M0,0 L477,0 L477,167 L0,167 L0,0"></path></clipPath><clipPath id="id-404"><rect width="477" height="167"></rect></clipPath><filter id="filter-id-51" filterUnits="objectBoundingBox" width="200%" height="200%" x="-50%" y="-50%"><feGaussianBlur result="blurOut" in="SourceGraphic" stdDeviation="1.5"></feGaussianBlur><feOffset result="offsetBlur" dx="1" dy="1"></feOffset><feFlood flood-color="#000000" flood-opacity="0.5"></feFlood><feComposite in2="offsetBlur" operator="in"></feComposite><feMerge><feMergeNode></feMergeNode><feMergeNode in="SourceGraphic"></feMergeNode></feMerge></filter><filter id="filter-id-66" filterUnits="objectBoundingBox" width="120%" height="120%" x="-10%" y="-10%"><feColorMatrix type="saturate" values="0"></feColorMatrix></filter><filter id="filter-id-137" filterUnits="objectBoundingBox" width="200%" height="200%" x="-50%" y="-50%"><feGaussianBlur result="blurOut" in="SourceGraphic" stdDeviation="1.5"></feGaussianBlur><feOffset result="offsetBlur" dx="1" dy="1"></feOffset><feFlood flood-color="#000000" flood-opacity="0.5"></feFlood><feComposite in2="offsetBlur" operator="in"></feComposite><feMerge><feMergeNode></feMergeNode><feMergeNode in="SourceGraphic"></feMergeNode></feMerge></filter></defs><g><g fill="#ffffff" fill-opacity="0"><rect width="568" height="240"></rect></g><g><g role="region" clip-path="url(&quot;#id-46&quot;)" opacity="1" aria-label="Chart"><g transform="translate(15,15)"><g><g><g><g><g><g><g transform="translate(61,0)"><g fill="#ffffff" fill-opacity="0"><rect width="477" height="167"></rect></g><g><g><g><g stroke="#000000" stroke-opacity="0.15" fill="none" transform="translate(0,223)" display="none"><path d=" M0,0  L477,0 " transform="translate(-0.5,-0.5)"></path></g><g stroke="#000000" stroke-opacity="0.15" fill="none" transform="translate(0,195)" display="none"><path d=" M0,0  L477,0 " transform="translate(-0.5,-0.5)"></path></g><g stroke="#000000" stroke-opacity="0.15" fill="none" transform="translate(0,167)"><path d=" M0,0  L477,0 " transform="translate(-0.5,-0.5)"></path></g><g stroke="#000000" stroke-opacity="0.15" fill="none" transform="translate(0,139)" display="none"><path d=" M0,0  L477,0 " transform="translate(-0.5,-0.5)"></path></g><g stroke="#000000" stroke-opacity="0.15" fill="none" transform="translate(0,111)"><path d=" M0,0  L477,0 " transform="translate(-0.5,-0.5)"></path></g><g stroke="#000000" stroke-opacity="0.15" fill="none" transform="translate(0,-56)" display="none"><path d=" M0,0  L477,0 " transform="translate(-0.5,-0.5)"></path></g><g stroke="#000000" stroke-opacity="0.15" fill="none" transform="translate(0,56)"><path d=" M0,0  L477,0 " transform="translate(-0.5,-0.5)"></path></g><g stroke="#000000" stroke-opacity="0.15" fill="none"><path d=" M0,0  L477,0 " transform="translate(-0.5,-0.5)"></path></g><g stroke="#000000" stroke-opacity="0.15" fill="none" transform="translate(0,278)" display="none"><path d=" M0,0  L477,0 " transform="translate(-0.5,-0.5)"></path></g></g></g><g><g><g stroke="#000000" stroke-opacity="0.15" fill="none"><path d=" M0,0  L0,167 " transform="translate(-0.5,-0.5)"></path></g><g stroke="#000000" stroke-opacity="0.15" fill="none" transform="translate(159,0)"><path d=" M0,0  L0,167 " transform="translate(-0.5,-0.5)"></path></g><g stroke="#000000" stroke-opacity="0.15" fill="none" transform="translate(318,0)"><path d=" M0,0  L0,167 " transform="translate(-0.5,-0.5)"></path></g><g stroke="#000000" stroke-opacity="0.15" fill="none" transform="translate(477,0)"><path d=" M0,0  L0,167 " transform="translate(-0.5,-0.5)"></path></g></g></g><g><g><g role="group" stroke-opacity="0" fill-opacity="1" stroke="#67b7dc" aria-label="Visits"><g><g clip-path="url(&quot;#id-132&quot;)"><g><g><g><g stroke-opacity="1" fill-opacity="0.8" stroke-width="2" fill="#67b7dc" stroke="#67b7dc" role="menuitem" focusable="true" tabindex="0" transform="translate(15.898,41.75)" id="id-406"><g><g><path d="M0,0 L127.202,0 a0,0 0 0 1 0,0 L127.202,125.25 a0,0 0 0 1 -0,0 L0,125.25 a0,0 0 0 1 -0,-0 L0,0 a0,0 0 0 1 0,-0 Z"></path></g></g></g><g stroke-opacity="1" fill-opacity="0.8" stroke-width="2" fill="#67b7dc" stroke="#67b7dc" role="menuitem" focusable="true" tabindex="0" transform="translate(174.902,121.353)"><g><g><path d="M0,0 L127.197,0 a0,0 0 0 1 0,0 L127.197,45.647 a0,0 0 0 1 -0,0 L0,45.647 a0,0 0 0 1 -0,-0 L0,0 a0,0 0 0 1 0,-0 Z"></path></g></g></g><g stroke-opacity="1" fill-opacity="0.8" stroke-width="2" fill="#67b7dc" stroke="#67b7dc" role="menuitem" focusable="true" tabindex="0" transform="translate(333.9,161.99)"><g><g><path d="M0,0 L127.202,0 a0,0 0 0 1 0,0 L127.202,5.01 a0,0 0 0 1 -0,0 L0,5.01 a0,0 0 0 1 -0,-0 L0,0 a0,0 0 0 1 0,-0 Z"></path></g></g></g></g></g></g></g><g></g></g></g></g></g><g clip-path="url(&quot;#id-404&quot;)"><g><g fill="#67b7dc" stroke="#67b7dc"><g></g></g></g></g><g><g><g><g></g></g><g><g></g></g></g></g><g><g></g></g><g><g></g></g><g role="button" focusable="true" tabindex="0" opacity="0" visibility="hidden" aria-hidden="true" transform="translate(437,-3)" aria-labelledby="id-37-title"><g fill="#6794dc" stroke="#ffffff" fill-opacity="1" stroke-opacity="0" transform="translate(0,8)"><path d="M17,0 L18,0 a17,17 0 0 1 17,17 L35,17 a17,17 0 0 1 -17,17 L17,34 a17,17 0 0 1 -17,-17 L0,17 a17,17 0 0 1 17,-17 Z"></path></g><g transform="translate(9,9)"><g stroke="#ffffff" style="pointer-events: none;" transform="translate(0,8)"><path d=" M0,0  L11,0 " transform="translate(2.5,7.5)"></path></g><g fill="#000000" style="pointer-events: none;" transform="translate(17,8)"><g display="none"></g></g></g><title id="id-37-title">Zoom Out</title></g></g></g><g><g><g aria-hidden="true"><g><g fill="#000000" transform="translate(0,83.5) rotate(-90)"><g display="none"></g></g><g stroke="#000000" stroke-opacity="0.15" fill="none" opacity="0" visibility="hidden" aria-hidden="true"><path transform="translate(-0.5,-0.5)" d=" M0,0  L568,0 "></path></g><g><g><g fill="#000000" aria-label="L" fill-opacity="0" opacity="0" stroke-opacity="0" style="pointer-events: none;" transform="translate(61,83.5)"><g transform="translate(-51,-11.5)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>1,800</tspan></text></g></g><g fill="#000000" transform="translate(61,222.667)" display="none"><g transform="translate(-51,-11.5)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>1,700</tspan></text></g></g><g fill="#000000" transform="translate(61,194.833)" display="none"><g transform="translate(-51,-11.5)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>1,750</tspan></text></g></g><g fill="#000000" transform="translate(61,167)"><g transform="translate(-51,-11.5)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>1,800</tspan></text></g></g><g fill="#000000" transform="translate(61,139.167)" display="none"><g transform="translate(-51,-11.5)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>1,850</tspan></text></g></g><g fill="#000000" transform="translate(61,111.333)"><g transform="translate(-51,-11.5)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>1,900</tspan></text></g></g><g fill="#000000" transform="translate(61,-55.667)" display="none"><g transform="translate(-51,-11.5)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>2,200</tspan></text></g></g><g fill="#000000" transform="translate(61,55.667)"><g transform="translate(-51,-11.5)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>2,000</tspan></text></g></g><g fill="#000000" transform="translate(61,0)"><g transform="translate(-51,-11.5)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>2,100</tspan></text></g></g><g fill="#000000" transform="translate(61,278.333)" display="none"><g transform="translate(-51,-11.5)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>1,600</tspan></text></g></g></g></g><g stroke="#000000" stroke-opacity="0" fill="none" style="pointer-events: none;" transform="translate(61,0)"><path d=" M0,0  L0,167 " transform="translate(-0.5,-0.5)"></path></g></g></g></g></g><g transform="translate(538,0)"><g></g></g></g></g><g><g transform="translate(61,0)"></g></g><g transform="translate(0,167)"><g transform="translate(61,0)"><g aria-hidden="true"><g><g stroke="#000000" stroke-opacity="0" fill="none" style="pointer-events: none;"><path d=" M0,0  L477,0 " transform="translate(-0.5,-0.5)"></path></g><g><g><g fill="#000000" aria-label="L" fill-opacity="0" opacity="0" stroke-opacity="0" style="pointer-events: none;" transform="translate(238.5,0)"><g transform="translate(-4.5,10)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>L</tspan></text></g></g><g fill="#000000" transform="translate(79.499,0)"><g transform="translate(-15.5,10)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>USA</tspan></text></g></g><g fill="#000000" transform="translate(238.5,25)"><g transform="translate(-21.5,10)" style="user-select: none;">
            <text x="0" y="23" dy="-6.21"><tspan>China</tspan></text></g></g><g fill="#000000" transform="translate(397.501,0)"><g transform="translate(-21.5,10)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>Japan</tspan></text></g></g><g fill="#000000" display="none" transform="translate(556.499,25)"><g transform="translate(0,10)" display="none"></g></g></g></g><g stroke="#000000" stroke-opacity="0.15" fill="none" display="none" transform="translate(477,-167)" opacity="0" visibility="hidden" aria-hidden="true"><path transform="translate(-0.5,-0.5)" d=" M0,0  L0,167 "></path></g><g fill="#000000" transform="translate(238.5,43)"><g display="none"></g></g></g></g></g></g></g></g></g></g></g></g><g><g><g role="tooltip" visibility="hidden" opacity="0"><g fill="#ffffff" style="pointer-events: none;" fill-opacity="0.9" stroke-width="1" stroke-opacity="1" stroke="#ffffff" filter="url(&quot;#filter-id-51&quot;)" transform="translate(0,6)"><path d="M3,0 L3,0 L0,-6 L13,0 L21,0 a3,3 0 0 1 3,3 L24,8 a3,3 0 0 1 -3,3 L3,11 a3,3 0 0 1 -3,-3 L0,3 a3,3 0 0 1 3,-3"></path></g><g><g fill="#ffffff" style="pointer-events: none;" transform="translate(12,6)"><g transform="translate(0,7)" display="none"></g></g></g></g><g visibility="hidden" style="pointer-events: none;" display="none"><g fill="#ffffff" opacity="1"><rect width="568" height="240"></rect></g><g><g transform="translate(284,120)"><g><g stroke-opacity="1" fill="#f3f3f3" fill-opacity="0.8"><g><g><path d=" M53,0  a53,53,0,0,1,-106,0 a53,53,0,0,1,106,0 M42,0  a42,42,0,0,0,-84,0 a42,42,0,0,0,84,0 L42,0 "></path></g></g></g><g stroke-opacity="1" fill="#000000" fill-opacity="0.2"><g><g><path d=" M50,0  a50,50,0,0,1,-100,0 a50,50,0,0,1,100,0 M45,0  a45,45,0,0,0,-90,0 a45,45,0,0,0,90,0 L45,0 "></path></g></g></g><g fill="#000000" fill-opacity="0.4"><g style="user-select: none;" transform="translate(-20.5,-11.5)"><text x="20.5px" y="23" dy="-6.21" text-anchor="middle"><tspan>100%</tspan></text></g></g></g></g></g></g><g opacity="0.3" style="cursor: pointer;" aria-labelledby="id-66-title" filter="url(&quot;#filter-id-66&quot;)" transform="translate(0,219)"><g fill="#ffffff" opacity="0"><rect width="66" height="21"></rect></g><g><g shape-rendering="auto" fill="none" stroke-opacity="1" stroke-width="1.7999999999999998" stroke="#3cabff"><path d=" M15,15  C17.4001,15 22.7998,15.0001 27,15 C31.2002,14.9999 33.2999,6 36,6 C38.7001,6 38.6999,10.5 40.5,10.5 C42.3001,10.5 42.2999,6 45,6 C47.7001,6 50.9999,14.9999 54,15 C57.0002,15.0001 58.7999,15 60,15"></path></g><g shape-rendering="auto" fill="none" stroke-opacity="1" stroke-width="1.7999999999999998" stroke="url(&quot;#gradient-id-69&quot;)"><path d=" M6,15  C8.2501,15 9.7498,15.0001 15,15 C20.2502,14.9999 20.7748,3.6 27,3.6 C33.2252,3.6 33.8998,14.9999 39.9,15 C45.9002,15.0001 45.9748,15 51,15 C56.0252,15 57.7499,15 60,15"></path></g></g><title id="id-66-title">Chart created using amCharts library</title></g><g role="tooltip" visibility="hidden" opacity="0"><g fill="#000000" style="pointer-events: none;" fill-opacity="1" stroke-width="1" stroke-opacity="1" stroke="#000000" transform="translate(76,182)"><path d="M0,0 L20,0 a0,0 0 0 1 0,0 L20,10 a0,0 0 0 1 -0,0 L0,10 a0,0 0 0 1 -0,-0 L0,0 a0,0 0 0 1 0,-0"></path></g><g><g fill="#ffffff" style="pointer-events: none;" transform="translate(86,182)"><g transform="translate(0,5)" display="none"></g></g></g></g><g role="tooltip" visibility="hidden" opacity="0"><g fill="#000000" style="pointer-events: none;" fill-opacity="1" stroke-width="1" stroke-opacity="1" stroke="#000000" transform="translate(-25,15)"><path d="M0,0 L20,0 a0,0 0 0 1 0,0 L20,0 L20,0 L25,0 L20,10 L20,10 a0,0 0 0 1 -0,0 L0,10 a0,0 0 0 1 -0,-0 L0,0 a0,0 0 0 1 0,-0"></path></g><g><g fill="#ffffff" style="pointer-events: none;" transform="translate(-15,15)"><g transform="translate(0,5)" display="none"></g></g></g></g><g role="tooltip" opacity="0" aria-describedby="id-406" transform="translate(155.499,119.375)" aria-hidden="true" visibility="hidden"><g fill="#67b7dc" style="pointer-events: none;" fill-opacity="0.9" stroke-width="1" stroke-opacity="1" stroke="#ffffff" filter="url(&quot;#filter-id-137&quot;)" transform="translate(6,-17)"><path d="M3,0 L102,0 a3,3 0 0 1 3,3 L105,31 a3,3 0 0 1 -3,3 L3,34 a3,3 0 0 1 -3,-3 L0,31 L0,22 L-6,17 L0,12 L0,3 a3,3 0 0 1 3,-3"></path></g><g><g fill="#000000" style="pointer-events: none;" transform="translate(58.5,-17)"><g transform="translate(-40.5,7)" style="user-select: none;"><text x="0" y="23" dy="-6.21"><tspan>USA: </tspan><tspan style="font-weight:bold">2,025</tspan></text></g></g></g></g></g></g></g></g></svg></div></div>

          </div>

  </DraggableCard>


</div>
