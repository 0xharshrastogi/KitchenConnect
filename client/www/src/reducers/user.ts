import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { User } from "../common/types";

type UserReducerState = {
  loading: boolean;
  info: User | null;
};

const initialState: UserReducerState = {
  loading: false,
  info: null,
};

export const { reducer, actions } = createSlice({
  name: "user",
  initialState,
  reducers: {
    reset: (state, action: PayloadAction<{ loading: true } | undefined>) => {
      if (!action.payload) return initialState;
      return { ...state, loading: action.payload.loading };
    },

    activate: (state, action: PayloadAction<User>) => {
      return { ...state, loading: false, user: action.payload };
    },
  },
});
