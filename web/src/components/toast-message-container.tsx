import {toast} from "sonner";
import {XIcon} from "lucide-react";

export const ToastMessageContainer = (
  title: any,
  body: any,
) => {
  toast.custom((t) => (
    <a href="#" className="flex gap-2 p-4 rounded-lg bg-gray-950 text-gray-300">
      <div className="w-full">
        <h1 className="text-sm font-bold">{title}</h1>
        <div className="text-gray-400 text-xs font-normal">{body}</div>
      </div>
      <XIcon className="w-6 h-6" onClick={() => toast.dismiss(t)}/>
    </a>
  ), {duration: 15000});
}