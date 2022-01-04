test:
	go test ./commands/
remove_tag:
	git push --delete origin $(GIT_TAG_NAME)
	git tag -d $(GIT_TAG_NAME)