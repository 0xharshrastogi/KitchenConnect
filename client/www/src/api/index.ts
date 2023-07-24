import { AuthApiHandler, IAuthApiHandler } from "./AuthApiHandler";

interface ApiHandler {
  auth: IAuthApiHandler;
}

export const api: ApiHandler = {
  auth: new AuthApiHandler(),
};

export * from "./endpoints";
