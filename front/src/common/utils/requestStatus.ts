import * as Sentry from "@sentry/browser";
import { t } from "./intl";
import { IPromiseFailureAction } from "./reduxUtils";
import { IApiErrorAction } from "./requestTypes";

export enum RequestStatusType {
  active = "active",
  failed = "failed",
  inactive = "inactive",
  successful = "successful"
}
export interface IRequestStatusProgress {
  type: RequestStatusType.active;
}

export interface IRequestStatusFailed {
  type: RequestStatusType.failed;
  errorCode?: number;
  error?: string;
}

export interface IRequestStatusSuccessful {
  type: RequestStatusType.successful;
}

export interface IRequestStatusInactive {
  type: RequestStatusType.inactive;
  message?: string;
}

export type RequestStatus =
  | IRequestStatusProgress
  | IRequestStatusFailed
  | IRequestStatusInactive
  | IRequestStatusSuccessful;

/**
 * Extract error text from the server response
 * @param action
 */
export const extractRequestError = (
  action: IApiErrorAction | IPromiseFailureAction
): string => {
  try {
    if (action.error && action.error.response && action.error.response.data) {
      if (action.error.response.data.Error) {
        return action.error.response.data.Error;
      }

      if (action.error.response.data.msg) {
        return action.error.response.data.msg;
      }

      if (action.error.response.data.message) {
        return action.error.response.data.message;
      }

      if (typeof action.error.response.data.error_description === "string") {
        return action.error.response.data.error_description;
      }

      if (typeof action.error.response.data.error === "string") {
        return action.error.response.data.error;
      }

      if (
        action.error.response.data.error &&
        action.error.response.data.error.data
      ) {
        return action.error.response.data.error.data;
      }
    } else if (action.error && (action as IPromiseFailureAction).errorText) {
      return (action as IPromiseFailureAction).errorText;
    } else if (typeof action.error.message === "string") {
      return action.error.message;
    }

    return t("errors.unknown");
  } catch (error) {
    Sentry.configureScope(scope => {
      scope.setExtra("action", action);
      Sentry.captureException(error);
    });
    return t("errors.unknown");
  }
};

export function requestInProgress(progress?: number) {
  return { progress, type: RequestStatusType.active } as RequestStatus;
}

export function requestFailed(error?: string, errorCode?: number) {
  return { error, errorCode, type: RequestStatusType.failed } as RequestStatus;
}

const REQUEST_SUCCESSFUL = { type: RequestStatusType.successful };
export function requestSuccessful() {
  return REQUEST_SUCCESSFUL as RequestStatus;
}

const REQUEST_INACTIVE_EMPTY = {
  type: RequestStatusType.inactive
} as RequestStatus;
export function requestInactive(message?: string) {
  if (message === undefined) return REQUEST_INACTIVE_EMPTY;
  return { type: RequestStatusType.inactive, message } as RequestStatus;
}

export function isRequestInProgress(status: RequestStatus) {
  return status.type === RequestStatusType.active;
}

export function isRequestFailed(status: RequestStatus) {
  return status.type === RequestStatusType.failed;
}

export function isRequestSuccessful(status: RequestStatus) {
  return status.type === RequestStatusType.successful;
}

export function isRequestInactive(status: RequestStatus) {
  return status.type === RequestStatusType.inactive;
}

export function getRequestError(...statuses: RequestStatus[]) {
  const status = statuses.find(
    status => status.type === RequestStatusType.failed
  );

  if (status && status.type === RequestStatusType.failed) {
    return status.error ? status.error : "Error";
  }

  return undefined;
}

export function getRequestErrorCode(...statuses: RequestStatus[]) {
  const status = statuses.find(
    status => status.type === RequestStatusType.failed
  );

  if (status && status.type === RequestStatusType.failed) {
    return status.errorCode;
  }

  return undefined;
}

export const getRequestsError = getRequestError;
