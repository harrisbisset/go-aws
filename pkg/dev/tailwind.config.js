/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './pkg/dev/render/views/**/*.templ',
    ],
    darkMode: "class",
    theme: {
        extend: {
            colors: {
                slate: "#1e2d3b",
                orange: "#E87A1A",
            },
            fontSize: {
                l: "1.2rem",
                lm: "1.3rem",
                ml: "1.4rem",
                xl: "1.7rem",
            },
        },
    },
    plugins: [],
    corePlugins: {
        preflight: true,
    },
};