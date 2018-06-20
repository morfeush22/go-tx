defmodule Scheduler do
  @moduledoc false
  use Application

  defp verify_args([queue_host, queue_port | _]) do
    case [queue_host, queue_port] do
      [nil, _] -> {:error}
      [_, nil] -> {:error}
      [_, _] -> %{queue_host: queue_host, queue_port: queue_port}
    end
  end

  defp verify_args(_) do
    {:error}
  end

  def start(_type, _args) do
    [
      System.get_env("QUEUE_HOST"),
      System.get_env("QUEUE_PORT")
    ]
    |> verify_args
    |> start
  end

  defp start ({:error}) do
    {:error, nil}
  end

  defp start (verified_args) do
    children = [
      {Scheduler.Supervisor, verified_args}
    ]
    Supervisor.start_link(children, strategy: :one_for_one)
  end
end
