set dotenv-load

dev:
  air --build.cmd "templ generate && go build ." --build.bin "./leet_tutor" --build.exclude_regex ".*_templ.go" --build.include_ext "go,templ"
