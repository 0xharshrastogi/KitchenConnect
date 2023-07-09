import { RouteObject } from "react-router-dom";
import { ErrorBoundary, Login } from "./feature";

export const routes: RouteObject[] = [
  { path: "login", Component: Login, ErrorBoundary },
  { path: "*", Component: ErrorBoundary.NotFound },
];
