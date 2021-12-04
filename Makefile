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
	cat testdata/gapminder-2007-population-life.csv | ./treemap -w 1080 -h 1080 > docs/gapminder-2007-population-life-1080x1080.svg
	cat testdata/gapminder-2007-population-life.csv | ./treemap -w 1080 -h 360 > docs/gapminder-2007-population-life-1080x360.svg
	cat testdata/gapminder-2007-population-life.csv | ./treemap -color none > docs/gapminder-2007-population-life-nocolor.svg
	cat testdata/gapminder-2007-population-life.csv | ./treemap -color RedBlu > docs/gapminder-2007-population-life-RedBlu.svg

.PHONY: all clean build cover docs
