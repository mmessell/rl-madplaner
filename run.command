#!/usr/bin/env bash

echo "START RUN.COMMAND"

BASE_DIR=$(dirname "$0")
MADPLAN_DIR="$BASE_DIR/madplaner-go"
LATEX_DIR="$BASE_DIR/configuration/latex/texfiles"
PDF_DIR="$BASE_DIR/pdf"

echo "START DELETE PDF"
cd "$PDF_DIR"
rm  -r */
echo "END DELETE PDF"

echo "START DELETE TEX"
cd "$LATEX_DIR"
rm  -r */
echo "END DELETE TEX"

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

cp -r ./* ../../../pdf
cd "$PDF_DIR"
rm -r */*.tex

echo "END LATEX"

echo "END RUN.COMMAND"