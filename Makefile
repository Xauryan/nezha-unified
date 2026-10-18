# The pinned Certimate/PocketBase baseline requires Go 1.26 JSON semantics.
export GOTOOLCHAIN := go1.26.6
CERTIMATE_TEST_PACKAGES := github.com/certimate-go/certimate/pkg/embedded github.com/certimate-go/certimate/internal/embeddedhooks github.com/certimate-go/certimate/internal/embedcrypto github.com/certimate-go/certimate/internal/certacme github.com/certimate-go/certimate/internal/workflow/engine

.PHONY: prepare frontend themes test test-core test-web test-e2e test-agent race build build-agent check vet
prepare:
	bash script/prepare-unified.sh

frontend:
	cd web/admin && npm ci --no-audit --no-fund && npm run build
	python3 script/copy-admin.py

themes:
	python3 script/fetch-public-themes.py

test-core: prepare
	go test -mod=readonly -p 2 ./service/unified/... ./service/unifiedbackup ./service/agentcall ./pkg/certdeploy github.com/certimate-go/certimate/pkg/embedded

test: prepare
	go test -tags=certimate_full -mod=readonly -p 2 ./... $(CERTIMATE_TEST_PACKAGES)

test-web:
	cd web/admin && npm ci --no-audit --no-fund && npm test && npm run lint

race: prepare
	go test -tags=certimate_full -mod=readonly -p 2 -race ./service/unified/... ./service/unifiedbackup ./service/agentcall ./pkg/certdeploy ./cmd/dashboard/controller $(CERTIMATE_TEST_PACKAGES)

vet: prepare
	go vet -tags=certimate_full -mod=readonly -p 2 ./... $(CERTIMATE_TEST_PACKAGES)

test-agent:
	cd third_party/agent && go test -mod=readonly -p 2 ./...
	cd third_party/agent && go test -mod=readonly -p 2 -race ./pkg/agentproto ./pkg/certdeploy ./cmd/agent -run 'TestUnified|TestAgent|TestDeploy'

build-agent:
	mkdir -p bin
	cd third_party/agent && go build -mod=readonly -trimpath -o ../../bin/nezha-agent ./cmd/agent

test-e2e:
	cd web/admin && npx playwright install --with-deps chromium && npm run e2e

build: prepare frontend themes
	mkdir -p bin
	go build -tags=certimate_full -mod=readonly -p 2 -trimpath -ldflags '-X github.com/nezhahq/nezha/service/singleton.Version=unified-0.2.0-dev' -o bin/nezha-unified ./cmd/dashboard

check:
	python3 script/check-integration.py
