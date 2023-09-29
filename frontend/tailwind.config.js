/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
  ],
  theme: {
    extend: {
      colors: {
        gBlack: "#161616",
        gGray: "#2e2e2e",
        gGraySoft: "#18181b",
        retroPurple: "#451952",
        retroPurplePink: "#662549",
        retroYellow: "#F39F5A",
        retroCream: "#DAC0A3",
        retroLightCream: "#EADBC8",
      },
      fontFamily: {
        Poppins: "Poppins, sans-serif",
      },
    },
  },
  plugins: [],
};
