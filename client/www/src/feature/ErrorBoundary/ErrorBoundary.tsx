import { FC, ReactNode } from "react";
import { useLocation } from "react-router-dom";
import "./ErrorBoundary.css";

type ErrorBoundaryProps =
  | { type?: "route-not-found"; path: string }
  | { type?: "default" };

type TErrorBoundaryFC = FC<ErrorBoundaryProps> & { NotFound: FC };

export const ErrorBoundary: TErrorBoundaryFC = (props) => {
  const { type } = props;
  let children: ReactNode = null;

  if (type === "route-not-found") {
    const { path } = props;

    children = (
      <div>
        <h1>404 : Route Not Found</h1>
        <span>
          Invalid route{" "}
          <code className="path-value">
            {window.location.host}
            {path}
          </code>
        </span>
      </div>
    );
  }
  return <section className="error-boundary-container">{children}</section>;
};

export const NotFound: FC = () => {
  const location = useLocation();

  return <ErrorBoundary type="route-not-found" path={location.pathname} />;
};

ErrorBoundary.NotFound = NotFound;
