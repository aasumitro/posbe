import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/avatar"
import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "@/components/ui/hover-card"
import {useUserDetail} from "@/hooks/use-user";

export function UserHoverCard({id}: {id: number}) {
  const {data: user, isPending} = useUserDetail(id)

  if (isPending) return <>Loading . . .</>

  return (
    <HoverCard>
      <HoverCardTrigger asChild>
        <Avatar key={id.toString()}>
          <AvatarImage
            src="https://github.com/evilrabbit.png"
            alt={id.toString()}
          />
          <AvatarFallback>u{id}</AvatarFallback>
        </Avatar>
      </HoverCardTrigger>
      <HoverCardContent className="w-80">
        <div className="flex justify-between gap-4">
          <Avatar>
            <AvatarImage
              src="https://github.com/evilrabbit.png"
              alt={user?.data?.username}
            />
            <AvatarFallback>{user?.data?.username?.slice(0, 2)}</AvatarFallback>
          </Avatar>
          <div className="space-y-1">
            <h4 className="text-sm font-semibold">@{user?.data?.username}</h4>
            <div className="text-muted-foreground text-xs">
              {user?.data?.role?.name}
            </div>
            <p className="text-sm">
              {user?.data?.email}
            </p>
          </div>
        </div>
      </HoverCardContent>
    </HoverCard>
  )
}

