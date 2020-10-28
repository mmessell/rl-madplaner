package latex

import (
	"strings"

	"../io"
	"../log"
	"../models"
	. "../types"
)

var courseHeading = "part*"
var recipeHeading = "section*"
var shoppingListHeading = "section*"
var ingredientHeading = "subsection*"

func CreateLatexFile(mp models.Madplan, latexDir string, recipeImgDir string) {
	log.Start("latexfile", "CreateLatexFile")

	for _, week := range mp.Weeks {
		latex := getLatex(mp, week, recipeImgDir)
		io.CreateFile(getLatexDir(latexDir, mp.Filename), getTexFilename(mp.Filename, week), latex)
	}

	log.End("latexfile", "CreateLatexFile")
}

func getLatex(mp models.Madplan, week Double, recipeImgDir string) string {
	latex := "\\documentclass[oneside]{book}\n"
	latex += "\\author{Anne Skaarup}\n"
	latex += "\\title{Ren Energi}\n"
	latex += "\\usepackage{../../packages/renlykke}\n"
	latex += "\\usepackage{color, colortbl}\n"
	latex += "\\usepackage{graphicx}\n"
	latex += "\\usepackage{float}\n"
	latex += "\\usepackage[margin=3cm]{geometry}\n"

	latex += "\\pagenumbering{gobble}\n"

	latex += "\\begin{document}\n\n"
	latex += getContent(mp, week, recipeImgDir) + "\n"
	latex += "\\end{document}\n"

	return latex
}

func getContent(mp models.Madplan, week Double, recipeImgDir string) string {
	content := "\\" + courseHeading + "{" + mp.Name + "}\n"
	content += getTableContent(mp, week)
	content += getRecipesContent("Morgenmadsopskrifter", mp.GetUniqueBreakfastRecipes(), recipeImgDir)
	content += getRecipesContent("Frokostopskrifter", mp.GetUniqueLunchRecipes(), recipeImgDir)
	content += getRecipesContent("Aftensmadsopskrifter", mp.GetUniqueDinnerRecipes(), recipeImgDir)
	content += getShoppingListContent(mp)
	return content
}

func getTableContent(mp models.Madplan, week Double) string {
	content := "\\begin{table}[] \n"
	content += "\\begin{tabular}{|m{2cm}|m{3cm}|m{3cm}|m{3cm}|m{2cm}|} \n"
	content += "\\hline \n"
	content += "\\rowcolor{tblHeader} \n"

	content += "\\vspace{3mm} \\cellcolor{tblColumn} \\bold{Uge " + week.ToString() + "} \\vspace{3mm}& \\bold{Morgenmad} & \\bold{Middagsmad} & \\bold{Aftensmad} & \\bold{Kalorier} \\\\ \\hline \n"

	for _, dp := range mp.Dagsplaner {
		cBreakfast := dp.Breakfast.GetCaloriesPerPortions()
		cLunch := dp.Lunch.GetCaloriesPerPortions()
		cDinner := dp.Dinner.GetCaloriesPerPortions()
		cTotal := cBreakfast + cLunch + cDinner

		content += "\\vspace{1mm} \\cellcolor{tblColumn}\\bold{" + dp.Day + "}\\vspace{1mm}&"
		content += "\\vspace{1mm}" + dp.Breakfast.Name + " (" + cBreakfast.ToString() + "kcal) \\vspace{1mm}&"
		content += "\\vspace{1mm}" + dp.Lunch.Name + " (" + cLunch.ToString() + "kcal) \\vspace{1mm}&"
		content += "\\vspace{1mm}" + dp.Dinner.Name + " (" + cDinner.ToString() + "kcal) \\vspace{1mm}&"
		content += "\\vspace{1mm}" + cTotal.ToString() + "kcal \\vspace{1mm} \\\\ \\hline \n"
	}

	content += "\\end{tabular} \n"
	content += "\\end{table} \n"

	return content
}

func getRecipesContent(course string, recipes []models.Recipe, recipeImgDir string) string {
	content := "\\" + courseHeading + "{" + course + "}\n"

	for _, recipe := range recipes {
		content += "\\" + recipeHeading + "{" + recipe.Name + "}\n"

		content += "\\begin{figure}[H]"
		content += "\\includegraphics[width=\\linewidth]{" + recipeImgDir + "/" + recipe.Image + "}"
		content += "\\end{figure}"

		content += "Portion(er): " + recipe.Portions.ToString() + "\n\n"
		content += "Kalorier: " + recipe.Calories.ToString() + "\n\n"

		if len(recipe.Ingredients) > 0 {
			content += "\\" + ingredientHeading + "{Ingredienser}\n"
			content += "\\begin{itemize}\n"

			for _, ing := range recipe.Ingredients {
				content += "\\item " + ing.Amount.ToString() + " " + ing.Unit + " " + ing.Name + "\n"
			}

			content += "\\end{itemize}\n"
		}

		if len(recipe.Procedure) > 0 {
			content += "\\" + ingredientHeading + "{Fremgangsmåde}\n"
			for _, pro := range recipe.Procedure {
				content += pro + "\n\n"
			}
		}

		content += "\\newpage\n"
	}

	return content
}

func getShoppingListContent(mp models.Madplan) string {
	content := "\\" + shoppingListHeading + "{Det skal du bruge}\n"

	content += "\\begin{itemize}\n"
	for _, ing := range mp.GetShoppingList() {
		if ing.ShoppingList {
			content += "\\item " + ing.Amount.RoundUp().ToString() + " " + ing.Unit + " " + ing.Name + "\n"
		}
	}
	content += "\\end{itemize}\n"

	return content
}

func getLatexDir(latexDir string, filename string) string {
	return latexDir + "/" + strings.Replace(filename, ".json", "", -1)
}

func getTexFilename(fn string, week Double) string {
	return strings.Replace(fn, ".json", "-"+week.ToString()+".tex", -1)
}
