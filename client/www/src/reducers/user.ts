import {
  CreateSliceOptions,
  PayloadAction,
  createSlice,
} from "@reduxjs/toolkit";
import { User, type Reducer } from "../common/types";

const initialState: Reducer.IUserReducer = null;

const option: CreateSliceOptions<Reducer.IUserReducer> = {
  name: "user",
  initialState,
  reducers: {
    setUserInfo: (state, action: PayloadAction<User>) => {
      console.log(state, action);
    },
  },
};

export const { reducer } = createSlice(option);
