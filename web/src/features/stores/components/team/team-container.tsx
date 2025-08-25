import {TeamActionAdd} from "@/features/stores/components/team/team-action-add";
import {TeamCard} from "@/features/stores/components/team/team-card";
import {TeamDetailActionSheet} from "@/features/stores/components/team/team-detail-sheet";
import {useUserState} from "@/states/user-state";
import {TeamConfirmDeleteAlertDialog} from "@/features/stores/components/team/team-confirm-delete";

export function TeamContainer() {
  const {users} = useUserState()

  return (
    <aside className="flex flex-wrap gap-4 px-8 py-4">
      <TeamActionAdd />

      {users?.map((user, index) => (
        <TeamCard key={index} user={user} />
      ))}

      <TeamDetailActionSheet />
      <TeamConfirmDeleteAlertDialog />
    </aside>
  )
}