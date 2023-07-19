Download atilwindcss standalone CLI from
https://github.com/tailwindlabs/tailwindcss/releases/tag/v3.3.3

Move the executable to /usr/local/bin
mv ~/desktop/tailwind\* /usr/local/bin/tailwindcss
cd /usr/local/bin
chmod +x tailwindcss

tailwindcss init

tailwindcss -c tailwind.config.js > ./public/styles.css
