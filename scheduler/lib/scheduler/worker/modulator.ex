defmodule Scheduler.Worker.Modulator do
  @moduledoc false
  use Scheduler.Worker.Worker, queue_name: "modulator"

  defp handle_payload(calculation) do
    calculation
  end
end
