import {Badge} from "@/components/ui/badge";
import {Button} from "@/components/ui/button";
import { IconPlus } from "@tabler/icons-react";
import {Coffee, Soup} from "lucide-react";

export function MenuList() {
  const menus = [1,2]

  return (
    <div className="grid gap-4 grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 auto-rows-fr">
      {menus.map((index) => (
        <div
          key={index}
          className="border-1 rounded-xl p-4 space-y-6 select-none flex flex-col h-full"
        >
          <div className="flex-grow space-y-6">
            <div className="flex gap-4">
              <div className="rounded-lg w-20 h-20 bg-gray-50/50 flex items-center justify-center">
                {index % 2 == 0 ? <Soup /> : <Coffee />}
              </div>
              <div className="space-y-1">
                <h5 className="text-xl font-semibold">
                  {index % 2 == 0 ? "Daging Enak Banget" : "Teh Manis Banget"}
                </h5>
                <p className="text-sm text-muted-foreground">
                  {index % 2 == 0 ? "24 Available" : "99 Available"}
                </p>
                <h5 className="text-lg font-semibold text-green-500">
                  {index % 2 == 0 ? " IDR 99K" : " IDR 15K"}
                </h5>
              </div>
            </div>

            <div className="grid grid-cols-2 space-y-4">
            <div className="flex flex-col gap-2">
              <p className="text-sm">
                {index % 2 == 0 ? "Portion?" : " Cup Size?"}
              </p>
              <div className="flex flex-row gap-2">
                {index % 2 == 0 && (
                  ["HALF", "NORMAL"].map((size) => (
                    <Badge
                      key={size}
                      variant="outline"
                      className="hover:bg-green-50 hover:text-green-600 cursor-pointer"
                    >
                      {size}
                    </Badge>
                  ))
                )}

                {index % 2 != 0 && (
                  ["S", "M", "L", "XL"].map((size) => (
                    <Badge
                      key={size}
                      variant="outline"
                      className="hover:bg-green-50 hover:text-green-600 cursor-pointer"
                    >
                      {size}
                    </Badge>
                  ))
                )}
              </div>
            </div>

            {index % 2 != 0 && (
              <div className="flex flex-col gap-2">
                <p className="text-sm">
                  Hot or Cold?
                </p>
                <div className="flex flex-row gap-2">
                  {["HOT", "COLD"].map((level) => (
                    <Badge
                      key={level}
                      variant="outline"
                      className="hover:bg-green-50 hover:text-green-600 cursor-pointer"
                    >
                      {level}
                    </Badge>
                  ))}
                </div>
              </div>
            )}

            {index % 2 != 0 && (
              <div className="flex flex-col gap-2">
              <p className="text-sm">
                Ice Level (%)?
              </p>
              <div className="flex flex-row gap-2">
                {[10, 25, 50].map((level) => (
                  <Badge
                    key={level}
                    variant="outline"
                    className="hover:bg-green-50 hover:text-green-600 cursor-pointer"
                  >
                    {level}
                  </Badge>
                ))}
              </div>
            </div>
            )}
          </div>
          </div>

          <Button
            className="w-full cursor-pointer rounded-full hover:bg-green-50 hover:text-green-600"
            variant="outline"
            disabled
          >
            <IconPlus className="w-4 h-4"/>
            Add to Cart
          </Button>
        </div>
      ))}
    </div>
  )
}