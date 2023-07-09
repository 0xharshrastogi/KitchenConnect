import { FC } from "react";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { routes } from "./routes";

export const App: FC = () => {
  const router = createBrowserRouter(routes);

  return <RouterProvider router={router} />;
};
