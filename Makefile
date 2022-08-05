gen-query-gorm:
	rm -f dal/*/*
	cd cmd/gen && go run generator.go

.PHONY: gen-query-gorm