defmodule Secrets do
  def secret_add(secret) do
    &(&1 + secret)
  end

  def secret_subtract(secret) do
    &(&1 - secret)
  end

  def secret_multiply(secret) do
    multiplier_func = fn param ->
      param * secret
    end

    multiplier_func
  end

  def secret_divide(secret) do
    &div(&1, secret)
  end

  def secret_and(secret) do
    &Bitwise.band(&1, secret)
  end

  def secret_xor(secret) do
    &Bitwise.bxor(&1, secret)
  end

  def secret_combine(secret_function1, secret_function2) do
    combined_func = fn param ->
      secret_function2.(secret_function1.(param))
    end

    combined_func
  end
end
