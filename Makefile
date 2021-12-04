e2e:
	-rm treemap
	go build ./cmd/treemap
	cat testdata/gapminder-2007-population-life.csv | ./treemap > 1.svg
	open 1.svg

