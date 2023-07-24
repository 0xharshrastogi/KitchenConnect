import { login, slice } from "./user";

export const userAction = { ...slice.actions, login };
