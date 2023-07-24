import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { User } from "../../common/shared";
import { UserReducerState } from "../shared";
import { login } from "./user.action";

const getDefaultState = (): UserReducerState => ({
  loggedIn: false,
});

export const slice = createSlice({
  name: "USER",
  initialState: getDefaultState(),
  reducers: {
    setUser: (_, action: PayloadAction<User>) => {
      return { loggedIn: true, info: action.payload };
    },
  },

  extraReducers: (builder) => {
    builder.addCase(login.pending, () => {
      return { loggedIn: false };
    });

    builder.addCase(login.fulfilled, (_, action) => {
      action.meta.requestStatus;
      return { loggedIn: true, info: action.payload.user };
    });

    builder.addCase(login.rejected, () => ({ loggedIn: false }));
  },
});
