import Ts from "rollup-plugin-typescript2";

export default {
  input: [ "entries/test/index.ts"],
  output: {
    file: 'build/bundle.js',
  },
  format: "iife",
  plugins: [Ts()],
};
