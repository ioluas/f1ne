.DE: run

run: build
	./f1ne

build: bundle dep
	go build -v -x .

bundle: bundle.sh
	./bundle.sh

dep:
	go mod tidy

vet:
	go vet .
	revive -config revive.toml -formatter friendly ./...

clean:
	rm -f f1ne ui/bundled.go

all: bundle
	fyne package --appID ioluas/f1ne --release -os darwin --icon res/icons/Icon.png
	fyne package --appID ioluas/f1ne --release -os linux --icon res/icons/Icon.png
	fyne package --appID ioluas/f1ne --release -os windows --icon res/icons/Icon.png
