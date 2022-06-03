<div class="title font-bold text-lg">Todo Application</div>
<div class="flex items-center text-sm mt-2">
  <button>
    <svg
      class="w-3 h-3 mr-3 focus:outline-none"
      fill="none"
      stroke-linecap="round"
      stroke-linejoin="round"
      stroke-width="2"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path d="M12 4v16m8-8H4" />
    </svg>
  </button>
  <span>Click to add task</span>
</div>
<input
  type="text"
  placeholder="what is your plan for today"
  class=" rounded-sm shadow-sm px-4 py-2 border border-gray-200 w-full mt-4"
  x-model="todoTitle"
  x-ref="addTask"
/>

<!-- todo list -->
<ul class="todo-list mt-4">
  <template x-for="todo in todos" :key="todo.id">
    <li class="flex justify-between items-center mt-3">
      <div class="flex items-center">
        <input type="checkbox" name="" id="" x-model="todo.isComplete" />
        <div
          class="capitalize ml-3 text-sm font-semibold"
          x-text="todo.title"
        />
      </div>
      <div>
        <button>
          <svg
            class=" w-4 h-4 text-gray-600 fill-current"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
    </li>
  </template>
</ul>
