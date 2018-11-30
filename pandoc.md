---
title: "Title here"
date: "Nov 29, 2018"
documentclass: "article"
papersize: "letter"
fontsize: "10pt"
classoption: "twocolumn"
#geometry: "margin=2cm"
listings: true
header-includes:
  - \usepackage{xcolor}
  - \lstset{
      basicstyle=\scriptsize,
      breaklines=true,
      keywordstyle=\color[HTML]{8959a8},
      commentstyle=\color[HTML]{8e908c},
      stringstyle=\color[HTML]{718c00},
      numberstyle=\color[HTML]{f5871f},
    }
bibliography: "bibliography.bib"
link-citations: true
csl: "/home/kevin/wiki/acm.csl"
abstract: |
  Abstract here
---

# Running

As seen in Appendix [@lst:pandocmake]

# Appendix

```{#lst:pandocmake .makefile caption="Pandoc Makefile"}
MD:=$(shell find . -name '*.md')
MDARGS=--listings --filter pandoc-crossref --filter pandoc-citeproc

$(MD:.md=.pdf): %.pdf: %.md
	pandoc $(MDARGS) $^ -o $@
```
