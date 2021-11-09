import { FormApi } from "final-form";

export type ValidateFn = (values: object) => object | Promise<object>;

export type onFormSubmit = (
  values: {},
  form?: FormApi,
  callback?: (errors?: object) => void
) => object | Promise<object | undefined> | undefined | void;

export type FormErrors<T> = Partial<{ [key in keyof T]: string }>;
