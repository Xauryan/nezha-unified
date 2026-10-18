# The pinned Certimate/PocketBase baseline requires Go 1.26 JSON semantics.
export GOTOOLCHAIN := go1.26.6

.PHONY: prepare frontend themes test test-core test-web race build check
prepare:
	bash script/prepare-unified.sh

frontend:
	cd web/admin && npm ci --no-audit --no-fund && npm run build
	python3 script/copy-admin.py

themes:
	python3 script/fetch-public-themes.py

test-core: prepare
	go test -mod=readonly -p 2 ./service/unified/... github.com/certimate-go/certimate/pkg/embedded

test: prepare
	go test -mod=readonly -p 2 ./... github.com/certimate-go/certimate/pkg/embedded

test-web:
	cd web/admin && npm ci --no-audit --no-fund && npm test && npm run lint

race: prepare
	go test -mod=readonly -p 2 -race ./service/unified/... ./cmd/dashboard/controller github.com/certimate-go/certimate/pkg/embedded

build: prepare frontend themes
	mkdir -p bin
	go build -mod=readonly -trimpath -ldflags '-X github.com/nezhahq/nezha/service/singleton.Version=unified-0.1.0-dev' -o bin/nezha-unified ./cmd/dashboard

check:
	python3 script/check-integration.py
