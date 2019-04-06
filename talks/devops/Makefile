NAME=devops
SRC=$(NAME).tex
OUT=$(NAME).pdf

IMGS=$(shell find . -name '*.dot')

$(OUT): $(SRC) images
	latexmk -pdf -bibtex -pdflatex="pdflatex -interaction=nonstopmode" $<

.PHONY: clean init images clean-images

clean:
	latexmk -bibtex -CA

images: $(IMGS:.dot=_gen.png)

$(IMGS:.dot=_gen.png): %_gen.png: %.dot
	dot -Tpng $^ > $@

clean-images:
	rm *_gen.png
