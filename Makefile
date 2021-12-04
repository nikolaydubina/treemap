all: build docs

clean:
	-rm treemap
	-rm *.svg
	-rm coverage.txt

build: clean
	go build ./cmd/treemap

cover:
	go test -cover ./...

docs: 
	-rm docs/*.svg
	cat testdata/gapminder-2007-population-life.csv | ./treemap > docs/gapminder-2007-population-life.svg

.PHONY: all clean build cover docs
