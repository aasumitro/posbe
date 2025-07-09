import {TeamActionAdd} from "@/features/stores/components/team/team-action-add";
import {TeamCard} from "@/features/stores/components/team/team-card";
import {TeamDetailActionSheet} from "@/features/stores/components/team/team-detail-sheet";

export function TeamContainer() {
  const teams = [1]

  return (
    <aside className="flex flex-wrap gap-4 px-8 py-4">
      <TeamActionAdd />

      {teams.map((index) => (
        <TeamCard key={index} />
      ))}

      <TeamDetailActionSheet />
    </aside>
  )
}