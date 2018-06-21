defmodule Scheduler.Server.Server do
  @moduledoc false
  use Supervisor

  def start_link(opts) do
    Supervisor.start_link(__MODULE__, opts, name: __MODULE__)
  end

  def init(_opts) do
    children = [
      Plug.Adapters.Cowboy.child_spec(scheme: :http, plug: Scheduler.Server.Router, options: [port: 8080])
    ]
    Supervisor.init(children, strategy: :one_for_one)
  end
end
