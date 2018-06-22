defmodule Scheduler.Worker.CRCCalc do
  @moduledoc false
  use Scheduler.Worker.Worker, queue_name: "crc-calc"

  defp handle_payload(calculation) do
    with {:ok, json} <- Poison.decode(calculation),
         {:ok, data} <- Base.decode64(json["data"]) do
      data
    end
  end
end
