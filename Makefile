install:
	go build -o /tmp/chains

run: install
	/tmp/chains

# TODO: Tirso : Add target to run the tests


swagger_validate:
	swagger validate ./swagger.yaml

generate:
	swagger generate server -A chains -f ./swagger.yaml

models:
	# -s the package to save the server specific code (default: restapi)
	# --exclude-main if not, then a cmd dir would be created
	swagger generate model --keep-spec-order -f swagger.yaml
