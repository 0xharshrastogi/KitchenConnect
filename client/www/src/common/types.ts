export * as Reducer from "./reducer.type";
export * from "./shared";

export type Func<TParams extends any[], TReturn> = (...args: TParams) => TReturn;
