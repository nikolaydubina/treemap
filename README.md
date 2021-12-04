Pretty Treemaps

```
$ go install github.com/nikolaydubina/treemap/cmd/treemap@latest 
$ echo '
Africa/Algeria,33333216,72
Africa/Angola,12420476,42
Africa/Benin,8078314,56
' | treemap > out.svg
```

Default options
```
$ cat testdata/gapminder-2007-population-life.csv | ./treemap > docs/gapminder-2007-population-life.svg
```
![example-1](./docs/gapminder-2007-population-life.svg)

Adjusting size
```
$ cat testdata/gapminder-2007-population-life.csv | ./treemap -w 1080 -h 640 > docs/gapminder-2007-population-life-1080x640.svg
```
![example-2](./docs/gapminder-2007-population-life-1080x640.svg)

## Format

Size is required. Heat is optional. For more customizations  

```
</ delimitered path>,<size>,<heat>
```

## Algorithms

* `Squarified` algorithm for treemap layout problem. This is very common algorithm used in Plotly and most of visualization packages. _"Squarified Treemaps", Mark Bruls, Kees Huizing, and Jarke J. van Wijk, 2000_

## Contributions

Welcomed!

## References

* Plotly treemaps: https://plotly.com/python/treemaps/
* go-colorful: https://github.com/lucasb-eyer/go-colorful
* D3 treemap is using Squerified: https://github.com/d3/d3-hierarchy
* Interactive treemap: https://github.com/vasturiano/treemap-chart
* Squerified in Rust: https://github.com/bacongobbler/treemap-rs
* Squerified in JavaScript: https://github.com/clementbat/treemap
* Squerified in Python: https://github.com/laserson/squarify
* Treemap Go tool: https://github.com/willpoint/treemap
