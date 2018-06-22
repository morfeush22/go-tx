defmodule Scheduler.Worker.Server do
  @moduledoc false
  use GenServer

  def start_link(_opts) do
    GenServer.start_link(__MODULE__, nil, name: __MODULE__)
  end

  def handle_call({:handle_message, message}, _from, _state) do
    result = message
             |> calc_crc
             |> modulate
    {:reply, result, nil}
  end

  def calc_crc(message) do
    GenServer.call(Scheduler.Worker.CRCCalc, message)
  end

  def modulate(message) do
    GenServer.call(Scheduler.Worker.Modulator, message)
  end

  def init(_opts) do
    {:ok, nil}
  end
end
