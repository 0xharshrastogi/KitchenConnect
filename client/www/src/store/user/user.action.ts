import { createAsyncThunk } from "@reduxjs/toolkit";
import * as Endpoint from "../../common/endpoints";
import { ContentType, HttpMethod } from "../../common/enums";
import { User, UserCredential } from "../../common/shared";

export const login = createAsyncThunk("user/@LOGIN", async (credential: UserCredential) => {
  const resp = await fetch(Endpoint.LOGIN, {
    method: HttpMethod.POST,
    body: JSON.stringify(credential),
    headers: { ["Content-Type"]: ContentType.Json },
  });

  return (await resp.json()) as { token: string; user: User };
});
