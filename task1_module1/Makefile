build:
	hugo -d dist/
clean:
	rm -r dist/

post:
	echo "---" > content/posts/$(POST_NAME).md
	echo "title: $(POST_TITLE)" >> content/posts/$(POST_NAME).md
	echo "---" >> content/posts/$(POST_NAME).md
#	hugo new $(POST_NAME).md
	

# clean:
# 	echo cleaning...
