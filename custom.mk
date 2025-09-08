# Add custom targets below...

#
# compile is the default target; it builds the 
# application for the default platform (linux/amd64)
#
.DEFAULT_GOAL := compile

.PHONY: compile 
compile: goreleaser-dev ## build for the default linux/amd64 platform

.PHONY: snapshot 
snapshot: goreleaser-snapshot ## build a snapshot version for the supported platforms

.PHONY: release 
release: goreleaser-release ## build a release version (requires a valid tag)

.PHONY: clean 
clean: #clean the binary directory 
	@rm -rf dist

.PHONY: local-install
local-install: #install the plugin locally
	@mkdir -p ~/.opentofu.d/plugins/dihedron.org/edu/example/1.0.0/linux_amd64