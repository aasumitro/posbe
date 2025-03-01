import {useLocation} from "react-router-dom";

export const isCurrentRoute = (path: string) => {
  const location = useLocation();
  const currentPath = location.pathname.replace(/\/$/, '').toLowerCase();
  const targetPath = path.replace(/\/$/, '').toLowerCase();
  return currentPath.includes(targetPath);
};