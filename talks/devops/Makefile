NAME=devops
SRC=$(NAME).tex
OUT=$(NAME).pdf
NOTESNAME=$(NAME).notes
NOTESSRC=$(NOTESNAME).tex
NOTES=$(NOTESNAME).pdf

IMGS=$(shell find . -name '*.dot')

all: build notes

build: $(OUT)

$(OUT): $(SRC) images
	latexmk -pdfxe -pdfxelatex="xelatex -interaction=nonstopmode" $<

notes: $(NOTES)

$(NOTES): $(NOTESSRC) images
	latexmk -pdfxe -pdfxelatex="xelatex -interaction=nonstopmode" $<

.PHONY: clean init images clean-images

clean:
	latexmk -CA

images: $(IMGS:.dot=_gen.png)

$(IMGS:.dot=_gen.png): %_gen.png: %.dot
	dot -Tpng $^ > $@

clean-images:
	rm *_gen.png
