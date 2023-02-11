defmodule TakeANumber do
  def start() do
    spawn(&listener/0)
  end

  defp listener(state \\ 0) do
    receive do
      {:report_state, sender_pid} ->
        send(sender_pid, state)

        listener(state)
      {:take_a_number, sender_pid} ->
        send(sender_pid, state + 1)

        listener(state + 1)
      :stop -> :stop
      _ -> listener(state)
    end
  end
end
