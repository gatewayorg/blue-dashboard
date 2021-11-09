import { AnyAction } from "redux";
import {
  extractRequestError,
  requestFailed,
  requestInactive,
  requestInProgress,
  requestSuccessful
} from "./requestStatus";
import {
  IApiErrorAction,
  IApiErrorResponse,
  IExtendedAxiosResponse,
  IExtendedRequestAction
} from "./requestTypes";

/**
 * Redux related utils
 */

type Reducer<State> = (State: any, Action: any) => State;

const createReducer = (
  initialState: any,
  handlers: { [key: string]: Reducer<any> }
): Reducer<any> => {
  return (state: any = initialState, action: Action): any => {
    return handlers.hasOwnProperty(action.type)
      ? handlers[action.type](state, action)
      : state;
  };
};

// taken from https://github.com/Hotell/rex-tils
type Action<T extends string = string, P = void> = P extends void
  ? Readonly<{ type: T }>
  : Readonly<{ type: T; payload: P }>;

export function createAction<T extends string>(type: T): Action<T>;
export function createAction<T extends string, P>(
  type: T,
  payload: P
): Action<T, P>;

export function createAction<T extends string, P>(type: T, payload?: P) {
  return payload === undefined ? { type } : { type, payload };
}

export interface IPromiseActionMeta {
  debug?: boolean;
}

export interface IPromiseActionPayload<F extends () => Promise<R>, R = any> {
  func: F;
  params: Parameters<F>;
}

export interface IPromiseAction<F extends () => Promise<R>, R = any>
  extends AnyAction {
  type: string;
  promises: IPromiseActionPayload<F>[];
  meta?: IPromiseActionMeta;
  hasSinglePayload: boolean;
}

export interface IPromiseRequestAction<
  F extends () => Promise<R> = any,
  R = any
> {
  type: string;
  meta: {
    requestAction: IPromiseAction<F, R>;
  };
}

export interface IPromiseSuccessAction<
  F extends () => Promise<R> = any,
  R = any
> {
  type: string;
  data: R;
  meta: {
    requestAction: IPromiseAction<F, R>;
  };
}

export interface IPromiseFailureAction<
  F extends () => Promise<R> = any,
  R = any
> {
  type: string;
  error: IApiErrorResponse;
  errorText: string;
  meta: {
    requestAction: IPromiseAction<F, R>;
  };
}

export function createPromiseAction<F extends () => Promise<R>, R = any>(
  type: string,
  payload: IPromiseActionPayload<F> | IPromiseActionPayload<F>[],
  meta?: IPromiseActionMeta
): IPromiseAction<F> {
  return {
    type,
    promises: Array.isArray(payload) ? payload : [payload],
    meta,
    hasSinglePayload: !Array.isArray(payload)
  };
}

/**
 * Creates typed API reducer
 * Generic types:
 *    * T - Store State Type
 *    * A - Response action
 *    * R - Request action
 *    * C - Reset action
 * @param action
 * @param statusProperty
 * @param subReducers
 */
export const createAPIReducer = <
  T,
  A = IExtendedAxiosResponse<any>,
  R = IExtendedRequestAction,
  C = AnyAction
>(
  action: string,
  statusProperty: keyof T,
  subReducers?: {
    onRequest?: (state: T, action: R) => T;
    onError?: (state: T, action: IApiErrorAction) => T;
    onSuccess?: (state: T, action: A) => T;
    onReset?: (state: T, action: C) => T;
  }
) => ({
  [action]: (state: T, action: R): T => {
    const newState: T = {
      ...state,
      [statusProperty]: requestInProgress()
    };
    return {
      ...(subReducers && subReducers.onRequest
        ? subReducers.onRequest(newState, action)
        : newState)
    };
  },
  [`${action}_SUCCESS`]: (state: T, action: A): T => {
    const newState: T = {
      ...state,
      [statusProperty]: requestSuccessful()
    };
    return {
      ...(subReducers && subReducers.onSuccess
        ? subReducers.onSuccess(newState, action)
        : newState)
    };
  },
  [`${action}_ERROR`]: (state: T, action: IApiErrorAction): T => {
    const newState: T = {
      ...state,
      [statusProperty]: requestFailed(
        extractRequestError(action),
        action?.error?.response?.status
      )
    };
    return {
      ...(subReducers && subReducers.onError
        ? subReducers.onError(newState, action)
        : newState)
    };
  },
  [`${action}_RESET`]: (state: T, action: C): T => {
    const newState: T = {
      ...state,
      [statusProperty]: requestInactive()
    };
    return {
      ...(subReducers && subReducers.onReset
        ? subReducers.onReset(newState, action)
        : newState)
    };
  }
});

export { createReducer };
