/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      'templates/**/*.templ',
    ],
    darkMode: 'class',
    theme: {
      extend: {
        fontFamily: {
          mono: ['Courier Prime', 'monospace'],
        }
      },
    },
    plugins: [
      require('@tailwindcss/forms'),
    ],
    corePlugins: {
      preflight: true,
    }
  }