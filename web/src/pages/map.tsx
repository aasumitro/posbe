import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs.tsx";
import { Button } from "@/components/ui/button.tsx";
import { Edit3 } from "lucide-react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card.tsx";

export function Map() {
  return (<>
    <main className="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
      <Tabs defaultValue="all">
        <div className="flex items-center">
          <TabsList>
            <TabsTrigger value="all">1st Floor</TabsTrigger>
            <TabsTrigger value="active">2nd Floor</TabsTrigger>
          </TabsList>
          <div className="ml-auto flex items-center gap-2">
            <Button size="sm" className="h-8 gap-1">
              <Edit3 className="h-3.5 w-3.5" />
              <span className="sr-only sm:not-sr-only sm:whitespace-nowrap">
                Edit Map
              </span>
            </Button>
          </div>
        </div>
        <TabsContent value="all">
          <Card x-chunk="dashboard-06-chunk-0">
            <CardHeader>
              <CardTitle>Store Maps</CardTitle>
              <CardDescription>
                Manage your store floor, table and room.
              </CardDescription>
            </CardHeader>
            <CardContent className="min-h-[calc(100vh_-_theme(spacing.16))]">

            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
    </main>
  </>)
}
