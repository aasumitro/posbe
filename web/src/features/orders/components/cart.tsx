import { Button } from "@/components/ui/button"
import {/*CreditCard, Banknote,*/ Edit2, Soup, Coffee} from "lucide-react"
// import {IconWallet} from "@tabler/icons-react";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/components/ui/select";
import {cn} from "@/lib/utils";
import {Separator} from "@/components/ui/separator";
import {formatShortNumber} from "@/lib/numbers";

const cartItems = [
  { title: "Teh Manis Banget", price: 15000, quantity: 1 },
  { title: "Daging Enak Banget", price: 99000, quantity: 1 },
]

export function Cart() {
  const subtotal = cartItems.reduce((acc, item) => acc + item.price * item.quantity, 0)
  const tax = subtotal * 0.05
  const service = subtotal * 0.10
  const total = subtotal + tax + service

  return (
    <div className="flex flex-col h-full select-none">
      <div className="p-4 border-b flex justify-between items-center">
        <h2 className="text-xl font-bold">
          Order
          <span className="text-xs font-mono text-muted-foreground">
            #001
          </span>
        </h2>

        <p className="text-sm text-gray-500 flex gap-2 cursor-pointer">
          Floyd Miles
          <Edit2 className="h-2 w-2" />
        </p>
      </div>

      <div className="p-4 border-b space-y-4">
        <div className="flex gap-2">
          <Button variant="secondary" className="flex-1 rounded-full cursor-pointer bg-green-50 text-green-600">
            Dine In - TA1
          </Button>
          <Button variant="outline" className="flex-1 rounded-full cursor-pointer hover:bg-green-50 hover:text-green-600">
            Take Away
          </Button>
        </div>

       <Select defaultValue="ta1">
         <SelectTrigger className="w-full rounded-full uppercase">
           <SelectValue placeholder="Select Table" />
         </SelectTrigger>
         <SelectContent position="popper">
           {["ta1", "ta2", "ta3", "ta4", "ta5"].map((table, index) => (
             <SelectItem
               key={table}
               value={table}
               className={cn(
                 "uppercase cursor-pointer",
                 index > 2 && " text-muted-foreground"
               )}
             >
               {table}
               <span className={cn(
                 "text-xs lowercase",
                 index > 2 ? "text-red-500" : "text-green-500"
               )}>
                 {index > 2 ? "*Unavailable" : "*Available"}
               </span>
             </SelectItem>
           ))}
         </SelectContent>
       </Select>
      </div>

      <div className="flex-1 overflow-auto p-4">
        <div className="mb-4 flex items-center justify-between">
          <p className="text-sm text-muted-foreground">1 item{cartItems.length>1&&"s"} selected</p>
          <Button variant="link" className="text-red-500 cursor-pointer">Clear</Button>
        </div>
        {cartItems.map((item, index) => (
          <div key={index} className="flex gap-3 mb-4 border px-2 pt-2 pb-4 rounded-lg">
            <div className="rounded-lg w-20 h-20 bg-white flex items-center justify-center">
              {index === 1 ? <Soup /> : <Coffee />}
            </div>

            <div className="flex-1">
              <div className="flex justify-between items-center mt-1">
                <div>
                  <h4 className="text-sm font-medium">{item.title}</h4>
                  <span className="text-green-600 font-bold">IDR {formatShortNumber(item.price)}</span>
                </div>

                <div className="flex items-center gap-2">
                  <Button variant="outline" className="p-2 h-6 cursor-pointer hover:bg-red-50 hover:text-red-600">-</Button>
                  <span className="text-sm text-gray-500">{item.quantity}</span>
                  <Button variant="outline" className="p-2 h-6 cursor-pointer hover:bg-green-50 hover:text-green-600">+</Button>
                </div>
              </div>

              <div className="grid grid-cols-2 mt-2 space-y-1">
                {index === 0 && (<>
                  <p className="text-xs text-muted-foreground">
                    Cup Size:
                    <span className={cn(
                      "hover:cursor-pointer hover:px-1 ml-2",
                      "hover:border hover:rounded-full font-semibold",
                      "hover:bg-green-50 hover:text-green-600"
                    )}>XL</span>
                  </p>
                  <p className="text-xs text-muted-foreground">
                    Hot or Cold:
                    <span className={cn(
                      "hover:cursor-pointer hover:px-1 ml-2",
                      "hover:border hover:rounded-full font-semibold",
                      "hover:bg-green-50 hover:text-green-600"
                    )}>COLD</span>
                  </p>
                  <p className="text-xs text-muted-foreground">
                    Ice Level (%):
                    <span className={cn(
                      "hover:cursor-pointer hover:px-1 ml-2",
                      "hover:border hover:rounded-full font-semibold",
                      "hover:bg-green-50 hover:text-green-600"
                    )}>25</span>
                  </p>
                </>)}

                {index === 1 && (
                  <p className="text-xs text-muted-foreground">
                    Portion:
                    <span className={cn(
                      "hover:cursor-pointer hover:px-1 ml-2",
                      "hover:border hover:rounded-full font-semibold",
                      "hover:bg-green-50 hover:text-green-600"
                    )}>NORMAL</span>
                  </p>
                )}

                <p className="text-xs text-muted-foreground group">
                  Addons:
                  <span className={cn(
                    "group-hover:cursor-pointer group-hover:px-1 ml-2",
                    "group-hover:border group-hover:rounded-full font-semibold",
                    "group-hover:bg-green-50 group-hover:text-green-600"
                  )}>
                    <span className="group-hover:hidden max-w-[60px] truncate whitespace-nowrap overflow-hidden text-ellipsis inline-block align-bottom">this add on bla bla bla bla lbal bal</span>
                    <span className="hidden group-hover:inline">+</span>
                  </span>
                </p>

                <Separator className="col-span-2"/>

                <p className="text-xs text-muted-foreground col-span-2">
                  Notes:
                  <span className={cn(
                    "hover:cursor-pointer hover:px-1 ml-2 group",
                    "hover:border hover:rounded-full font-semibold",
                    "hover:bg-green-50 hover:text-green-600"
                  )}>
                      <span className="group-hover:hidden">-</span>
                      <span className="hidden group-hover:inline">+</span>
                  </span>
                </p>
              </div>
            </div>
          </div>
        ))}
      </div>

      <div className="border-t p-4">
        <div className="space-y-2 mb-4">
          <div className="flex justify-between text-sm">
            <span className="text-gray-600">Sub Total</span>
            <span>{new Intl.NumberFormat("id-ID", { style: "currency", currency: "IDR" }).format(subtotal)}</span>
          </div>
          <div className="flex justify-between text-sm">
            <span className="text-gray-600">Tax 5%</span>
            <span>{new Intl.NumberFormat("id-ID", { style: "currency", currency: "IDR" }).format(tax)}</span>
          </div>
          <div className="flex justify-between text-sm">
            <span className="text-gray-600">Service 10%</span>
            <span>{new Intl.NumberFormat("id-ID", { style: "currency", currency: "IDR" }).format(service)}</span>
          </div>
          <div className="flex justify-between font-bold">
            <span>Total Amount</span>
            <span>{new Intl.NumberFormat("id-ID", { style: "currency", currency: "IDR" }).format(total)}</span>
          </div>
        </div>

        {/*/!* TODO: hide this if items for dine in or ask to display pay now or others *!/*/}
        {/*<div className="grid grid-cols-3 gap-2 mb-4">*/}
        {/*  <Button variant="outline" className="flex flex-col items-center py-2 h-16  gap-0 cursor-pointer hover:bg-green-50 hover:text-green-600">*/}
        {/*    <Banknote className="h-4 w-4 mb-1" />*/}
        {/*    <span className="text-xs font-light">Cash</span>*/}
        {/*  </Button>*/}
        {/*  <Button variant="outline" className="flex flex-col items-center py-2 h-16 gap-0 cursor-pointer hover:bg-green-50 hover:text-green-600">*/}
        {/*    <CreditCard className="h-4 w-4 mb-1" />*/}
        {/*    <span className="text-xs font-light">Card</span>*/}
        {/*  </Button>*/}
        {/*  <Button variant="outline" className="flex flex-col items-center py-2 h-16 gap-0 cursor-pointer hover:bg-green-50 hover:text-green-600">*/}
        {/*    <IconWallet className="h-4 w-4 mb-1" />*/}
        {/*    <span className="text-xs font-light">E-Wallet</span>*/}
        {/*  </Button>*/}
        {/*</div>*/}

        <Button className="w-full bg-green-600 hover:bg-green-700 text-white h-12 cursor-pointer">
          Make Order
        </Button>
      </div>
    </div>
  )
}
