BAZEL_BIN:=bazel

build: update-gazelle
	$(BAZEL_BIN) build //...

update-gazelle: bazel-fmt
	$(BAZEL_BIN) run //:gazelle

bazel-fmt:
	$(BAZEL_BIN) run //:buildifier

update-go-mods:
	cd src/ && go mod tidy && go mod verify
	$(BAZEL_BIN) run //:gazelle
	$(BAZEL_BIN) run //:gazelle -- \
		update-repos \
		-from_file=src/go.mod \
		-to_macro=src/go.bzl%go_dependencies \
		-build_file_proto_mode=disable_global
	$(BAZEL_BIN) run //:gazelle
