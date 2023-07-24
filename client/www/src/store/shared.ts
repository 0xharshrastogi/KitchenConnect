import { User } from "../common/shared";

export interface UserAuthenticatedStore {
  loggedIn: true;

  info: User;
}

export interface UserNotAuthenticatedStore {
  loggedIn: false;
}

export type UserReducerState = UserAuthenticatedStore | UserNotAuthenticatedStore;
