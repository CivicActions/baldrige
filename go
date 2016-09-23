#!/usr/bin/env bash
command -v wget >/dev/null 2>&1 || { echo >&2 "I require wget but it's not installed.  Aborting."; exit 1; }
command -v pdfseparate >/dev/null 2>&1 || { echo >&2 "I require pdfseparate but it's not installed.  Aborting."; exit 1; }
command -v tet >/dev/null 2>&1 || { echo >&2 "I require tet but it's not installed.  Aborting."; exit 1; }
rm README.md
wget -N https://www.nist.gov/sites/default/files/documents/2016/09/15/baldrige-cybersecurity-excellence-builder-draft-09.2016.pdf
wget -N https://www.pdflib.com/fileadmin/pdflib/TET-Cookbook/tetml_and_xslt/code/tetml2html.xsl
pdfseparate baldrige-cybersecurity-excellence-builder-draft-09.2016.pdf 'baldrige-page-%d.pdf'
for file in baldrige-page-*.pdf; do tet --tetml wordplus --outfile $file.tetml $file; xsltproc tetml2html.xsl $file.tetml > $file.html; pandoc -s -r html -t markdown $file.html -o $file.md; done
for i in $(seq 1 35); do cat "baldrige-page-$i.pdf.md" >> README.md; done
rm -rf baldrige-page-*
