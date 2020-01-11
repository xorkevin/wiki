NAME=agiletools
SRC=$(NAME).tex
OUT=$(NAME).pdf
NOTESNAME=$(NAME).notes
NOTESSRC=$(NOTESNAME).tex
NOTES=$(NOTESNAME).pdf

.PHONY: all build notes

all: build notes

build: $(OUT)

$(OUT): $(SRC)
	latexmk -pdfxe -pdfxelatex="xelatex -interaction=nonstopmode" $<

notes: $(NOTES)

$(NOTES): $(NOTESSRC) $(SRC)
	latexmk -pdfxe -pdfxelatex="xelatex -interaction=nonstopmode" $<

.PHONY: clean

clean:
	latexmk -CA
