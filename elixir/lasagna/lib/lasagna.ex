defmodule Lasagna do
  def expected_minutes_in_oven, do: 40

  def remaining_minutes_in_oven(oven_min) do
    expected_minutes_in_oven() - oven_min
  end

  def preparation_time_in_minutes(layer_count), do: 2 * layer_count

  def total_time_in_minutes(layer_count, oven_time) do
    oven_time + preparation_time_in_minutes(layer_count)
  end

  def alarm(), do: "Ding!"
end
