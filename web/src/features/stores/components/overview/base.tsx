import type {ReactNode} from "react";

export function BaseSection({
   title,
   description,
   additionalHeading,
   children,
   ...props
 }: {
  id?: string,
  title: string
  description: string,
  additionalHeading?: ReactNode,
  children: ReactNode
}) {
  return (
    <div {...props} className="mx-auto w-full max-w-[1400px] px-8">
      <div className="flex flex-col first:pt-12 py-6 gap-4 lg:grid md:grid-cols-12">
        <div className="col-span-4 xl:col-span-5 prose text-sm lg:mr-12">
          <div className="sticky space-y-6 top-12">
            <div className="space-y-2 mb-4">
              <p className="text-primary text-lg font-normal m-0 transition">
                {title}
              </p>
              <p className="text-sm font-light text-muted-foreground m-0 transition">
                {description}
              </p>
            </div>

            {additionalHeading}
          </div>
        </div>
        <div className="col-span-8 xl:col-span-7 flex flex-col gap-6">
          {children}
        </div>
      </div>
    </div>
  )
}