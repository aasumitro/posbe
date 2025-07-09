import {TeamContainer} from "@/features/stores/components/team/team-container";

export function TeamMemberPage() {
  return (
    <div className="w-full pt-4 xl:pt-6 space-y-6">
      <aside className="flex items-center justify-between gap-4 px-8">
        <div className="block">
          <h3 className="text-lg font-semibold mb-2">Teams</h3>
          <p className="text-sm text-muted-foreground">
            Manage your organization's teams and their members. Add, view, and organize team structures to match your operational needs.
          </p>
        </div>
      </aside>

      <TeamContainer />
    </div>
  )
}