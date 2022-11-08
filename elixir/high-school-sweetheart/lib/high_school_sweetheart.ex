defmodule HighSchoolSweetheart do
  def first_letter(name) do
    name
    |> String.trim()
    |> String.first()
  end

  def initial(name) do
    first_letter_str =
      name
      |> first_letter()
      |> String.capitalize()

    first_letter_str <> "."
  end

  def initials(full_name) do
    [first_name, last_name | _] =
      full_name
      |> String.trim()
      |> String.split(" ")

    initial(first_name) <> " " <> initial(last_name)
  end

  def pair(full_name1, full_name2) do
    couple_initials = initials(full_name1) <> "  +  " <> initials(full_name2)

    heart = """
         ******       ******
       **      **   **      **
     **         ** **         **
    **            *            **
    **                         **
    **     #{couple_initials}     **
     **                       **
       **                   **
         **               **
           **           **
             **       **
               **   **
                 ***
                  *
    """

    heart
  end
end
