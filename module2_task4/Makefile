build:	## Builds the website using gohugo
	hugo -d dist/

clean:	## Cleans the dist/ directory (deletes whole folder)
	rm -rf dist/

post:	## Creates a new post with parameters POST_TITLE and POST_NAME
	hugo new posts/$(POST_NAME).md
	sed -i -E 's/(title: ")(.+)/\1$(POST_TITLE)"/' ./content/posts/$(POST_NAME).md 
	#cat ./content/posts/$(POST_NAME).md | sed 's/$(POST_NAME)/$(POST_TITLE)/' > ./content/posts/$(POST_NAME).md
	#echo > content/posts/$(POST_NAME).md
	#echo "title: $(POST_TITLE)" >> content/posts/$(POST_NAME).md
	#echo "---" >> content/posts/$(POST_NAME).md

check:
	docker run -v ${PWD}:/workdir ghcr.io/igorshubovych/markdownlint-cli:latest "content/**/*.md"
	for f in $$(find -wholename "*posts/*.md"); do \
		fpath=$$(realpath $$(dirname $$f)); echo $$fpath; \
		fname=$$(basename $$f); echo $$fname; \
		docker run -v $$fpath:/tmp:ro --rm -i ghcr.io/tcort/markdown-link-check:stable /tmp/$$fname; \
		done
		
validate:
	git clone https://github.com/holbertonschool/W3C-Validator.git testtools/
	-./testtools/w3c_validator.py ./dist/index.html
	rm -rf testtools

help:	## Show this help
	$(info 'build: Builds the website using gohugo')
	$(info 'clean: Cleans the dist/ directory (deletes whole folder)')
	$(info 'post: Creates a new post with parameters POST_TITLE and POST_NAME')
	$(info 'help: Show this help')
