# buf generates go code
buf lint
buf format -w
buf generate

# poetry generates python library code
poetry run python -m grpc_tools.protoc -I . --python_betterproto_out=./../py/lib ./nanala/v1/nanala.proto
