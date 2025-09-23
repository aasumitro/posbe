import {Coffee, Soup} from "lucide-react";
import {Badge} from "@/components/ui/badge";
import {Separator} from "@/components/ui/separator";
import {
  DropdownMenu,
  DropdownMenuContent, DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import {Button} from "@/components/ui/button";
import {IconDotsVertical} from "@tabler/icons-react";
import {useActionState} from "@/states/action-state";
import {ProductActionDeleteModalState} from "@/features/stores/components/catalog/product-action-delete";
import {useNavigate} from "@tanstack/react-router";

export function  ProductContainer() {
  const products = [1,2]
  const { setBoolState } = useActionState();
  const navigate = useNavigate();

  return (
    <div className="grid gap-4 grid-cols-2 2xl:grid-cols-3 auto-rows-fr mt-4">
      {products.map((index) => (
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

                <div className="text-muted-foreground text-xs flex gap-2">
                  {index % 2 == 0 ? (
                    <>
                      <p>#foods</p>
                      <p>#meat</p>
                    </>
                  ) : (
                    <>
                      <p>#beverages</p>
                      <p>#tea</p>
                    </>
                  )}
                </div>

                {index == 1 && (
                  <h5 className="text-3xl font-semibold text-green-500">
                    IDR 15K
                  </h5>
                )}

                {index == 2 && (
                  <section className="flex">
                  <span className="text-sm font-normal text-muted-foreground mr-1">
                    from
                    {/*  if variant exists `start from` if no variants `$` */}
                  </span>
                    <div className="flex items-baseline gap-1 text-3xl font-bold tabular-nums leading-none text-green-500">
                      {/* remove `$` if variant not exist */}
                      IDR 99K <span className="text-sm font-normal text-muted-foreground">/ 150gr</span>
                    </div>
                  </section>
                )}
              </div>

              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button className="ml-auto my-auto" variant="ghost">
                    <IconDotsVertical />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                  <DropdownMenuLabel>Actions</DropdownMenuLabel>
                  <DropdownMenuSeparator/>
                  <DropdownMenuItem onClick={async (e) => {
                    e.preventDefault();
                    await navigate({to: `/stores/products/${index}`})
                  }}>Edit</DropdownMenuItem>
                  <DropdownMenuItem variant="destructive" onClick={(e) => {
                    e.preventDefault();
                    setBoolState(ProductActionDeleteModalState, true)
                  }}>Delete</DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
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
                      >
                        {level}
                      </Badge>
                    ))}
                  </div>
                </div>
              )}
            </div>
          </div>

          <div className="flex flex-row border-t p-4">
            <div className="flex w-full items-center gap-2">
              <div className="grid flex-1 auto-rows-min gap-0.5">
                <div className="text-xs text-muted-foreground">This Year</div>
                <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
                  562
                  <span className="text-sm font-normal text-muted-foreground">
                  sales
                </span>
                </div>
              </div>
              <Separator orientation="vertical" className="mx-2 h-10 w-px" />
              <div className="grid flex-1 auto-rows-min gap-0.5">
                <div className="text-xs text-muted-foreground">This Week</div>
                <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
                  73
                  <span className="text-sm font-normal text-muted-foreground">
                  sales
                </span>
                </div>
              </div>
              <Separator orientation="vertical" className="mx-2 h-10 w-px" />
              <div className="grid flex-1 auto-rows-min gap-0.5">
                <div className="text-xs text-muted-foreground">Today</div>
                <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
                  14
                  <span className="text-sm font-normal text-muted-foreground">
                  sales
                </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      ))}
    </div>
  )
}