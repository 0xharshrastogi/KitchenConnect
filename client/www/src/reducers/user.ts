import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { User, type Reducer } from "../common/types";

const initialState: Reducer.IUserReducer = null;

export const { reducer } = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUserInfo: (_state, action: PayloadAction<User>) => {
      console.log(_state, action);
    },
  },
});
