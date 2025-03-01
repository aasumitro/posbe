import {Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {ChartContainer} from "@/components/ui/chart.tsx";
import {Bar, BarChart, Rectangle} from "recharts";
import {Separator} from "@/components/ui/separator.tsx";
import {Tooltip, TooltipContent, TooltipTrigger} from "@/components/ui/tooltip.tsx";

export const ProductCard = () => {
  return (
    <Card className="w-full lg:w-96 h-60 hover:bg-gray-100">
      <CardHeader className="p-4 pb-0 flex flex-row gap-2">
        <div className="w-[80px] h-fit">
          <img
            className="w-full h-full"
            alt="placeholder"
            src="https://placehold.jp/150x150.png"
          />
        </div>
        <div className="w-fit">
          <CardTitle className="flex gap-2">
            Wagyu A5 Steak
            <Tooltip>
              <TooltipTrigger className="text-sm font-normal text-muted-foreground mr-1" asChild>
                <span>(*2)</span>
              </TooltipTrigger>
              <TooltipContent>
                <span> Available 2 different variants </span>
              </TooltipContent>
            </Tooltip>
          </CardTitle>
          <CardDescription>
            #foods #meat
            <Tooltip>
              <TooltipTrigger className="text-left line-clamp-1">
                The highest yield grade and meat quality grade for Wagyu beef is A5, where A represents the yield grade, and 5 represents the meat quality grade. A5 Wagyu beef denotes meat with ideal firmness and texture, coloring, yield, and beef marbling score.
              </TooltipTrigger>
              <TooltipContent className="w-[400px]">
                <p>The highest yield grade and meat quality grade for Wagyu beef is A5, where A represents the yield grade,
                  and 5 represents the meat quality grade. A5 Wagyu beef denotes meat with ideal firmness and texture,
                  coloring, yield, and beef marbling score.</p>
              </TooltipContent>
            </Tooltip>
          </CardDescription>
        </div>
      </CardHeader>
      <CardContent className="flex flex-row items-center gap-4 p-4">
        <section className="flex">
          <span className="text-sm font-normal text-muted-foreground mr-1">
            from
            {/*  if variant exists `start from` if no variants `$` */}
          </span>
          <div className="flex items-baseline gap-1 text-3xl font-bold tabular-nums leading-none">
            {/* remove `$` if variant not exist */}
            $100 <span className="text-sm font-normal text-muted-foreground">/ 250gr</span>
          </div>
        </section>

        <ChartContainer
          config={{
            steps: {
              label: "Steps",
              color: "hsl(var(--chart-1))",
            },
          }}
          className="ml-auto w-[72px]"
        >
          <BarChart
            accessibilityLayer
            margin={{
              left: 0,
              right: 0,
              top: 0,
              bottom: 0,
            }}
            data={[
              {
                date: "2024-01-01",
                steps: 2000,
              },
              {
                date: "2024-01-02",
                steps: 2100,
              },
              {
                date: "2024-01-03",
                steps: 2200,
              },
              {
                date: "2024-01-04",
                steps: 1300,
              },
              {
                date: "2024-01-05",
                steps: 1400,
              },
              {
                date: "2024-01-06",
                steps: 2500,
              },
              {
                date: "2024-01-07",
                steps: 1600,
              },
            ]}
          >
            <Bar
              dataKey="steps"
              fill="var(--color-steps)"
              radius={2}
              fillOpacity={0.2}
              activeIndex={6}
              activeBar={<Rectangle fillOpacity={0.8}/>}
            />
          </BarChart>
        </ChartContainer>
      </CardContent>
      <CardFooter className="flex flex-row border-t p-4">
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
      </CardFooter>
    </Card>
  )
}