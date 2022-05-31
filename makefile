check: # check if binaries required are in $PATH
	@which wasm2wat
	@which cargo
	@which docker
	@which docker-compose
	@which go
	@which rustc

#############BACKEND############################
backend_docker_run:
	docker-compose up
backend_server_run:
	go run .

############# FRONTEND ############################
frontend_run:
	cd frontend && npm run dev
frontend_portal_build:
	cd frontend/src/lib/launcher && go run buildentry.go
	cd frontend && npm run build:portal
frontend_auth_build:
	cd frontend && npm run build:auth
frontend_operator_build:
	cd frontend && npm run build:operator
frontend_gen_launcher_template:
	cd frontend/src/lib/launcher && go run buildentry.go

############# EXECUTOR ############################
executor_sdk_rust_smoke_test:
	cd backend/sdk/rust_plug_wasm2_sdk && \
	cargo test -- --nocapture
executor_sdk_rust_test_wasm2:
	cd backend/sdk/rust_plug_wasm2_test && cargo build
	wasm2wat target/wasm32-unknown-unknown/debug/rust_plug_wasm2_test.wasm -o target/wasm32-unknown-unknown/debug/rust_plug_wasm2_test.wat
	go run backend/sdk/rust_plug_wasm2_test/*.go target/wasm32-unknown-unknown/debug/rust_plug_wasm2_test.wat
executor_sdk_goja_test:
	cd backend/sdk/goja_plug_sdk && rollup -c && go run entries/test/main.go


dev_pg:
	psql -p 7032 -h localhost -U temphia