
buildyarn:
	yarn install
	yarn build
	cp build/index.html ../../server/main/templates/view_location.tmpl
	rm -rf ../../server/main/static/js/main.*.*
	rm -rf ../../server/main/static/css/main.*.*
	rsync -avrP build/static/ ../../server/main/static/
	git add ../../server/main/static/

dev:
	npm run start

production:
	npm run build

install:
	rm -rf node_modules
	npm install
