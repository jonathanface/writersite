/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['public/**/*.{html,js}'],
  theme: {
    extend: {
      fontFamily: {
        'gauchi': ['Gauchi+Hand', 'sans-serif'],
        'roboto': ['Roboto', 'sans-serif'] 
      },
    },
  },
  plugins: [],
}
