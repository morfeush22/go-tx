defmodule Scheduler.Worker.Supervisor do
  @moduledoc false
  use Supervisor

  def start_link(opts) do
    Supervisor.start_link(__MODULE__, opts, name: __MODULE__)
  end

  def init(opts) do
    children = [
      {Scheduler.Worker.CRCCalc, opts},
      {Scheduler.Worker.Server, []},
    ]
    Supervisor.init(children, strategy: :one_for_one)
  end
end
