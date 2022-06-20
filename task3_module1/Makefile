build:	## Builds the website using gohugo
	hugo -d dist/

clean:	## Cleans the dist/ directory (deletes whole folder)
	rm -r dist/

post:	## Creates a new post with parameters POST_TITLE and POST_NAME
	echo > content/posts/$(POST_NAME).md
	echo "title: $(POST_TITLE)" >> content/posts/$(POST_NAME).md
	echo "---" >> content/posts/$(POST_NAME).md

help:	## Show this help
	$(info 'build: Builds the website using gohugo')
	$(info 'clean: Cleans the dist/ directory (deletes whole folder)')
	$(info 'post: Creates a new post with parameters POST_TITLE and POST_NAME')
	$(info 'help: Show this help')
