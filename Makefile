.DE: run

run: build
	./f1ne

build: package dep
	go build -v -x .

package: res/icons
	./bundle.sh

dep:
	go mod tidy

vet:
	go vet .

clean:
	rm -f f1ne ui/bundled.go

all: package
	fyne package --appID ioluas/f1ne --release -os darwin --icon res/icons/Icon.png
	fyne package --appID ioluas/f1ne --release -os linux --icon res/icons/Icon.png
	fyne package --appID ioluas/f1ne --release -os windows --icon res/icons/Icon.png
