defmodule Scheduler.Supervisor do
  @moduledoc false
  use Supervisor

  def start_link(verified_args) do
    Supervisor.start_link(__MODULE__, verified_args, name: __MODULE__)
  end

  def init(verified_args) do
    children = [
      {Scheduler.Server.Server, verified_args},
      {Scheduler.Worker.Supervisor, verified_args}
    ]
    Supervisor.init(children, strategy: :one_for_one)
  end
end
