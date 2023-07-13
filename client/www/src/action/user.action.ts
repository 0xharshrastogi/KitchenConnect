import { createAsyncThunk } from "@reduxjs/toolkit";
import { api } from "../api";
import { UserCredential } from "../common/shared";

const login = createAsyncThunk("user/login-user", async (credential: UserCredential) => {
  return await api.auth.login(credential);
});

export const authAction = { login };
