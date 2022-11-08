defmodule KitchenCalculator do
  def get_volume(volume_pair), do: elem(volume_pair, 1)

  def to_milliliter({:cup, volume}), do: {:milliliter, 240.0 * volume}

  def to_milliliter({:fluid_ounce, volume}), do: {:milliliter, 30.0 * volume}

  def to_milliliter({:teaspoon, volume}), do: {:milliliter, 5.0 * volume}

  def to_milliliter({:tablespoon, volume}), do: {:milliliter, 15.0 * volume}

  def to_milliliter(volume_pair), do: volume_pair

  def from_milliliter({:milliliter, volume}, :milliliter), do: {:milliliter, volume}

  def from_milliliter({:milliliter, volume}, :cup), do: {:cup, volume / 240.0}

  def from_milliliter({:milliliter, volume}, :fluid_ounce), do: {:fluid_ounce, volume / 30.0}

  def from_milliliter({:milliliter, volume}, :teaspoon), do: {:teaspoon, volume / 5.0}

  def from_milliliter({:milliliter, volume}, :tablespoon), do: {:tablespoon, volume / 15.0}

  def convert(volume_pair, unit) do
    {_, milliliter_volume} = to_milliliter(volume_pair)

    {_, convert_to_ml_per_unit} = to_milliliter({unit, 1})

    {unit, milliliter_volume / convert_to_ml_per_unit}
  end
end
