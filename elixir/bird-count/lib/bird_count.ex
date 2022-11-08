defmodule BirdCount do
  def today([]), do: nil

  def today([today_count | _]), do: today_count

  def increment_day_count([]), do: [1]

  def increment_day_count([today_count | past_list]), do: [today_count + 1 | past_list]

  def has_day_without_birds?([]), do: false

  def has_day_without_birds?([0 | _]), do: true

  def has_day_without_birds?([_ | past_list]), do: has_day_without_birds?(past_list)

  def total([]), do: 0

  def total([today_count | past_list]), do: today_count + total(past_list)

  def busy_days([]), do: 0

  def busy_days([today_count | past_list]) when today_count > 4, do: 1 + busy_days(past_list)

  def busy_days([_ | past_list]), do: busy_days(past_list)
end
