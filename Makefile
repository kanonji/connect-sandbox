.PHONY: buf_generate
buf_generate:
	buf generate --include-imports --template buf.gen-es.yaml
	buf generate --template buf.gen-go.yaml
