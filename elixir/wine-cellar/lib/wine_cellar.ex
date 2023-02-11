defmodule WineCellar do
  def explain_colors do
    [
      white: "Fermented without skin contact.",
      red: "Fermented with skin contact using dark-colored grapes.",
      rose: "Fermented with some skin contact, but not enough to qualify as a red wine."
    ]
  end

  def filter(cellar, color), do: Keyword.get_values(cellar, color)

  def filter(cellar, color, year: year) do
    filter_by_year(filter(cellar, color), year)
  end

  def filter(cellar, color, country: country) do
    filter_by_country(filter(cellar, color), country)
  end

  def filter(cellar, color, year: year, country: country) do
    filter_by_country(filter(cellar, color, year: year), country)
  end

  def filter(cellar, color, country: country, year: year) do
    filter_by_country(filter(cellar, color, year: year), country)
  end

  # The functions below do not need to be modified.

  defp filter_by_year(wines, year)
  defp filter_by_year([], _year), do: []

  defp filter_by_year([{_, year, _} = wine | tail], year) do
    [wine | filter_by_year(tail, year)]
  end

  defp filter_by_year([{_, _, _} | tail], year) do
    filter_by_year(tail, year)
  end

  defp filter_by_country(wines, country)
  defp filter_by_country([], _country), do: []

  defp filter_by_country([{_, _, country} = wine | tail], country) do
    [wine | filter_by_country(tail, country)]
  end

  defp filter_by_country([{_, _, _} | tail], country) do
    filter_by_country(tail, country)
  end
end
