REGISTRY=gcr.io/keep-up-199521
TAG ?= ${TRAVIS_TAG}
IMAGE="$(REGISTRY)/keep-up-graphql:$(TAG)"

build:
	docker build -t $(IMAGE) .

publish:
	gcloud docker -- push $(IMAGE)
