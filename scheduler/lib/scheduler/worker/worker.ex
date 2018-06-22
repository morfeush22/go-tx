defmodule Scheduler.Worker.Worker do
  @moduledoc false

  defmacro __using__([queue_name: queue_name]) do
    quote location: :keep do
      import Task
      require Logger
      use GenServer

      def start_link(opts) do
        GenServer.start_link(__MODULE__, opts, name: __MODULE__)
      end

      def handle_call(message, _from, %{channel: channel} = state) do
        callback_queue = setup_callback_queue(channel)
        consumer = async(fn -> wait_for_messages(channel, on: callback_queue) end)
        trigger_calculation(message, channel, reply_to: callback_queue)
        json = await(consumer) |> handle_payload
        {:reply, json, state}
      end

      def init(%{queue_host: host, queue_port: port} = opts) do
        Process.flag(:trap_exit, true)
        with {:ok, conn} <- AMQP.Connection.open("amqp://guest:guest@#{host}:#{port}"),
             {:ok, channel} <- AMQP.Channel.open(conn) do
          {:ok, %{channel: channel, conn: conn}}
        else
          {:error, :econnrefused} -> Logger.error("Can not connect to AMQP server")
                                     :timer.sleep(1000)
                                     init(opts)
        end
      end

      def terminate(reason, %{conn: conn} = state) do
        AMQP.Connection.close(conn)
      end

      defp setup_callback_queue(channel) do
        with {:ok, %{queue: callback_queue}} <- AMQP.Queue.declare(channel, "", exclusive: true) do
          callback_queue
        end
      end

      defp wait_for_messages(channel, [on: callback_queue]) do
        AMQP.Basic.consume(channel, callback_queue, nil, no_ack: true)
        receive do
          {:basic_deliver, payload, _meta} -> payload
        end
      end

      defp trigger_calculation(message, channel, [reply_to: callback_queue]) do
        AMQP.Basic.publish(
          channel,
          "",
          unquote(queue_name),
          message,
          reply_to: callback_queue
        )
      end
    end
  end
end
