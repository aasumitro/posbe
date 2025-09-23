import { createFileRoute } from "@tanstack/react-router"
import { TeamMemberPage } from "@/features/stores/teams";

export const Route = createFileRoute("/_authenticated/stores/teams")({
  component: TeamMemberPage,
})