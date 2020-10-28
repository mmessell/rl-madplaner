#!/usr/bin/env bash

echo "START RUN.COMMAND"

BASE_DIR=$PWD
CONFIG_DIR="$BASE_DIR/configuration"
MADPLAN_DIR="$BASE_DIR/madplaner-go"
LATEX_DIR="$BASE_DIR/output/latexfiles"
PDF_DIR="$BASE_DIR/output/pdf"

####
echo "START DELETE OUTPUT"
rm  -r output
echo "END DELETE OUTPUT"

echo "START BUILD.COMMAND"
cd "$MADPLAN_DIR"
./build.command
echo "END BUILD.COMMAND"

echo "START PROGRAM"
cd "$MADPLAN_DIR"
./program
echo "END PROGRAM"

echo "START LATEX"
cd "$LATEX_DIR"
for d in */ ; do
    cd "$d"
    for f in *.tex ; do
        pdflatex "$f"
        rm *.aux
        rm *.log
    done
    cd ..
done

cp -r ./* ../pdf
cd "$PDF_DIR"
rm -r */*.tex

echo "END LATEX"

echo "END RUN.COMMAND"