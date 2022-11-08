defmodule FreelancerRates do
  def daily_rate(hourly_rate), do: 8.0 * hourly_rate

  def apply_discount(before_discount, discount),
    do: before_discount - discount / 100.0 * before_discount

  def monthly_rate(hourly_rate, discount) do
    round(Float.ceil(22 * apply_discount(daily_rate(hourly_rate), discount)))
  end

  def days_in_budget(budget, hourly_rate, discount) do
    daily_discounted_charges = apply_discount(daily_rate(hourly_rate), discount)

    Float.floor(budget / daily_discounted_charges, 1)
  end
end
