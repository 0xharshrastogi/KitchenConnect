import { useRef } from "react";
import { shallowEqual } from "react-redux";
import { api } from "../api";
import { type User, type UserCredential } from "../common/shared";
import { ErrorHelper } from "../helpers";
import { userAction } from "../store/actions";
import { useAppDispatch, useAppSelector } from "./useAppSelector";

type AuthChangeFn = (status: boolean) => void;

type LoginFn = (
  credential: UserCredential,
  config?: {
    success?: () => void;

    error: (error: Error) => void;
  }
) => Promise<void>;

interface UserHookSharedAPI {
  login: LoginFn;

  /**
   * Registers a callback function to be invoked when the authentication status changes.
   *
   * @param {boolean} status - A boolean value indicating the current authentication status.
   * `true` represents authenticated, and `false` represents unauthenticated.
   * @returns
   */
  onAuthenticationStatusChange(cb: AuthChangeFn): void;
}

interface UserAuthenticatedAPI {
  authenticated: true;

  user: User;
}

interface UnAuthenticatedHookAPI {
  authenticated: false;
}

type UseAuthenticationHookAPI = UserHookSharedAPI & (UserAuthenticatedAPI | UnAuthenticatedHookAPI);

export const useAuthentication = (): UseAuthenticationHookAPI => {
  const user = useAppSelector(({ user }) => user, shallowEqual);
  const dispatch = useAppDispatch();
  const onAuthStatusChange = useRef<((status: boolean) => void) | undefined>();

  const login: LoginFn = async (value, config) => {
    try {
      const { user } = await api.auth.login(value);
      onAuthStatusChange.current?.(true);
      dispatch(userAction.setUser(user));
      config?.success?.();
    } catch (error) {
      config?.error?.(ErrorHelper.toHumanReadable(error));
    }
  };

  const onAuthenticationStatusChange = (onStatusChange: (status: boolean) => void) => {
    onAuthStatusChange.current = onStatusChange;
  };

  const shared: UserHookSharedAPI = { login, onAuthenticationStatusChange };

  if (!user.loggedIn) {
    return { ...shared, authenticated: false };
  }

  return { ...shared, authenticated: true, user: user.info };
};
