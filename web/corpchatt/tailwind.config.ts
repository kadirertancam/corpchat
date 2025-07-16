import type { Config } from "tailwindcss";
export default {
  content: ["./src/**/*.{js,ts,jsx,tsx,mdx}"],
  theme: {
    extend: {
      colors: {
        "corp-primary": "#0052D4",
        "corp-secondary": "#7FBDFF",
        "corp-accent": "#FFAB40",
      },
      fontFamily: {
        sans: ["Inter var", "system-ui", "sans-serif"],
        heading: ["Poppins", "sans-serif"],
      },
    },
  },
  plugins: [require("@tailwindcss/typography")],
} satisfies Config;