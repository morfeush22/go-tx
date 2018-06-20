defmodule Scheduler.Worker.CRCCalc do
  @moduledoc false
  alias AMQP.{Basic, Channel, Connection, Queue}
  import Task
  use GenServer

  @queue "crc-calc"

  def start_link(opts) do
    GenServer.start_link(__MODULE__, opts, name: __MODULE__)
  end

  def handle_call({:calc_crc, message}, _from, %{channel: channel} = state) do
    callback_queue = setup_callback_queue(channel)
    consumer = async(fn -> wait_for_messages(channel, on: callback_queue) end)
    trigger_calculation(message, channel, reply_to: callback_queue)
    json = await(consumer) |> handle_payload
    {:reply, json, state}
  end

  def init(%{queue_host: host, queue_port: port}) do
    with {:ok, conn} <- Connection.open("amqp://guest:guest@#{host}:#{port}"),
         {:ok, channel} <- Channel.open(conn) do
      {:ok, %{channel: channel}}
    end
  end

  defp setup_callback_queue(channel) do
    with {:ok, %{queue: callback_queue}} <- Queue.declare(channel, "", exclusive: true) do
      callback_queue
    end
  end

  defp wait_for_messages(channel, [on: callback_queue]) do
    Basic.consume(channel, callback_queue, nil, no_ack: true)
    receive do
      {:basic_deliver, payload, _meta} -> payload
    end
  end

  defp trigger_calculation(message, channel, [reply_to: callback_queue]) do
    Basic.publish(
      channel,
      "",
      @queue,
      message,
      reply_to: callback_queue
    )
  end

  defp handle_payload(calculation) do
    calculation
  end
end
