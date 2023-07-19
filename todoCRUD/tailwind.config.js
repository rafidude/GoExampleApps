/** @type {import('tailwindcss').Config} */
module.exports = {
    content: {
        relative: true,
        files: ['./views/**/*.html', './public/**/*.html'],
    },
    theme: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
    ],
}
