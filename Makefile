GO_LDFLAGS := -s -w

.PHONY: go
go:
	@go run scripts/build-inject.go
	CC=arm-linux-gnueabihf-gcc CGO_ENABLED=1 GOARCH=arm GOOS=linux go build -ldflags="$(GO_LDFLAGS)" -o App/MiyooPod/MiyooPod src/*.go

.PHONY: install-hooks
install-hooks:
	@echo "Installing git hooks..."
	@chmod +x .git-hooks/pre-commit
	@mkdir -p $$(git rev-parse --git-dir)/hooks
	@cp .git-hooks/pre-commit $$(git rev-parse --git-dir)/hooks/pre-commit
	@chmod +x $$(git rev-parse --git-dir)/hooks/pre-commit
	@echo "✓ Pre-commit hook installed"

.PHONY: package
package:
	@echo "Creating MiyooPod release package..."
	@read -p "Enter version (e.g., 1.0.0): " VERSION; \
	read -p "Enter changelog: " CHANGELOG; \
	if [ -z "$$VERSION" ]; then \
		echo "Error: Version cannot be empty"; \
		exit 1; \
	fi; \
	echo "Updating APP_VERSION to $$VERSION..."; \
	go run scripts/update-version.go $$VERSION; \
	echo "Building app (stripped)..."; \
	go run scripts/build-inject.go; \
	rm -f App/MiyooPod/updater App/MiyooPod/updater_new; \
	CC=arm-linux-gnueabihf-gcc CGO_ENABLED=1 GOARCH=arm GOOS=linux go build -ldflags="$(GO_LDFLAGS)" -o App/MiyooPod/MiyooPod src/*.go; \
	echo "Creating release directory..."; \
	mkdir -p releases; \
	echo "Packaging release..."; \
	cd App && zip -r ../releases/MiyooPod-$$VERSION.zip MiyooPod -x "MiyooPod/updater" "MiyooPod/updater_new" "MiyooPod/*.log" && cd ..; \
	cp releases/MiyooPod-$$VERSION.zip releases/MiyooPod.zip; \
	echo "Updating version.json with checksum..."; \
	go run scripts/update-version.go $$VERSION releases/MiyooPod.zip $$CHANGELOG; \
	echo "✓ Release package created: releases/MiyooPod-$$VERSION.zip"; \
	echo "✓ Latest version copied to: releases/MiyooPod.zip"; \
	echo "✓ version.json updated with checksum and changelog"
