import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { authAction } from "../action/user.action";
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

  extraReducers: (builder) => {
    builder.addCase(authAction.login.pending, () => {
      return { loading: true, info: null };
    });

    builder.addCase(authAction.login.fulfilled, (state, action) => {
      state.loading = false;
      state.info = action.payload.user;

      return state;
    });

    builder.addCase(authAction.login.rejected, (_, action) => {
      console.error(action.error.message, action.error);
      return { loading: false, info: null };
    });
  },
});
