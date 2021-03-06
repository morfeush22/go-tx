defmodule Scheduler.Server.Router do
  @moduledoc false
  import Plug.Conn
  use Plug.Router
  use Plug.Debugger

  plug(:match)
  plug(:dispatch)

  @data "data"

  get "/tx" do
    conn = conn |> fetch_query_params
    case conn.params do
      %{@data => data} ->
        response = GenServer.call(Scheduler.Worker.Server, {:handle_message, data})
        send_resp(conn, 200, response)
      _ -> send_resp(conn, 404, "")
    end
  end

  match _ do
    send_resp(conn, 404, "")
  end
end
