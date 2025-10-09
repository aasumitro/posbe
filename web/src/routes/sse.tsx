import {createFileRoute} from "@tanstack/react-router";
import {Button} from "@/components/ui/button";
import {type FormEvent, useReducer} from "react";
import type {HTTPResponse} from "@/types/http-response";
import {api, API_URL, catchHTTPError} from "@/lib/api";
import {useMutation} from "@tanstack/react-query";
import {useEventSource, useEventSourceListener} from "@/hooks/use-sse";
import {ScrollArea} from "@/components/ui/scroll-area";

export const Route = createFileRoute("/sse")({
  component: SSETestPage
})

function useFireEvent() {
  const fire = async (
    body: string,
  ): Promise<HTTPResponse<string>> => {
    try {
      const response =
        await api.post<HTTPResponse<string>>(
          `/orders/events`, body);
      return response.data;
    } catch (error: unknown) {
      return catchHTTPError(error);
    }
  };

  return useMutation({ mutationFn: fire })
}

type Message = { id: number }

function messageReducer(state: Message[], action: Message): Message[] {
  return [...state, action]
}

function SSETestPage() {
  const {mutate: fire} = useFireEvent();
  const [messages, addMessages] = useReducer(messageReducer, []);

  const fireEvent = (e: FormEvent, type: string) => {
    e.preventDefault();
    fire(JSON.stringify({"type": type}), {
      onSuccess: (resp) => console.log(resp),
      onError: (err) => console.log(err),
    })
  }

  const [eventSource] = useEventSource(`${API_URL}/orders/events`, true);

  useEventSourceListener(eventSource, ["update"], (evt) =>
    addMessages(JSON.parse(evt.data)), [addMessages]);

  return (
    <section className="relative flex flex-col items-center justify-center w-screen h-screen">
      <div className="flex gap-2">
        <Button
          className="cursor-pointer"
          onClick={(e) => fireEvent(e, "T1")}
        >Fire Table 1</Button>

        <Button
          className="cursor-pointer"
          onClick={(e) => fireEvent(e, "T3")}
        >Fire Table 3</Button>

        <Button
          className="cursor-pointer"
          onClick={(e) => fireEvent(e, "RELOAD")}
        >Fire Reload</Button>
      </div>

      <ScrollArea className="h-96 w-96 mt-4 p-4 rounded-md border">
        {[...messages].reverse().map((msg, idx) => (
          <div key={idx}>{msg.id}</div>
        ))}
      </ScrollArea>
    </section>
  )
}