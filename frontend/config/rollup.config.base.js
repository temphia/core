import svelte from "rollup-plugin-svelte";
import commonjs from "@rollup/plugin-commonjs";
import resolve from "@rollup/plugin-node-resolve";
import sveltePreprocess from "svelte-preprocess";
import typescript from "@rollup/plugin-typescript";
import css from "rollup-plugin-css-only";
import json from "rollup-plugin-json";

export const config = (entryFile, prod) => ({
  input: `src/entries/${entryFile}.ts`,
  output: {
    sourcemap: true,
    format: "iife",
    name: "app",
    file: `public/build/${entryFile}.js`,
  },
  plugins: [
    svelte({
      preprocess: sveltePreprocess({ postcss: true }),
      compilerOptions: {
        // enable run-time checks when not in production
        dev: !prod,
      },
    }),
    // we'll extract any component CSS out into
    // a separate file - better for performance
    css({ output: `${entryFile}.css` }),

    // If you have external dependencies installed from
    // npm, you'll most likely need these plugins. In
    // some cases you'll need additional configuration -
    // consult the documentation for details:
    // https://github.com/rollup/plugins/tree/master/packages/commonjs
    resolve({
      browser: true,
      dedupe: ["svelte"],
    }),
    commonjs(),
    json(),
    typescript({
      sourceMap: !prod,
      inlineSources: !prod,
    }),
  ],
});
