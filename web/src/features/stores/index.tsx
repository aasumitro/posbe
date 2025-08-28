import {InfoSection} from "@/features/stores/components/overview/info-section";
import {ServiceAndTaxRateSection} from "@/features/stores/components/overview/service-tax-rate-section";
import {ShiftTableSection} from "@/features/stores/components/overview/shift-section";
import {FeatureSection} from "@/features/stores/components/overview/feature-section";
import {BaseSection} from "@/features/stores/components/overview/base";
import {Separator} from "@/components/ui/separator";
import {Button} from "@/components/ui/button";
import {IconExternalLink} from "@tabler/icons-react";
import {useNavigate} from "@tanstack/react-router";
import {useEffect} from "react";
import {useStoreShifts, useStoresSetting} from "@/hooks/use-store";
import {useStoreState} from "@/states/store-state";

export function StorePage() {
  const {data: settings, isPending: isLoadSetting} = useStoresSetting()
  const {data: shifts, isPending: isLoadShifts} = useStoreShifts()
  const {setSettings, setShifts} =  useStoreState();

  useEffect(() => {
    if (settings?.data) {
      setSettings(settings?.data ?? null)
    }

    if (shifts?.data) {
      setShifts(shifts?.data ?? null)
    }
  }, [settings?.data, shifts?.data]);

  if (isLoadSetting || isLoadShifts) return <>Loading . . .</>

  const navigate = useNavigate();

  const sections = [
    {
      id: "store-information",
      title: "Store Information",
      description: "Update your store's name, contact, category/type, address, and other essential details.",
      content: <InfoSection />
    },
    {
      id: "service-tax-rate",
      title: "Service Tax Rate",
      description: "Configure your store's service categories, rates, and applicable taxes.",
      content: <ServiceAndTaxRateSection />
    },
    {
      id: "shift-table",
      title: "Store Shift",
      description: "Manage your store's shift schedule, working hours, and staff assignments.",      content: <ShiftTableSection />
    },
    {
      id: "feature-enabled",
      title: "Store Feature",
      description: "Enable or disable key store features like floors, etc.",
      additionalHeading: (
        <Button
          variant="link"
          size="sm"
          type="button"
          className="flex-0 mr-auto text-xs cursor-pointer"
          onClick={async (e) => {
            e.preventDefault();
            await navigate({to: "/stores/floors"})
          }}
        >
          <IconExternalLink />
          Open floor designer
        </Button>
      ),
      content: <FeatureSection />
    }
  ]

  const renderSection = (section: typeof sections[number], index: number) => (
    <div key={section.id}>
      <BaseSection
        id={section.id}
        title={section.title}
        description={section.description}
        additionalHeading={section.additionalHeading}
      >
        {section.content}
      </BaseSection>
      {index !== sections.length - 1 && <Separator className="mt-8" />}
    </div>
  );

  return (
    <div className="w-full pt-4 xl:pt-6 space-y-6">
      <aside className="flex items-center justify-between gap-4 px-8">
        <div className="block">
          <h3 className="text-lg font-semibold mb-2">Store Overview</h3>
          <p className="text-sm text-muted-foreground">
            Update your store details, preferences, and basic configuration settings.
          </p>
        </div>
      </aside>

      <aside className="@container/main flex flex-1 flex-col gap-2 mb-12">
        {sections.map(renderSection)}
      </aside>
    </div>
  )
}