<script>
  let schema = {
    name: "Kitchen Schema Example",
    slug: "ks2",
    description: "This is test Group 2",
    tables: [
      {
        name: "big table1",
        slug: "b1",
        description: "this is very big table",
        icon: "",
        main_column: "",
        columns: [
          {
            name: "words",
            slug: "w1",
            ctype: "shorttext",
            description: "it contain words",
            icon: "",
            options: [],
            not_nullable: true,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Short Text",
            slug: "sh1",
            ctype: "shorttext",
            description: "this is shorttext exmaple",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Phone Number",
            slug: "pn1",
            ctype: "phonenumber",
            description: "this is phonenumber example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Select",
            slug: "ss1",
            ctype: "select",
            description: "this is select example",
            icon: "",
            options: ["apple", "mango", "pineapple"],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Row Formula",
            slug: "rf1",
            ctype: "rowformula",
            description: "this is select example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "File",
            slug: "f1",
            ctype: "file",
            description: "this is file example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Multi File",
            slug: "mf1",
            ctype: "multifile",
            description: "this is multi file example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Checkbox",
            slug: "cb1",
            ctype: "checkbox",
            description: "this is checkbox example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Currency",
            slug: "cy1",
            ctype: "currency",
            description: "this is currency example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Number",
            slug: "n1",
            ctype: "number",
            description: "this is number example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Location",
            slug: "l1",
            ctype: "location",
            description: "this is location example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Datetime",
            slug: "dt1",
            ctype: "datetime",
            description: "this is datetime example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Multiselect",
            slug: "ms1",
            ctype: "multiselect",
            description: "this is multiselect example",
            icon: "",
            options: ["zebra", "cat", "dog", "mouse", "snake", "elephant"],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Longtext",
            slug: "long1",
            ctype: "longtext",
            description: "this is longtext example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Single User",
            slug: "siguser",
            ctype: "singleuser",
            description: "this is singleuser example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Multi User",
            slug: "muser",
            ctype: "multiuser",
            description: "this is multiuser example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Email",
            slug: "email",
            ctype: "email",
            description: "this is email example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Json",
            slug: "js1",
            ctype: "json",
            description: "this is json example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Range Number",
            slug: "rg1",
            ctype: "rangenumber",
            description: "this is Range Number example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Color",
            slug: "clr1",
            ctype: "color",
            description: "this is color example",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Hard Num Ref",
            slug: "hnref",
            ctype: "number",
            description: "this is  Hard reference number",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Soft Num Ref",
            slug: "snref",
            ctype: "number",
            description: "this is  Soft reference number",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Hard Text Ref",
            slug: "htref",
            ctype: "shorttext",
            description: "this is hard reference text",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Soft Text Ref",
            slug: "stref",
            ctype: "shorttext",
            description: "this is Soft reference text",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
        ],
        column_refs: [
          {
            slug: "ref1",
            type: "hard_pri",
            target: "ex2",
            from_cols: ["hnref"],
            to_cols: [],
          },
          {
            slug: "ref2",
            type: "soft_pri",
            target: "ex2",
            from_cols: ["snref"],
            to_cols: [],
          },
          {
            slug: "ref3",
            type: "soft_text",
            target: "ex2",
            from_cols: ["stref"],
            to_cols: ["title"],
          },
        ],
      },
      {
        name: "Example 2",
        slug: "ex2",
        description: "this is example 2",
        icon: "",
        main_column: "title",
        columns: [
          {
            name: "Title",
            slug: "title",
            ctype: "shorttext",
            description: "it contain words",
            icon: "",
            options: [],
            not_nullable: true,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Notes",
            slug: "notes",
            ctype: "longtext",
            description: "it contains paragraph",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Images",
            slug: "image",
            ctype: "multifile",
            description: "it contains images",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Entry date",
            slug: "date1",
            ctype: "datetime",
            description: "contains entry date",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },

          {
            name: "Published",
            slug: "pub1",
            ctype: "checkbox",
            description: "contains if its published",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
        ],
      },

      {
        name: "Example 3",
        slug: "ex3",
        description: "this is example 3",
        icon: "",
        main_column: "title",
        columns: [
          {
            name: "Title",
            slug: "title",
            ctype: "shorttext",
            description: "it contain words",
            icon: "",
            options: [],
            not_nullable: true,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Notes",
            slug: "notes",
            ctype: "longtext",
            description: "it contains paragraph",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Images",
            slug: "image",
            ctype: "multifile",
            description: "it contains images",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
          {
            name: "Published",
            slug: "pub1",
            ctype: "checkbox",
            description: "contains if its published",
            icon: "",
            options: [],
            not_nullable: false,
            pattern: "",
            strict_pattern: false,
          },
        ],
      },
    ],
  };
</script>

<div class="bg-blue-100 p-10">
  {#each schema.tables || [] as table}
    <div class="p-2 border bg-white shadow m-2">
      <div class="grid gap-6 mb-6 lg:grid-cols-2">
        <div>
          <label
            for="first_name"
            class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
            >Name</label
          >
          <input
            type="text"
            id="first_name"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            required
          />
        </div>
        <div class="flex justify-end content-end justify-items-end">
          <button
            class="p-1 h-8 bg-blue-400 rounded text-white hover:bg-blue-500 flex text-sm"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z"
                clip-rule="evenodd"
              />
            </svg>
            Add</button
          >
        </div>
      </div>

      <table class="table-auto border-collapse w-full border">
        <thead
          ><tr class="rounded-lg text-sm font-medium text-gray-700 text-left"
            ><th class="px-2 py-1" style="background-color: rgb(248, 248, 248);"
              >Name</th
            ><th class="px-2 py-1" style="background-color: rgb(248, 248, 248);"
              >Slug</th
            >

            <th class="px-2 py-1" style="background-color: rgb(248, 248, 248);"
              >ctype</th
            >

            <th
              class="px-2 py-1"
              style="background-color: rgb(248, 248, 248);"
            />
          </tr>
        </thead>

        <tbody class="text-sm font-normal text-gray-700">
          {#each table.columns as col}
            <tr
              class="hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700"
            >
              <td class="px-3 py-1"
                ><span class="p-1 rounded-lg" style="">{col.name}</span></td
              >
              <td class="px-3 py-1"
                ><span class="p-1 rounded-lg" style="">{col.slug}</span></td
              >
              <td class="px-3 py-1"
                ><span class="p-1 rounded-lg" style="">{col.ctype}</span></td
              >

              <td class="px-3 py-1"
                ><button
                  class="p-1 m-1 text-sm font-semibold text-white rounded transform hover:scale-110 bg-blue-400 inline-flex"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-5 w-5"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      d="M17.414 2.586a2 2 0 00-2.828 0L7 10.172V13h2.828l7.586-7.586a2 2 0 000-2.828z"
                    />
                    <path
                      fill-rule="evenodd"
                      d="M2 6a2 2 0 012-2h4a1 1 0 010 2H4v10h10v-4a1 1 0 112 0v4a2 2 0 01-2 2H4a2 2 0 01-2-2V6z"
                      clip-rule="evenodd"
                    />
                  </svg>
                  Edit</button
                >
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/each}
</div>
