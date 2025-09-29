import {IconPlant} from "@tabler/icons-react";

export function AppComingSoon({feature, type}: {feature: string, type: string}) {
  return (
    <div className="text-center py-12">
      <IconPlant className="my-8 w-24 h-24 mx-auto"/>
      <h4 className="text-primary text-xl font-bold tracking-tight">
        Coming Soon
      </h4>
      <p className="text-secondary-foreground text-md font-normal">
        We’re working hard to bring you powerful tools to view and manage your {feature}. <br />
        Stay tuned for updates — this {type} will be available soon!
      </p>
    </div>
  )
}