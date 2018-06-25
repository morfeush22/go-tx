defmodule Scheduler do
  @moduledoc false
  require Logger
  use Application

  defp verify_args([nil, nil | _]) do
    Logger.error("QUEUE_HOST and QUEUE_PORT empty")
    {:error}
  end

  defp verify_args([nil, _queue_port | _]) do
    Logger.error("QUEUE_HOST empty")
    {:error}
  end

  defp verify_args([_queue_host, nil | _]) do
    Logger.error("QUEUE_PORT empty")
    {:error}
  end

  defp verify_args([queue_host, queue_port | _]) do
    %{queue_host: queue_host, queue_port: queue_port}
  end

  def start(_type, _args) do
    [
      System.get_env("QUEUE_HOST") || Application.get_env(:scheduler, :queue_host),
      System.get_env("QUEUE_PORT") || Application.get_env(:scheduler, :queue_port)
    ]
    |> verify_args
    |> start
  end

  defp start ({:error}) do
    {:error, nil}
  end

  defp start (verified_args) do
    Scheduler.Supervisor.start_link(verified_args)
  end
end
