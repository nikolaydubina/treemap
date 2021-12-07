all: build docs

clean:
	-rm treemap
	-rm *.svg
	-rm coverage.txt
	-rm docs/*.svg

build: clean
	go build ./cmd/treemap

cover:
	go test -cover ./...

docs: 
	cat testdata/gapminder-2007-population-life.csv | ./treemap > docs/gapminder-2007-population-life.svg
	cat testdata/gapminder-2007-population-life.csv | ./treemap -w 1080 -h 1080 > docs/gapminder-2007-population-life-1080x1080.svg
	cat testdata/gapminder-2007-population-life.csv | ./treemap -w 1080 -h 360 > docs/gapminder-2007-population-life-1080x360.svg
	cat testdata/gapminder-2007-population-life.csv | ./treemap -color none > docs/gapminder-2007-population-life-nocolor.svg
	cat testdata/gapminder-2007-population-life.csv | ./treemap -color RedBlu > docs/gapminder-2007-population-life-RedBlu.svg
	cat testdata/gapminder-2007-population-life.csv | ./treemap -impute-heat > docs/gapminder-2007-population-life-impute-heat.svg
	cat testdata/gapminder-2007-population-life.csv | ./treemap -color balanced > docs/gapminder-2007-population-life-balanced.svg
	#cat testdata/find-src-go-dir.csv | ./treemap -h 4096 -w 4096 > docs/find-src-go-dir.svg
	cat testdata/gapminder-2007-population-life.csv | ./treemap -color RdYlGn > docs/gapminder-2007-population-RdYlGn.svg

.PHONY: all clean build cover docs
