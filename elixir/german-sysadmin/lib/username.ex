defmodule Username do
  def sanitize(''), do: ''

  def sanitize([head | tail]) do
    sanitized_character =
      case head do
        ?_ -> '_'
        ?ä -> 'ae'
        ?ö -> 'oe'
        ?ü -> 'ue'
        ?ß -> 'ss'
        char when char >= ?a and char <= ?z -> [char]
        _ -> ''
      end

    sanitized_character ++ sanitize(tail)
  end
end
