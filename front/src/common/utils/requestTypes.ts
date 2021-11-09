import { RequestActionMeta } from "redux-saga-requests";
import { AxiosRequestConfig, AxiosResponse } from "axios";

export interface IExtendedRequestActionMeta extends RequestActionMeta {
  noRefresh?: boolean;
  noAuth?: boolean;
  redirectOnSuccess?: string;
  id?: string;
}

export interface IRequestConfig<D = any, P = D> extends AxiosRequestConfig {
  data?: D;
  params?: P;
}

export interface IExtendedRequestAction<D = any> {
  type: string;
  request: IRequestConfig<D>;
  meta?: IExtendedRequestActionMeta;
}

export interface IExtendedMultiRequestsAction<D = any> {
  type: string;
  request: IRequestConfig<D>[];
  meta?: IExtendedRequestActionMeta;
}

export interface IExtendedAxiosResponseMeta<A>
  extends IExtendedRequestActionMeta {
  requestAction: A;
}

export interface IExtendedAxiosResponse<T, A = IExtendedRequestAction> {
  type: string;
  data: T;
  response: AxiosResponse<T>;
  meta: IExtendedAxiosResponseMeta<A>;
}

export interface IApiErrorResponse {
  response?: {
    status?: number;
    data?: {
      Error?: string;
      Code?: string;
      Method?: string;
      message?: string;
      msg?: string;
      code?: number;
      status?: number;
      error_description?: string;
      error?:
        | {
            data?: string;
          }
        | string;
    };
  };
  message?: string;
}

export interface IApiErrorAction<T = any, A = IExtendedRequestAction>
  extends IExtendedAxiosResponse<T, A> {
  error: IApiErrorResponse;
}

export type SendRequestResponse<T = any> =
  | { response: AxiosResponse<T> }
  | IApiErrorAction<T>;

export type SendMultiRequestResponse<T = any> =
  | { response: AxiosResponse<T>[] }
  | IApiErrorAction<T>;
