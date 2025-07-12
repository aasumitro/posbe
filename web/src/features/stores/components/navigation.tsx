"use client"

import {
  IconBuildingStore,
  IconDatabase,
  IconLibrary,
  IconShoppingCart,
  IconStairs,
  IconUsers,
  IconUsersGroup,
} from "@tabler/icons-react";
import {Children} from "react";
import type { ReactNode, ReactElement } from "react";
import {cn} from "@/lib/utils";
import type {Icon} from "@tabler/icons-react";
import {Link, useRouterState} from "@tanstack/react-router";
import * as React from "react";

const data = {
  users: [
    {
      url: "/stores",
      name: "Overview",
      icon: IconBuildingStore,
    },
    {
      url: "/stores/attributes",
      name: "Attributes",
      icon: IconDatabase,
    },
    {
      url: "/stores/teams",
      name: "Teams",
      icon: IconUsers,
    },
    {
      url: "/stores/floors",
      name: "Floor Plan",
      icon: IconStairs,
    },
    {
      url: "/stores/catalogs",
      name: "Catalogs",
      icon: IconLibrary,
    },
    {
      url: "/stores/customers",
      name: "Customers",
      icon: IconUsersGroup,
    },
    {
      url: "/stores/transactions",
      name: "Transactions",
      icon: IconShoppingCart,
    },
  ],
}

export const Navigation = () => {
  return (
    <MenuContainer>
      <MenuSection label="Sections" items={data.users}/>
    </MenuContainer>
  )
}

export const NavigationInset = ({ className, ...props }: React.ComponentProps<"div">) => {
  return (
    <main
      className={cn(
        "bg-background relative flex w-full flex-1 flex-col",
        "rounded-r-xl xl:mt-0 xl:ml-1/4 xl:ml-64",
        className
      )}
      {...props}
    />
  )
}

const MenuContainer = ({
  children
}: { children: ReactNode }) => {
  return (
    <div className={cn(
      "xl:fixed inset-y-1 z-10 xl:my-1 w-full xl:w-64 pt-4",
      "border-b xl:border-b-0 xl:border-r",
    )}>
      <div className="hidden xl:block">{children}</div>
      <ul className="xl:hidden flex gap-2 overflow-x-auto px-4 pb-3">
        {Children.map(children, (section) => {
          const typedSection = section as ReactElement<{ items: {
              url: string
              name: string
              icon: Icon
              enabled?: boolean
            }[] }>;
          return typedSection.props.items.map((item, index) => {
            return (
              <Menuitem
                key={`${item.name}-${index}`}
                route={item.url}
                name={item.name}
                icon={item.icon}
                enabled={item.enabled}
              />
            )
          });
        })}
      </ul>
    </div>
  );
}

const MenuSection = ({
  label,
  items
}: {
  label: string,
  items: {
    url: string
    name: string
    icon: Icon
    enabled?: boolean
  }[],
}) => {
  return (
    <div className="hidden xl:block">
      <h6 className="py-2 px-4 text-sidebar-foreground/70 text-xs font-medium">{label}</h6>
      <ul className="py-2 px-3 flex flex-col gap-1">
        {items.map((item, index) => (
          <Menuitem key={index} route={item.url} name={item.name} icon={item.icon} enabled={item.enabled}/>
        ))}
      </ul>
    </div>
  );
}

const Menuitem = ({
  route,
  name,
  icon: Icon,
  enabled = true,
}: {
  route: string
  name: string
  icon: Icon
  enabled?: boolean
}) => {
  const { location } = useRouterState();
  const pathname = location.pathname;
  let isActive = pathname === route;

  if (pathname.startsWith('/stores/products') && route === "/stores/catalogs") {
    isActive = true
  }

  return (
    <li
      className={cn(
        "flex items-center gap-2 text-sm rounded-md whitespace-nowrap", {
          "hover:bg-sidebar-accent hover:text-sidebar-accent-foreground": enabled,
          "bg-sidebar-accent text-sidebar-accent-foreground font-medium": isActive && enabled,
          "cursor-not-allowed opacity-50": !enabled,
        })}
    >
      <Link
        to={route}
        onClick={(e) => {
          if (!enabled) e.preventDefault();
        }}
        className="flex items-center gap-2 w-full h-full py-2 px-3"
      >
        <Icon className="h-4 w-4" />
        <span className="whitespace-nowrap">{name}</span>
      </Link>
    </li>
  );
}